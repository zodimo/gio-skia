// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	gpaint "gioui.org/op/paint"
	"github.com/zodimo/gio-skia/pkg/stroke"
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
)

// Compile-time check that canvas implements Canvas interface
var _ Canvas = (*canvas)(nil)

type canvas struct {
	ops   *op.Ops
	stack []context
}

type context struct {
	xform f32.Affine2D
}

// NewCanvas returns a Canvas implementation backed by Gio's GPU renderer.
func NewCanvas(ops *op.Ops) Canvas {
	return &canvas{
		ops: ops,
		stack: []context{{
			xform: f32.Affine2D{},
		}},
	}
}

// ── State management ───────────────────────────────────────────────────

func (c *canvas) Save() int {
	saveCount := len(c.stack)
	top := c.stack[len(c.stack)-1]
	c.stack = append(c.stack, top)
	return saveCount
}

func (c *canvas) Restore() {
	if len(c.stack) > 1 {
		c.stack = c.stack[:len(c.stack)-1]
	}
}

func (c *canvas) Concat(matrix SkMatrix) {
	top := &c.stack[len(c.stack)-1]
	// Convert current f32.Affine2D to SkMatrix, concat, then convert back
	currentMatrix := affine2DToSkMatrix(top.xform)
	// Create a new matrix for the result
	resultMatrix := impl.NewMatrixIdentity()
	resultMatrix.SetConcat(matrix, currentMatrix)
	top.xform = skMatrixToAffine2D(resultMatrix)
}

func (c *canvas) Translate(dx, dy Scalar) {
	matrix := impl.NewMatrixTranslate(dx, dy)
	c.Concat(matrix)
}

func (c *canvas) Scale(sx, sy Scalar) {
	matrix := impl.NewMatrixScale(sx, sy)
	c.Concat(matrix)
}

func (c *canvas) Rotate(degrees Scalar) {
	matrix := impl.NewMatrixRotate(degrees)
	c.Concat(matrix)
}

// ── Convenience methods ───────────────────────────────────────────────────

func (c *canvas) TranslateFloat32(x, y float32) {
	c.Translate(Scalar(x), Scalar(y))
}

func (c *canvas) ScaleFloat32(x, y float32) {
	c.Scale(Scalar(x), Scalar(y))
}

func (c *canvas) RotateFloat32(degrees float32) {
	c.Rotate(Scalar(degrees))
}

// drawPathInternal is the internal implementation that handles the actual drawing.
func (c *canvas) drawPathInternal(path SkPath, paint SkPaint) {
	// Convert SkPaint to our internal Paint type for rendering
	internalPaint := skPaintToPaint(paint)
	transformSave := op.Affine(c.stack[len(c.stack)-1].xform).Push(c.ops)
	defer transformSave.Pop()

	if path.IsEmpty() {
		return
	}

	// Get path data for iteration
	verbCount := path.CountVerbs()
	verbs := make([]enums.PathVerb, verbCount)
	path.GetVerbs(verbs)

	pointCount := path.CountPoints()
	points := make([]models.Point, pointCount)
	path.GetPoints(points)

	conicWeights := path.ConicWeights()

	// Use go-skia-support's PathIter for proper path iteration
	iter := impl.NewPathIter(points, verbs, conicWeights)

	// Build GioUI paths
	var b clip.Path
	b.Begin(c.ops)
	// Build stroke.Path in parallel
	var s stroke.Path
	var start f32.Point
	var current f32.Point
	hasStart := false

	for rec := iter.Next(); rec != nil; rec = iter.Next() {
		verb := rec.Verb
		pts := rec.Points

		if len(pts) == 0 {
			continue
		}

		switch verb {
		case enums.PathVerbMove:
			pt := f32.Pt(float32(pts[0].X), float32(pts[0].Y))
			b.MoveTo(pt)
			s.Segments = append(s.Segments, stroke.MoveTo(pt))
			start = pt
			current = pt
			hasStart = true

		case enums.PathVerbLine:
			if len(pts) >= 2 {
				pt := f32.Pt(float32(pts[1].X), float32(pts[1].Y))
				b.LineTo(pt)
				s.Segments = append(s.Segments, stroke.LineTo(pt))
				current = pt
			}

		case enums.PathVerbQuad:
			if len(pts) >= 3 {
				ctrl := f32.Pt(float32(pts[1].X), float32(pts[1].Y))
				end := f32.Pt(float32(pts[2].X), float32(pts[2].Y))
				// Convert quadratic to cubic: CP1 = current + 2/3*(ctrl - current), CP2 = end + 2/3*(ctrl - end)
				cp1 := f32.Pt(
					current.X+2.0/3.0*(ctrl.X-current.X),
					current.Y+2.0/3.0*(ctrl.Y-current.Y),
				)
				cp2 := f32.Pt(
					end.X+2.0/3.0*(ctrl.X-end.X),
					end.Y+2.0/3.0*(ctrl.Y-end.Y),
				)
				b.CubeTo(cp1, cp2, end)
				s.Segments = append(s.Segments, stroke.QuadTo(ctrl, end))
				current = end
			}

		case enums.PathVerbConic:
			if len(pts) >= 3 {
				ctrl := f32.Pt(float32(pts[1].X), float32(pts[1].Y))
				end := f32.Pt(float32(pts[2].X), float32(pts[2].Y))
				weight := rec.ConicWeight
				if weight < 0 {
					weight = 1.0 // Default weight
				}
				// Convert conic to cubic using the weight
				// For now, treat as quadratic (weight=1) - can be enhanced later
				cp1 := f32.Pt(
					current.X+2.0/3.0*(ctrl.X-current.X),
					current.Y+2.0/3.0*(ctrl.Y-current.Y),
				)
				cp2 := f32.Pt(
					end.X+2.0/3.0*(ctrl.X-end.X),
					end.Y+2.0/3.0*(ctrl.Y-end.Y),
				)
				b.CubeTo(cp1, cp2, end)
				s.Segments = append(s.Segments, stroke.QuadTo(ctrl, end))
				current = end
			}

		case enums.PathVerbCubic:
			if len(pts) >= 4 {
				c1 := f32.Pt(float32(pts[1].X), float32(pts[1].Y))
				c2 := f32.Pt(float32(pts[2].X), float32(pts[2].Y))
				end := f32.Pt(float32(pts[3].X), float32(pts[3].Y))
				b.CubeTo(c1, c2, end)
				s.Segments = append(s.Segments, stroke.CubeTo(c1, c2, end))
				current = end
			}

		case enums.PathVerbClose:
			if hasStart {
				b.LineTo(start)
				s.Segments = append(s.Segments, stroke.LineTo(start))
				current = start
			}
		}
	}

	pathSpec := b.End()

	if internalPaint.Fill {
		clipSave := clip.Outline{Path: pathSpec}.Op().Push(c.ops)
		gpaint.ColorOp{Color: internalPaint.Color}.Add(c.ops)
		gpaint.PaintOp{}.Add(c.ops)
		clipSave.Pop()
	} else {
		contours := stroke.StrokedContours(s, internalPaint.Stroke)
		var stroked clip.Path
		stroked.Begin(c.ops)
		for _, contour := range contours {
			for i, seg := range contour {
				if i == 0 {
					stroked.MoveTo(f32.Point(seg.Start))
				}
				stroked.CubeTo(f32.Point(seg.CP1), f32.Point(seg.CP2), f32.Point(seg.End))
			}
		}
		strokePathSpec := stroked.End()
		gpaint.FillShape(c.ops, internalPaint.Color, clip.Outline{Path: strokePathSpec}.Op())
	}
}

