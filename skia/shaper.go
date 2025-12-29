// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/shaper"
)

// ── Trivial Iterators ────────────────────────────────────────────────────────
// These constructors provide simple single-run iterators for basic shaping needs.

const (
	// BidiLTR represents a Left-To-Right BiDi level (0).
	BidiLTR uint8 = 0
	// BidiRTL represents a Right-To-Left BiDi level (1).
	BidiRTL uint8 = 1
)

// NewTrivialFontRunIterator creates a trivial FontRunIterator that assumes the
// entire text uses the same font.
func NewTrivialFontRunIterator(font interfaces.SkFont, textLength int) FontRunIterator {
	return shaper.NewTrivialFontRunIterator(font, textLength)
}

// NewTrivialBiDiRunIterator creates a trivial BiDiRunIterator that assumes the
// entire text has the same BiDi level.
func NewTrivialBiDiRunIterator(bidiLevel uint8, textLength int) BiDiRunIterator {
	return shaper.NewTrivialBiDiRunIterator(bidiLevel, textLength)
}

// NewTrivialScriptRunIterator creates a trivial ScriptRunIterator that assumes the
// entire text uses the same script.
func NewTrivialScriptRunIterator(script uint32, textLength int) ScriptRunIterator {
	return shaper.NewTrivialScriptRunIterator(script, textLength)
}

// NewTrivialLanguageRunIterator creates a trivial LanguageRunIterator that assumes the
// entire text uses the same language.
func NewTrivialLanguageRunIterator(language string, textLength int) LanguageRunIterator {
	return shaper.NewTrivialLanguageRunIterator(language, textLength)
}
