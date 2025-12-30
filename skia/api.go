// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
	"github.com/zodimo/go-skia-support/skia/shaper"
)

// ── Type aliases for better developer experience ────────────────────────────
// These aliases allow clients to use skia.Scalar, skia.SkPath, etc.
// without needing to import the underlying packages.

// ── Shaper Types ────────────────────────────────────────────────────────
// These types allow using the SkShaper API directly from the skia package.

// Shaper is the interface for text shaping.
type Shaper = shaper.Shaper

// RunIterator is the base interface for iterators over runs of text.
type RunIterator = shaper.RunIterator

// FontRunIterator iterates over runs of fonts.
type FontRunIterator = shaper.FontRunIterator

// BiDiRunIterator iterates over runs of bidirectional levels.
type BiDiRunIterator = shaper.BiDiRunIterator

// ScriptRunIterator iterates over runs of scripts.
type ScriptRunIterator = shaper.ScriptRunIterator

// LanguageRunIterator iterates over runs of languages.
type LanguageRunIterator = shaper.LanguageRunIterator

// RunHandler is the interface for handling the results of text shaping.
type RunHandler = shaper.RunHandler

// RunInfo contains information about a shaped run.
type RunInfo = shaper.RunInfo

// Buffer contains the shaped glyphs and positions for a run.
type Buffer = shaper.Buffer

// Feature represents an OpenType feature.
type Feature = shaper.Feature

// Range represents a range of indices in the text.
type Range = shaper.Range

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
type Canvas = interfaces.SkCanvas