// DrawPath implements SkCanvas.DrawPath - matches SkCanvas signature.
func (c *canvas) DrawPath(path SkPath, paint SkPaint) {
	c.drawPathInternal(path, paint)
}

// ── State Management (additional methods) ───────────────────────────────────

func (c *canvas) SaveLayer(bounds *models.Rect, paint SkPaint) int {
	// For Gio, SaveLayer is equivalent to Save since Gio handles layers implicitly
	// TODO: Implement proper layer support with bounds and paint alpha if needed
	return c.Save()
}

func (c *canvas) RestoreToCount(saveCount int) {
	for len(c.stack) > saveCount && len(c.stack) > 1 {
		c.Restore()
	}
}

func (c *canvas) GetSaveCount() int {
	return len(c.stack)
}

func (c *canvas) Skew(sx, sy Scalar) {
	matrix := impl.NewMatrixSkew(sx, sy)
	c.Concat(matrix)
}

// ── Drawing Primitives ───────────────────────────────────────────────────

func (c *canvas) DrawPaint(paint SkPaint) {
	// Fill the entire clip region
	// Use a very large rectangle to approximate "infinite" canvas
	internalPaint := skPaintToPaint(paint)
	transformSave := op.Affine(c.stack[len(c.stack)-1].xform).Push(c.ops)
	defer transformSave.Pop()

	// Draw a full-coverage paint (relies on clip to bound it)
	gpaint.ColorOp{Color: internalPaint.Color}.Add(c.ops)
	gpaint.PaintOp{}.Add(c.ops)
}

func (c *canvas) DrawRect(rect interfaces.Rect, paint SkPaint) {
	path := impl.NewSkPath(enums.PathFillTypeWinding)
	path.AddRect(rect, enums.PathDirectionCW, 0)
	c.DrawPath(path, paint)
}

func (c *canvas) DrawRRect(rrect interfaces.RRect, paint SkPaint) {
	path := impl.NewSkPath(enums.PathFillTypeWinding)
	path.AddRRect(rrect, enums.PathDirectionCW)
	c.DrawPath(path, paint)
}

func (c *canvas) DrawDRRect(outer interfaces.RRect, inner interfaces.RRect, paint SkPaint) {
	// Draw "donut" - outer minus inner using even-odd fill
	path := impl.NewSkPath(enums.PathFillTypeEvenOdd)
	path.AddRRect(outer, enums.PathDirectionCW)
	path.AddRRect(inner, enums.PathDirectionCCW) // Counter-clockwise for hole
	c.DrawPath(path, paint)
}

