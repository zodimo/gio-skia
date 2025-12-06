// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/pkg/stroke"
)

type canvas struct {
	ops   *op.Ops
	stack []context
}

type context struct {
	xform f32.Affine2D
	paint Paint
}

// NewCanvas returns a Canvas implementation backed by Gio's GPU renderer.
func NewCanvas(ops *op.Ops) Canvas {
	return &canvas{
		ops: ops,
		stack: []context{{
			xform: f32.Affine2D{},
			paint: Paint{
				Color: color.NRGBA{A: 255},
				Fill:  true,
			},
		}},
	}
}

// ── State management ───────────────────────────────────────────────────

func (c *canvas) Save() {
	top := c.stack[len(c.stack)-1]
	c.stack = append(c.stack, top)
}

func (c *canvas) Restore() {
	if len(c.stack) > 1 {
		c.stack = c.stack[:len(c.stack)-1]
	}
}

func (c *canvas) Concat(m f32.Affine2D) {
	top := &c.stack[len(c.stack)-1]
	top.xform = top.xform.Mul(m)
}

func (c *canvas) Translate(x, y float32) {
	c.Concat(f32.Affine2D{}.Offset(f32.Pt(x, y)))
}

func (c *canvas) Scale(x, y float32) {
	c.Concat(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(x, y)))
}

func (c *canvas) Rotate(angle float32) {
	c.Concat(f32.Affine2D{}.Rotate(f32.Pt(0, 0), angle))
}

// ── Paint state ───────────────────────────────────────────────────────

func (c *canvas) SetColor(col color.NRGBA) {
	c.stack[len(c.stack)-1].paint.Color = col
}

func (c *canvas) SetStroke(opt StrokeOpts) {
	top := &c.stack[len(c.stack)-1]
	top.paint.Stroke = opt
	top.paint.Fill = false
}

func (c *canvas) Fill() {
	c.stack[len(c.stack)-1].paint.Fill = true
}

func (c *canvas) Stroke() {
	c.stack[len(c.stack)-1].paint.Fill = false
}

func (c *canvas) DrawPath(p Path) {
	transformSave := op.Affine(c.stack[len(c.stack)-1].xform).Push(c.ops)
	defer transformSave.Pop()

	// Build paths
	var b clip.Path
	b.Begin(c.ops)
	// Build stroke.Path in parallel
	var s stroke.Path
	var start f32.Point
	for _, cmd := range p.unwrap() {
		switch cmd.verb {
		case 0: // MoveTo
			pt := f32.Pt(cmd.pts[0], cmd.pts[1])
			s.Segments = append(s.Segments, stroke.MoveTo(pt))
			start = pt
		case 1: // LineTo
			pt := f32.Pt(cmd.pts[0], cmd.pts[1])
			s.Segments = append(s.Segments, stroke.LineTo(pt))
		case 2: // QuadTo
			ctrl := f32.Pt(cmd.pts[0], cmd.pts[1])
			end := f32.Pt(cmd.pts[2], cmd.pts[3])
			s.Segments = append(s.Segments, stroke.QuadTo(ctrl, end))
		case 3: // CubeTo
			c1 := f32.Pt(cmd.pts[0], cmd.pts[1])
			c2 := f32.Pt(cmd.pts[2], cmd.pts[3])
			end := f32.Pt(cmd.pts[4], cmd.pts[5])
			s.Segments = append(s.Segments, stroke.CubeTo(c1, c2, end))
		case 4: // Close
			s.Segments = append(s.Segments, stroke.LineTo(start))
		}
	}

	pathSpec := b.End()
	paintCfg := c.stack[len(c.stack)-1].paint

	if paintCfg.Fill {
		clipSave := clip.Outline{Path: pathSpec}.Op().Push(c.ops)
		paint.ColorOp{Color: paintCfg.Color}.Add(c.ops)
		paint.PaintOp{}.Add(c.ops)
		clipSave.Pop()
	} else {
		contours := stroke.StrokedContours(s, paintCfg.Stroke)
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
		paint.FillShape(c.ops, paintCfg.Color, clip.Outline{Path: strokePathSpec}.Op())
	}
}
