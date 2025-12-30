package main

import (
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
	"github.com/zodimo/go-skia-support/skia/shaper"
)

// makeTextBlob creates a text blob using the HarfBuzz shaper
func makeTextBlob(text string, font interfaces.SkFont) interfaces.SkTextBlob {
	hbShaper := shaper.NewHarfbuzzShaper()
	handler := shaper.NewTextBlobBuilderRunHandler(text, models.Point{X: 0, Y: 0})
	hbShaper.Shape(text, font, true, 0, handler, nil)
	return handler.MakeBlob()
}
