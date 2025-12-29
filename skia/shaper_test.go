package skia_test

import (
	"testing"

	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/interfaces"
)

// MockSkFont for testing
type MockSkFont struct {
	interfaces.SkFont
}

func TestTrivialIterators(t *testing.T) {
	textLength := 10

	// Test FontRunIterator
	font := &MockSkFont{} // Just need a type that implements the interface.
	// We can pass the font.
	fontIter := skia.NewTrivialFontRunIterator(font, textLength)
	if fontIter == nil {
		t.Error("NewTrivialFontRunIterator returned nil")
	}
	if fontIter.CurrentFont() != font {
		t.Error("CurrentFont() did not return the expected font")
	}

	// Test BiDiRunIterator
	bidiIter := skia.NewTrivialBiDiRunIterator(skia.BidiLTR, textLength)
	if bidiIter == nil {
		t.Error("NewTrivialBiDiRunIterator returned nil")
	}
	if bidiIter.CurrentLevel() != skia.BidiLTR {
		t.Errorf("Expected level %d, got %d", skia.BidiLTR, bidiIter.CurrentLevel())
	}

	// Test ScriptRunIterator
	scriptIter := skia.NewTrivialScriptRunIterator(0, textLength)
	if scriptIter == nil {
		t.Error("NewTrivialScriptRunIterator returned nil")
	}

	// Test LanguageRunIterator
	langIter := skia.NewTrivialLanguageRunIterator("en-US", textLength)
	if langIter == nil {
		t.Error("NewTrivialLanguageRunIterator returned nil")
	}
	if langIter.CurrentLanguage() != "en-US" {
		t.Errorf("Expected language %s, got %s", "en-US", langIter.CurrentLanguage())
	}
}