func (c *canvas) DrawOval(oval interfaces.Rect, paint SkPaint) {
	path := impl.NewSkPath(enums.PathFillTypeWinding)
	path.AddOval(oval, enums.PathDirectionCW)
	c.DrawPath(path, paint)
}

func (c *canvas) DrawArc(oval interfaces.Rect, startAngle, sweepAngle Scalar, useCenter bool, paint SkPaint) {
	if sweepAngle == 0 {
		return
	}
	path := impl.NewSkPath(enums.PathFillTypeWinding)
	if useCenter {
		// Wedge: start from center
		cx := (oval.Left + oval.Right) / 2
		cy := (oval.Top + oval.Bottom) / 2
		path.MoveTo(cx, cy)
		path.ArcTo(oval, startAngle, sweepAngle, false)
		path.Close()
	} else {
		// Arc only
		path.AddArc(oval, startAngle, sweepAngle)
	}
	c.DrawPath(path, paint)
}

func (c *canvas) DrawCircle(center interfaces.Point, radius Scalar, paint SkPaint) {
	path := impl.NewSkPath(enums.PathFillTypeWinding)
	path.AddCircle(center.X, center.Y, radius, enums.PathDirectionCW)
	c.DrawPath(path, paint)
}

func (c *canvas) DrawPoints(mode enums.PointMode, points []interfaces.Point, paint SkPaint) {
	if len(points) < 1 {
		return
	}

	path := impl.NewSkPath(enums.PathFillTypeWinding)
	switch mode {
	case enums.PointModePoints:
		// Draw each point as a small circle/square based on cap
		strokeWidth := paint.GetStrokeWidth()
		if strokeWidth <= 0 {
			strokeWidth = 1
		}
		for _, pt := range points {
			path.AddCircle(pt.X, pt.Y, strokeWidth/2, enums.PathDirectionCW)
		}
	case enums.PointModeLines:
		// Draw pairs as line segments
		for i := 0; i+1 < len(points); i += 2 {
			path.MoveTo(points[i].X, points[i].Y)
			path.LineTo(points[i+1].X, points[i+1].Y)
		}
	case enums.PointModePolygon:
		// Draw connected line segments
		if len(points) > 0 {
			path.MoveTo(points[0].X, points[0].Y)
			for i := 1; i < len(points); i++ {
				path.LineTo(points[i].X, points[i].Y)
			}
		}
	}
	c.DrawPath(path, paint)
}

func (c *canvas) DrawLine(p0, p1 interfaces.Point, paint SkPaint) {
	c.DrawPoints(enums.PointModeLines, []interfaces.Point{p0, p1}, paint)
}

// ── Image Drawing ───────────────────────────────────────────────────

func (c *canvas) DrawImage(image interfaces.SkImage, left, top Scalar, paint SkPaint) {
	// TODO: Implement image drawing using Gio's paint.NewImageOp
	// Requires converting SkImage to image.Image
	_ = image
	_ = left
	_ = top
	_ = paint
}

func (c *canvas) DrawImageRect(image interfaces.SkImage, src *interfaces.Rect, dst interfaces.Rect, paint SkPaint) {
	// TODO: Implement image rect drawing with scaling
	_ = image
	_ = src
	_ = dst
	_ = paint
}

// ── Clipping ───────────────────────────────────────────────────

func (c *canvas) ClipRect(rect interfaces.Rect, clipOp enums.ClipOp, doAntiAlias bool) {
	// TODO: Implement clipping - store clip state and apply during drawing
	// For now, Gio's clip operations are applied at draw time
	_ = rect
	_ = clipOp
	_ = doAntiAlias
}

func (c *canvas) ClipRRect(rrect interfaces.RRect, clipOp enums.ClipOp, doAntiAlias bool) {
	// TODO: Implement RRect clipping
	_ = rrect
	_ = clipOp
	_ = doAntiAlias
}

func (c *canvas) ClipPath(path SkPath, clipOp enums.ClipOp, doAntiAlias bool) {
	// TODO: Implement path clipping
	_ = path
	_ = clipOp
	_ = doAntiAlias
}

// ── Text Drawing ───────────────────────────────────────────────────

func (c *canvas) DrawTextBlob(blob interfaces.SkTextBlob, x, y Scalar, paint SkPaint) {
	// TODO: Implement using paragraph/shaper modules
	_ = blob
	_ = x
	_ = y
	_ = paint
}

func (c *canvas) DrawSimpleText(text []byte, encoding enums.TextEncoding, x, y Scalar, font interfaces.SkFont, paint SkPaint) {
	// TODO: Shape text using shaper, create blob, draw
	_ = text
	_ = encoding
	_ = x
	_ = y
	_ = font
	_ = paint
}

func (c *canvas) DrawString(str string, x, y Scalar, font interfaces.SkFont, paint SkPaint) {
	c.DrawSimpleText([]byte(str), enums.TextEncodingUTF8, x, y, font, paint)
}
