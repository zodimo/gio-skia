package shaper

import (
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/shaper"
)

// UseGoTextFace is an alias to the interface in go-skia-support,
// allowing consumers to use it without importing support package directly if they prefer.
type UseGoTextFace = shaper.UseGoTextFace

// ShaperImpl wraps the HarfbuzzShaper implementation from go-skia-support.
type ShaperImpl struct {
	*shaper.HarfbuzzShaper
}

// NewShaper creates a new ShaperImpl that delegates to go-skia-support.
func NewShaper() *ShaperImpl {
	return &ShaperImpl{
		HarfbuzzShaper: shaper.NewHarfbuzzShaper(),
	}
}

// Shape shapes the text using the font and runHandler.
// It matches the interface expected by gio-skia consumers, delegating to the support shaper.
func (s *ShaperImpl) Shape(text string, font interfaces.SkFont, leftToRight bool, width float32, runHandler shaper.RunHandler) {
	// Pass nil for descriptors/features as default
	s.HarfbuzzShaper.Shape(text, font, leftToRight, width, runHandler, nil)
}
