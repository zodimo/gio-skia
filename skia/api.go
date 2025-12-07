// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
)

// ── Type aliases for better developer experience ────────────────────────────
// These aliases allow clients to use skia.Scalar, skia.SkPath, etc.
// without needing to import the underlying packages.

// Scalar is a floating-point value used for coordinates, dimensions, and angles.
type Scalar = models.Scalar

// SkPath represents a path that can be drawn on a canvas.
// This is an alias for go-skia-support's SkPath interface.
type SkPath = interfaces.SkPath

// SkPaint represents paint properties for drawing operations.
// This is an alias for go-skia-support's SkPaint interface.
type SkPaint = interfaces.SkPaint

// SkMatrix represents a 3x3 transformation matrix.
// This is an alias for go-skia-support's SkMatrix interface.
type SkMatrix = interfaces.SkMatrix

// Canvas defines a Skia-style immediate-mode drawing context.
// All operations are GPU-accelerated via Gio's renderer.
// This interface matches SkCanvas method signatures for the methods we implement,
// allowing us to eventually use SkCanvas directly once all methods are implemented.
type Canvas interface {
	// ── State management ──────────────────────────────────────────────────
	// Save saves the current matrix and clip state to a stack.
	// Returns the save count (depth of stack before this save).
	Save() int
	// Restore removes the most recent save state from the stack.
	Restore()
	// Concat replaces current matrix with matrix premultiplied with existing matrix.
	Concat(matrix SkMatrix)
	// Translate translates the current matrix by dx along x-axis and dy along y-axis.
	Translate(dx, dy Scalar)
	// Scale scales the current matrix by sx on x-axis and sy on y-axis.
	Scale(sx, sy Scalar)
	// Rotate rotates the current matrix by degrees around the origin (0, 0).
	// Positive degrees rotates clockwise.
	Rotate(degrees Scalar)

	// ── Drawing ───────────────────────────────────────────────────────────
	// DrawPath draws a path containing one or more contours.
	// Path geometry is transformed by the current matrix before drawing.
	// Matches SkCanvas.DrawPath signature.
	DrawPath(path SkPath, paint SkPaint)

	// ── Convenience methods (not part of SkCanvas interface) ──────────────
	// Convenience transformation methods that accept float32 instead of Scalar
	TranslateFloat32(x, y float32)
	ScaleFloat32(x, y float32)
	RotateFloat32(degrees float32) // degrees, not radians
}

// Path helper functions for convenience when working with SkPath.
// NewPath() now returns SkPath directly, but these helpers make it easier
// to work with float32 coordinates instead of Scalar.
