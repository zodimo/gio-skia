// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"image/color"

	"gioui.org/f32"
	"github.com/zodimo/gio-skia/pkg/stroke"
)

// Canvas defines a Skia-style immediate-mode drawing context.
// All operations are GPU-accelerated via Gio's renderer.
type Canvas interface {
	// ── State management ──────────────────────────────────────────────────
	Save()                 // Push matrix & paint onto stack
	Restore()              // Pop state; restores matrix & paint
	Concat(m f32.Affine2D) // Multiply current transform matrix
	Translate(x, y float32)
	Scale(x, y float32)
	Rotate(angle float32) // radians, counter-clockwise

	// ── Paint state ───────────────────────────────────────────────────────
	SetColor(col color.NRGBA)
	SetStroke(opt StrokeOpts) // Configure stroke style
	Fill()                    // Switch to fill mode (default)
	Stroke()                  // Switch to stroke mode

	// ── Drawing ───────────────────────────────────────────────────────────
	DrawPath(p Path) // Render path with current paint & transform
}

// Path is a concrete path builder. Create with NewPath().
type Path interface { /* unexported */
	MoveTo(x, y float32)
	LineTo(x, y float32)
	QuadTo(cx, cy, x, y float32)             // Quadratic Bézier
	CubeTo(cx1, cy1, cx2, cy2, x, y float32) // Cubic Bézier
	Close()                                  // Close current contour
	AddRect(x, y, w, h float32)
	AddCircle(cx, cy, r float32)
	unwrap() []pathOp
}

// ── Stroke configuration ────────────────────────────────────────────────
type StrokeOpts struct {
	Width float32          // Stroke width in pixels
	Miter float32          // Miter limit (default 4)
	Cap   stroke.CapStyle  // Line cap style
	Join  stroke.JoinStyle // Line join style
	Dash  []float32        // Dash pattern (optional)
	Dash0 float32          // Dash phase
}

const (
	// MiterJoin joins path segments with a sharp corner.
	// It falls back to a bevel join if the miter limit is exceeded.
	MiterJoin = stroke.MiterJoin
	// RoundJoin joins path segments with a round segment.
	RoundJoin = stroke.RoundJoin
	// BevelJoin joins path segments with sharp bevels.
	BevelJoin = stroke.BevelJoin
)
