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
