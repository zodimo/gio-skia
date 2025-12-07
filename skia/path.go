package skia

import (
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
)

// NewPath creates a new path using go-skia-support's implementation.
// Returns SkPath directly for use with DrawPath.
func NewPath() SkPath {
	return impl.NewSkPath(enums.PathFillTypeDefault)
}
