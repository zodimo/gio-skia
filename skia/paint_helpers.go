// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"image/color"

	"github.com/zodimo/gio-skia/pkg/stroke"
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
)

// paintAdapter adapts *impl.Paint to interfaces.SkPaint
// This is needed because *impl.Paint.Equals takes *impl.Paint, not interfaces.SkPaint
type paintAdapter struct {
	*impl.Paint
}

var _ interfaces.SkPaint = (*paintAdapter)(nil)

func (a *paintAdapter) Equals(other interfaces.SkPaint) bool {
	if other == nil {
		return false
	}
	// Try to get the underlying *impl.Paint
	if otherAdapter, ok := other.(*paintAdapter); ok {
		return a.Paint.Equals(otherAdapter.Paint)
	}
	// Other implementations of SkPaint can't be compared directly
	return false
}

// NewPaint creates a new SkPaint with default values.
// This is a convenience wrapper around impl.NewPaint().
func NewPaint() SkPaint {
	return &paintAdapter{Paint: impl.NewPaint()}
}

// NewPaintWithColor creates a new SkPaint with the specified color.
// Style defaults to Fill.
func NewPaintWithColor(col color.NRGBA) SkPaint {
	paint := impl.NewPaint()
	color4f := models.Color4f{
		R: Scalar(float32(col.R) / 255.0),
		G: Scalar(float32(col.G) / 255.0),
		B: Scalar(float32(col.B) / 255.0),
		A: Scalar(float32(col.A) / 255.0),
	}
	paint.SetColor(color4f)
	return &paintAdapter{Paint: paint}
}

// NewPaintFill creates a new SkPaint configured for filling with the specified color.
func NewPaintFill(col color.NRGBA) SkPaint {
	paint := NewPaintWithColor(col)
	paint.SetStyle(enums.PaintStyleFill)
	return paint
}

// NewPaintStroke creates a new SkPaint configured for stroking with the specified color and width.
func NewPaintStroke(col color.NRGBA, width float32) SkPaint {
	paint := NewPaintWithColor(col)
	paint.SetStyle(enums.PaintStyleStroke)
	paint.SetStrokeWidth(Scalar(width))
	return paint
}

// ConfigureStrokePaint configures a SkPaint with stroke options.
// This helper converts StrokeOpts to SkPaint settings.
func ConfigureStrokePaint(paint SkPaint, opts stroke.StrokeOpts) SkPaint {
	paint.SetStrokeWidth(Scalar(opts.Width))
	paint.SetStrokeMiter(Scalar(opts.Miter))
	
	// Convert cap style
	switch opts.Cap {
	case stroke.RoundCap:
		paint.SetStrokeCap(enums.PaintCapRound)
	case stroke.SquareCap:
		paint.SetStrokeCap(enums.PaintCapSquare)
	case stroke.FlatCap:
		paint.SetStrokeCap(enums.PaintCapButt)
	case stroke.TriangularCap:
		// TriangularCap has no direct Skia equivalent, use RoundCap
		paint.SetStrokeCap(enums.PaintCapRound)
	}
	
	// Convert join style
	switch opts.Join {
	case stroke.RoundJoin:
		paint.SetStrokeJoin(enums.PaintJoinRound)
	case stroke.MiterJoin:
		paint.SetStrokeJoin(enums.PaintJoinMiter)
	case stroke.BevelJoin:
		paint.SetStrokeJoin(enums.PaintJoinBevel)
	}
	
	// TODO: Handle Dash pattern if go-skia-support implements PathEffect
	
	return paint
}

// PaintStyle aliases for convenience
const (
	PaintStyleFill            = enums.PaintStyleFill
	PaintStyleStroke          = enums.PaintStyleStroke
	PaintStyleStrokeAndFill   = enums.PaintStyleStrokeAndFill
)

// PaintCap aliases for convenience
const (
	PaintCapButt  = enums.PaintCapButt
	PaintCapRound = enums.PaintCapRound
	PaintCapSquare = enums.PaintCapSquare
)

// PaintJoin aliases for convenience
const (
	PaintJoinMiter = enums.PaintJoinMiter
	PaintJoinRound = enums.PaintJoinRound
	PaintJoinBevel = enums.PaintJoinBevel
)
