// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/models"
)

// Path helpers provide convenience functions for working with SkPath using float32 coordinates.
// These wrap SkPath methods to accept float32 instead of Scalar for easier usage.

// PathMoveTo moves to the specified point in the path.
func PathMoveTo(path SkPath, x, y float32) {
	path.MoveTo(Scalar(x), Scalar(y))
}

// PathLineTo adds a line to the specified point in the path.
func PathLineTo(path SkPath, x, y float32) {
	path.LineTo(Scalar(x), Scalar(y))
}

// PathQuadTo adds a quadratic bezier curve to the path.
func PathQuadTo(path SkPath, cx, cy, x, y float32) {
	path.QuadTo(Scalar(cx), Scalar(cy), Scalar(x), Scalar(y))
}

// PathCubeTo adds a cubic bezier curve to the path.
func PathCubeTo(path SkPath, cx1, cy1, cx2, cy2, x, y float32) {
	path.CubicTo(Scalar(cx1), Scalar(cy1), Scalar(cx2), Scalar(cy2), Scalar(x), Scalar(y))
}

// PathAddRect adds a rectangle to the path.
func PathAddRect(path SkPath, x, y, w, h float32) {
	rect := models.Rect{
		Left:   Scalar(x),
		Top:    Scalar(y),
		Right:  Scalar(x + w),
		Bottom: Scalar(y + h),
	}
	path.AddRect(rect, enums.PathDirectionDefault, 0)
}

// PathAddCircle adds a circle to the path.
func PathAddCircle(path SkPath, cx, cy, r float32) {
	path.AddCircle(Scalar(cx), Scalar(cy), Scalar(r), enums.PathDirectionDefault)
}
