package shaper

import (
	"bytes"
	"testing"

	"github.com/go-text/typesetting/font"
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"

	support_shaper "github.com/zodimo/go-skia-support/skia/shaper"
	"golang.org/x/image/font/gofont/goregular"
)

// MockTypeface implements interfaces.SkTypeface and UseGoTextFace
type MockTypeface struct {
	interfaces.SkTypeface
	face *font.Face
}

// GoTextFace implements UseGoTextFace interface from go-skia-support
func (m *MockTypeface) GoTextFace() *font.Face {
	return m.face
}

// MockFont implements interfaces.SkFont
type MockFont struct {
	interfaces.SkFont
	typeface *MockTypeface
	size     float32
}

func (m *MockFont) Typeface() interfaces.SkTypeface {
	return m.typeface
}

func (m *MockFont) Size() models.Scalar {
	return models.Scalar(m.size)
}

// MockRunHandler captures calls
type MockRunHandler struct {
	BeginLineCalled     bool
	RunInfoCalled       bool
	CommitRunInfoCalled bool
	CommitBufferCalled  bool
	CommitLineCalled    bool
	Infos               []support_shaper.RunInfo
}

func (m *MockRunHandler) BeginLine() {
	m.BeginLineCalled = true
}

func (m *MockRunHandler) RunInfo(info support_shaper.RunInfo) {
	m.RunInfoCalled = true
	m.Infos = append(m.Infos, info)
}

func (m *MockRunHandler) CommitRunInfo() {
	m.CommitRunInfoCalled = true
}

func (m *MockRunHandler) RunBuffer(info support_shaper.RunInfo) support_shaper.Buffer {
	return support_shaper.Buffer{
		Glyphs:    make([]uint16, info.GlyphCount),
		Positions: make([]models.Point, info.GlyphCount),
		Offsets:   make([]models.Point, info.GlyphCount),
		Clusters:  make([]uint32, info.GlyphCount),
	}
}

func (m *MockRunHandler) CommitRunBuffer(info support_shaper.RunInfo) {
	m.CommitBufferCalled = true
}

func (m *MockRunHandler) CommitLine() {
	m.CommitLineCalled = true
}

func TestShaper_Shape_Wrapper(t *testing.T) {
	// Parse a real font to avoid panic
	face, err := font.ParseTTF(bytes.NewReader(goregular.TTF))
	if err != nil {
		t.Fatalf("Failed to parse goregular font: %v", err)
	}

	mockTf := &MockTypeface{face: face}
	mockFont := &MockFont{typeface: mockTf, size: 12.0}

	s := NewShaper()
	handler := &MockRunHandler{}

	// Act
	s.Shape("Hello", mockFont, true, 100.0, handler)

	// Assert
	if !handler.BeginLineCalled {
		t.Error("BeginLine not called")
	}
	if !handler.RunInfoCalled {
		t.Error("RunInfo not called")
	}
	if !handler.CommitRunInfoCalled {
		t.Error("CommitRunInfo not called")
	}
	if !handler.CommitBufferCalled {
		t.Error("CommitRunBuffer not called")
	}
	if !handler.CommitLineCalled {
		t.Error("CommitLine not called")
	}

	if len(handler.Infos) == 0 {
		t.Fatal("No RunInfos generated")
	}
	info := handler.Infos[0]
	if info.GlyphCount == 0 {
		t.Error("GlyphCount is 0, expected > 0 for 'Hello'")
	}
}
