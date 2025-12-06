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

type Paint struct {
	Color  color.NRGBA
	Stroke StrokeOpts
	Fill   bool
}

type path struct {
	// Build both representations simultaneously
	verbs []pathOp
}

type pathOp struct {
	verb uint8
	pts  [6]float32
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

// ── Path construction ─────────────────────────────────────────────────

func NewPath() Path {
	return &path{}
}

func (p *path) MoveTo(x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 0, pts: [6]float32{x, y}})
}

func (p *path) LineTo(x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 1, pts: [6]float32{x, y}})
}

func (p *path) QuadTo(cx, cy, x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 2, pts: [6]float32{cx, cy, x, y}})
}

func (p *path) CubeTo(cx1, cy1, cx2, cy2, x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 3, pts: [6]float32{cx1, cy1, cx2, cy2, x, y}})
}

func (p *path) Close() {
	p.verbs = append(p.verbs, pathOp{verb: 4})
}

func (p *path) AddRect(x, y, w, h float32) {
	p.MoveTo(x, y)
	p.LineTo(x+w, y)
	p.LineTo(x+w, y+h)
	p.LineTo(x, y+h)
	p.Close()
}

func (p *path) AddCircle(cx, cy, r float32) {
	const k = 0.5522848
	p.MoveTo(cx+r, cy)
	p.CubeTo(cx+r, cy+k*r, cx+k*r, cy+r, cx, cy+r)
	p.CubeTo(cx-k*r, cy+r, cx-r, cy+k*r, cx-r, cy)
	p.CubeTo(cx-r, cy-k*r, cx-k*r, cy-r, cx, cy-r)
	p.CubeTo(cx+k*r, cy-r, cx+r, cy-k*r, cx+r, cy)
}

func (p *path) unwrap() []pathOp {
	return p.verbs
}

// ── Drawing ───────────────────────────────────────────────────────────

func (c *canvas) DrawPath(p Path) {
	save := op.TransformOp{}.Push(c.ops)
	xform := c.stack[len(c.stack)-1].xform
	op.Affine(xform).Add(c.ops)

	// Build clip.Path
	var b clip.Path
	b.Begin(c.ops)

	// Build stroke.Path in parallel
	var s stroke.Path

	for _, cmd := range p.unwrap() {
		switch cmd.verb {
		case 0:
			pt := f32.Pt(cmd.pts[0], cmd.pts[1])
			b.MoveTo(pt)
			s.Segments = append(s.Segments, stroke.MoveTo(pt))
		case 1:
			pt := f32.Pt(cmd.pts[0], cmd.pts[1])
			b.LineTo(pt)
			s.Segments = append(s.Segments, stroke.LineTo(pt))
		case 2:
			ctrl := f32.Pt(cmd.pts[0], cmd.pts[1])
			end := f32.Pt(cmd.pts[2], cmd.pts[3])
			b.QuadTo(ctrl, end)
			s.Segments = append(s.Segments, stroke.QuadTo(ctrl, end))
		case 3:
			c1 := f32.Pt(cmd.pts[0], cmd.pts[1])
			c2 := f32.Pt(cmd.pts[2], cmd.pts[3])
			end := f32.Pt(cmd.pts[4], cmd.pts[5])
			b.CubeTo(c1, c2, end)
			s.Segments = append(s.Segments, stroke.CubeTo(c1, c2, end))
		case 4:
			b.Close()
			// No explicit close needed for stroke.Path
		}
	}

	pathSpec := b.End()
	paintCfg := c.stack[len(c.stack)-1].paint

	if paintCfg.Fill {
		stack := clip.Outline{Path: pathSpec}.Op().Push(c.ops)
		paint.ColorOp{Color: paintCfg.Color}.Add(c.ops)
		paint.PaintOp{}.Add(c.ops)
		stack.Pop()
	} else {
		// Use expandStroke to get stroked path
		strokePathSpec := stroke.ExpandStroke(s, paintCfg.Stroke.Width,
			paintCfg.Stroke.Join, paintCfg.Stroke.Cap,
			paintCfg.Stroke.Miter, paintCfg.Stroke.Dash, paintCfg.Stroke.Dash0)

		stack := clip.Outline{Path: strokePathSpec}.Op().Push(c.ops)
		paint.ColorOp{Color: paintCfg.Color}.Add(c.ops)
		paint.PaintOp{}.Add(c.ops)
		stack.Pop()
	}

	save.Pop()
}
