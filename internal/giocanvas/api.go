package giocanvas

import (
	"image"
	"image/color"

	"gioui.org/op"
)

type Canvas interface {
	AbsArc(x float32, y float32, radius float32, start float64, end float64, fillcolor color.NRGBA)
	AbsCenterImage(name string, x float32, y float32, w int, h int, scale float32)
	AbsCenterRect(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	AbsCircle(x float32, y float32, radius float32, fillcolor color.NRGBA)
	AbsCubicBezier(x float32, y float32, cx1 float32, cy1 float32, cx2 float32, cy2 float32, ex float32, ey float32, size float32, fillcolor color.NRGBA)
	AbsEllipse(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	AbsGrid(width float32, height float32, size float32, interval float32, fillcolor color.NRGBA)
	AbsHLine(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	AbsImg(im image.Image, x float32, y float32, w int, h int, scale float32)
	AbsLine(x0 float32, y0 float32, x1 float32, y1 float32, size float32, fillcolor color.NRGBA)
	AbsPolygon(x []float32, y []float32, fillcolor color.NRGBA)
	AbsQuadBezier(x float32, y float32, cx float32, cy float32, ex float32, ey float32, size float32, fillcolor color.NRGBA)
	AbsRect(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	AbsRotate(x float32, y float32, angle float32) op.TransformStack
	AbsScale(x float32, y float32, factor float32) op.TransformStack
	AbsShear(x float32, y float32, ax float32, ay float32) op.TransformStack
	AbsStrokedCubicBezier(x float32, y float32, cx1 float32, cy1 float32, cx2 float32, cy2 float32, ex float32, ey float32, size float32, strokecolor color.NRGBA)
	AbsStrokedQuadBezier(x float32, y float32, cx float32, cy float32, ex float32, ey float32, size float32, strokecolor color.NRGBA)
	AbsText(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	AbsTextEnd(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	AbsTextMid(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	AbsTextWrap(x float32, y float32, size float32, width float32, s string, fillcolor color.NRGBA)
	AbsTranslate(x float32, y float32) op.TransformStack
	AbsVLine(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	Arc(x float32, y float32, r float32, a1 float64, a2 float64, fillcolor color.NRGBA)
	ArcLine(x float32, y float32, r float32, a1 float64, a2 float64, size float32, fillcolor color.NRGBA)
	Background(fillcolor color.NRGBA)
	CText(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	CenterImage(name string, x float32, y float32, w int, h int, scale float32)
	CenterRect(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	Circle(x float32, y float32, r float32, fillcolor color.NRGBA)
	Coord(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	CornerRect(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	CubeCurve(x float32, y float32, cx1 float32, cy1 float32, cx2 float32, cy2 float32, ex float32, ey float32, fillcolor color.NRGBA)
	CubeStrokedCurve(x float32, y float32, cx1 float32, cy1 float32, cx2 float32, cy2 float32, ex float32, ey float32, size float32, fillcolor color.NRGBA)
	Curve(x float32, y float32, cx float32, cy float32, ex float32, ey float32, fillcolor color.NRGBA)
	EText(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	Ellipse(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	Grid(x float32, y float32, w float32, h float32, size float32, interval float32, linecolor color.NRGBA)
	HLine(x float32, y float32, linewidth float32, size float32, linecolor color.NRGBA)
	Image(name string, x float32, y float32, w int, h int, scale float32)
	Img(im image.Image, x float32, y float32, w int, h int, scale float32)
	Line(x0 float32, y0 float32, x1 float32, y1 float32, size float32, strokecolor color.NRGBA)
	Polar(cx float32, cy float32, r float32, theta float32) (float32, float32)
	PolarDegrees(cx float32, cy float32, r float32, theta float32) (float32, float32)
	Polygon(x []float32, y []float32, fillcolor color.NRGBA)
	QuadCurve(x float32, y float32, cx float32, cy float32, ex float32, ey float32, fillcolor color.NRGBA)
	QuadStrokedCurve(x float32, y float32, cx float32, cy float32, ex float32, ey float32, size float32, strokecolor color.NRGBA)
	Rect(x float32, y float32, w float32, h float32, fillcolor color.NRGBA)
	Rotate(x float32, y float32, angle float32) op.TransformStack
	Scale(x float32, y float32, factor float32) op.TransformStack
	Shear(x float32, y float32, ax float32, ay float32) op.TransformStack
	Square(x float32, y float32, w float32, fillcolor color.NRGBA)
	StrokedCubeCurve(x float32, y float32, cx1 float32, cy1 float32, cx2 float32, cy2 float32, ex float32, ey float32, size float32, strokecolor color.NRGBA)
	StrokedCurve(x float32, y float32, cx float32, cy float32, ex float32, ey float32, size float32, fillcolor color.NRGBA)
	Text(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	TextEnd(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	TextMid(x float32, y float32, size float32, s string, fillcolor color.NRGBA)
	TextWrap(x float32, y float32, size float32, width float32, s string, fillcolor color.NRGBA)
	Translate(x float32, y float32) op.TransformStack
	VLine(x float32, y float32, lineheight float32, size float32, linecolor color.NRGBA)
}
