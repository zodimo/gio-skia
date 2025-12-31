package skia

import (
	"image/color"

	"github.com/zodimo/go-skia-support/skia/models"
)

// ColorToColor4f converts a color.NRGBA to a models.Color4f.
func ColorToColor4f(col color.NRGBA) models.Color4f {
	return models.Color4f{
		R: Scalar(float32(col.R) / 255.0),
		G: Scalar(float32(col.G) / 255.0),
		B: Scalar(float32(col.B) / 255.0),
		A: Scalar(float32(col.A) / 255.0),
	}
}
