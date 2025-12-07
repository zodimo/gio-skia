// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"image/color"
	"math"

	"gioui.org/f32"
	"github.com/zodimo/gio-skia/pkg/stroke"
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/interfaces"
)

// affine2DToSkMatrix converts f32.Affine2D to SkMatrix.
// f32.Affine2D stores a 2D affine transformation matrix.
// We extract the 6 affine values and create a 3x3 matrix.
func affine2DToSkMatrix(affine f32.Affine2D) SkMatrix {
	// f32.Affine2D.Elems() returns (sx, hx, ox, hy, sy, oy)
	// representing: [sx hx ox]
	//                [hy sy oy]
	//                [0  0  1 ]
	// SkMatrix stores as [scaleX, skewX, transX, skewY, scaleY, transY, persp0, persp1, persp2]
	sx, hx, ox, hy, sy, oy := affine.Elems()
	return impl.NewMatrixAll(
		Scalar(sx), Scalar(hx), Scalar(ox),
		Scalar(hy), Scalar(sy), Scalar(oy),
		0, 0, 1,
	)
}

// skMatrixToAffine2D converts SkMatrix to f32.Affine2D.
// Only works for affine matrices (no perspective).
func skMatrixToAffine2D(matrix SkMatrix) f32.Affine2D {
	// Extract the 6 affine values from the 3x3 matrix
	// f32.NewAffine2D takes (sx, hx, ox, hy, sy, oy)
	return f32.NewAffine2D(
		float32(matrix.GetScaleX()),
		float32(matrix.GetSkewX()),
		float32(matrix.GetTranslateX()),
		float32(matrix.GetSkewY()),
		float32(matrix.GetScaleY()),
		float32(matrix.GetTranslateY()),
	)
}

// Paint is an internal type used for rendering.
// It's converted from SkPaint for GioUI rendering.
type Paint struct {
	Color  color.NRGBA
	Stroke stroke.StrokeOpts
	Fill   bool
}

// skPaintToPaint converts SkPaint to our internal Paint type for rendering.
func skPaintToPaint(skPaint interfaces.SkPaint) Paint {
	p := Paint{}

	// Get color
	color4f := skPaint.GetColor()
	p.Color = color.NRGBA{
		R: uint8(color4f.R * 255),
		G: uint8(color4f.G * 255),
		B: uint8(color4f.B * 255),
		A: uint8(color4f.A * 255),
	}

	// Get style
	style := skPaint.GetStyle()
	p.Fill = (style == enums.PaintStyleFill || style == enums.PaintStyleStrokeAndFill)

	if style == enums.PaintStyleStroke || style == enums.PaintStyleStrokeAndFill {
		// Set stroke properties
		p.Stroke.Width = float32(skPaint.GetStrokeWidth())
		p.Stroke.Miter = float32(skPaint.GetStrokeMiter())

		// Convert cap
		cap := skPaint.GetStrokeCap()
		switch cap {
		case enums.PaintCapRound:
			p.Stroke.Cap = stroke.RoundCap
		case enums.PaintCapSquare:
			p.Stroke.Cap = stroke.SquareCap
		case enums.PaintCapButt:
			p.Stroke.Cap = stroke.FlatCap
		}

		// Convert join
		join := skPaint.GetStrokeJoin()
		switch join {
		case enums.PaintJoinRound:
			p.Stroke.Join = stroke.RoundJoin
		case enums.PaintJoinMiter:
			p.Stroke.Join = stroke.MiterJoin
		case enums.PaintJoinBevel:
			p.Stroke.Join = stroke.BevelJoin
		}
	}

	return p
}

// radiansToDegrees converts radians to degrees.
func radiansToDegrees(rad float32) Scalar {
	return Scalar(rad * 180.0 / math.Pi)
}

// degreesToRadians converts degrees to radians.
func degreesToRadians(deg Scalar) float32 {
	return float32(deg) * float32(math.Pi) / 180.0
}
