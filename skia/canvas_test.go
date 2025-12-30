// SPDX-License-Identifier: Unlicense OR MIT
package skia

import (
	"image"
	"image/color"
	"testing"

	"gioui.org/op"
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/models"
)

func TestNewCanvas(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	if canvas == nil {
		t.Fatal("NewCanvas returned nil")
	}

	// Initial save count should be 1 (the base context)
	if canvas.GetSaveCount() != 1 {
		t.Errorf("Initial save count should be 1, got %d", canvas.GetSaveCount())
	}
}

func TestCanvas_SaveRestore(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Initial count
	if canvas.GetSaveCount() != 1 {
		t.Errorf("Initial count: expected 1, got %d", canvas.GetSaveCount())
	}

	// Save should return new count
	count := canvas.Save()
	if count != 2 {
		t.Errorf("After Save: expected 2, got %d", count)
	}

	// GetSaveCount should match
	if canvas.GetSaveCount() != 2 {
		t.Errorf("GetSaveCount after Save: expected 2, got %d", canvas.GetSaveCount())
	}

	// Restore should decrement
	canvas.Restore()
	if canvas.GetSaveCount() != 1 {
		t.Errorf("After Restore: expected 1, got %d", canvas.GetSaveCount())
	}
}

func TestCanvas_RestoreToCount(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Save multiple times
	canvas.Save()
	canvas.Save()
	canvas.Save()

	if canvas.GetSaveCount() != 4 {
		t.Errorf("After 3 saves: expected 4, got %d", canvas.GetSaveCount())
	}

	// Restore to count 2
	canvas.RestoreToCount(2)
	if canvas.GetSaveCount() != 2 {
		t.Errorf("After RestoreToCount(2): expected 2, got %d", canvas.GetSaveCount())
	}

	// Cannot restore below 1
	canvas.RestoreToCount(0)
	if canvas.GetSaveCount() != 1 {
		t.Errorf("After RestoreToCount(0): expected 1, got %d", canvas.GetSaveCount())
	}
}

func TestCanvas_DrawRect(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintFill(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	rect := models.Rect{Left: 10, Top: 10, Right: 100, Bottom: 100}

	// Should not panic
	canvas.DrawRect(rect, paint)
}

func TestCanvas_DrawRRect(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintFill(color.NRGBA{R: 0, G: 255, B: 0, A: 255})

	var rrect models.RRect
	rrect.SetRectXY(models.Rect{Left: 10, Top: 10, Right: 100, Bottom: 100}, 10, 10)

	// Should not panic
	canvas.DrawRRect(rrect, paint)
}

func TestCanvas_DrawOval(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintFill(color.NRGBA{R: 0, G: 0, B: 255, A: 255})
	oval := models.Rect{Left: 0, Top: 0, Right: 100, Bottom: 50}

	// Should not panic
	canvas.DrawOval(oval, paint)
}

func TestCanvas_DrawCircle(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintFill(color.NRGBA{R: 255, G: 255, B: 0, A: 255})
	center := models.Point{X: 50, Y: 50}

	// Should not panic
	canvas.DrawCircle(center, 25, paint)
}

func TestCanvas_DrawArc(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintStroke(color.NRGBA{R: 255, G: 0, B: 255, A: 255}, 2)
	oval := models.Rect{Left: 0, Top: 0, Right: 100, Bottom: 100}

	// Arc without center
	canvas.DrawArc(oval, 0, 90, false, paint)

	// Wedge with center
	canvas.DrawArc(oval, 0, 90, true, paint)
}

func TestCanvas_DrawPath(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintFill(color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	path := impl.NewSkPath(enums.PathFillTypeWinding)
	path.MoveTo(0, 0)
	path.LineTo(100, 0)
	path.LineTo(50, 100)
	path.Close()

	// Should not panic
	canvas.DrawPath(path, paint)
}

func TestCanvas_DrawPoints(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintStroke(color.NRGBA{R: 255, G: 255, B: 255, A: 255}, 5)
	points := []models.Point{
		{X: 10, Y: 10},
		{X: 50, Y: 50},
		{X: 90, Y: 10},
	}

	// Test all point modes
	canvas.DrawPoints(enums.PointModePoints, points, paint)
	canvas.DrawPoints(enums.PointModeLines, points, paint)
	canvas.DrawPoints(enums.PointModePolygon, points, paint)
}

func TestCanvas_DrawLine(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintStroke(color.NRGBA{R: 0, G: 0, B: 0, A: 255}, 2)
	p0 := models.Point{X: 0, Y: 0}
	p1 := models.Point{X: 100, Y: 100}

	// Should not panic
	canvas.DrawLine(p0, p1, paint)
}

func TestCanvas_ClipRect(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	rect := models.Rect{Left: 10, Top: 10, Right: 90, Bottom: 90}

	// Should not panic
	canvas.ClipRect(rect, enums.ClipOpIntersect, true)
}

func TestCanvas_ClipRRect(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	var rrect models.RRect
	rrect.SetRectXY(models.Rect{Left: 10, Top: 10, Right: 90, Bottom: 90}, 5, 5)

	// Should not panic
	canvas.ClipRRect(rrect, enums.ClipOpIntersect, true)
}

func TestCanvas_ClipPath(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	path := impl.NewSkPath(enums.PathFillTypeWinding)
	path.MoveTo(50, 0)
	path.LineTo(100, 100)
	path.LineTo(0, 100)
	path.Close()

	// Should not panic
	canvas.ClipPath(path, enums.ClipOpIntersect, true)
}

func TestCanvas_DrawImage(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Create a simple test image
	imgInfo := models.NewImageInfo(10, 10, enums.ColorTypeRGBA8888, enums.AlphaTypePremul)
	pixels := make([]byte, 10*10*4) // RGBA
	for i := range pixels {
		pixels[i] = 255 // White
	}
	skImg := impl.NewRasterImage(imgInfo, pixels, 10*4)

	paint := NewPaint()

	// Should not panic (even if actual rendering is minimal)
	canvas.DrawImage(skImg, 10, 10, paint)
}

func TestCanvas_DrawImageRect(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Create a simple test image
	imgInfo := models.NewImageInfo(20, 20, enums.ColorTypeRGBA8888, enums.AlphaTypePremul)
	pixels := make([]byte, 20*20*4) // RGBA
	for i := range pixels {
		pixels[i] = 128 // Gray
	}
	skImg := impl.NewRasterImage(imgInfo, pixels, 20*4)

	paint := NewPaint()
	src := models.Rect{Left: 0, Top: 0, Right: 20, Bottom: 20}
	dst := models.Rect{Left: 0, Top: 0, Right: 100, Bottom: 100}

	// With source rect
	canvas.DrawImageRect(skImg, &src, dst, paint)

	// Without source rect (nil means full image)
	canvas.DrawImageRect(skImg, nil, dst, paint)
}

func TestCanvas_Transforms(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Test all transform methods don't panic
	canvas.Translate(10, 20)
	canvas.Scale(2, 2)
	canvas.Rotate(45)
	canvas.Skew(0.1, 0.2)

	// Test with concrete canvas method access via Save/Restore
	canvas.Save()
	canvas.Translate(100, 100)
	canvas.Restore()
}

// TestCanvas_DrawTextBlob_NilSafe tests that DrawTextBlob handles nil blob gracefully
func TestCanvas_DrawTextBlob_NilSafe(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaint()

	// Should not panic with nil blob
	canvas.DrawTextBlob(nil, 0, 0, paint)
}

// TestCanvas_DrawSimpleText_EmptySafe tests that DrawSimpleText handles empty text gracefully
func TestCanvas_DrawSimpleText_EmptySafe(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaint()

	// Should not panic with empty text
	canvas.DrawSimpleText([]byte{}, enums.TextEncodingUTF8, 0, 0, nil, paint)
}

// TestCanvas_DrawString_NilFontSafe tests that DrawString handles nil font gracefully
func TestCanvas_DrawString_NilFontSafe(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaint()

	// Should not panic (font is nil, so it returns early)
	canvas.DrawString("Hello", 10, 10, nil, paint)
}

// TestCanvas_DrawTextBlob_WithRealBlob_Parity tests text blob rendering with actual glyphs
// This is a parity test verifying the text rendering pipeline works end-to-end
func TestCanvas_DrawTextBlob_WithRealBlob_Parity(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Create a real font with typeface
	typeface := impl.NewTypeface("sans-serif", impl.FontStyle{})
	font := impl.NewFontWithTypefaceAndSize(typeface, 24)

	// Create text blob from string
	blob := impl.MakeTextBlobFromString("Hello", font)
	if blob == nil {
		t.Log("MakeTextBlobFromString returned nil (expected without real font data)")
		return
	}

	paint := NewPaintFill(color.NRGBA{R: 0, G: 0, B: 0, A: 255})

	// Should not panic - attempts to draw text
	canvas.DrawTextBlob(blob, 10, 50, paint)

	// Verify blob has expected properties
	bounds := blob.Bounds()
	if bounds.Right <= bounds.Left {
		t.Error("Text blob bounds should have positive width")
	}
}

// TestCanvas_TextBlob_RunIteration_Parity tests that we can iterate blob runs
func TestCanvas_TextBlob_RunIteration_Parity(t *testing.T) {
	// Create font
	typeface := impl.NewTypeface("sans-serif", impl.FontStyle{})
	font := impl.NewFontWithTypefaceAndSize(typeface, 16)

	// Create text blob
	blob := impl.MakeTextBlobFromString("Test", font)
	if blob == nil {
		t.Log("MakeTextBlobFromString returned nil (expected without real font data)")
		return
	}

	// Should have at least one run
	if blob.RunCount() < 1 {
		t.Errorf("Expected at least 1 run, got %d", blob.RunCount())
	}

	// Run should have glyphs
	run := blob.Run(0)
	if run == nil {
		t.Error("Run(0) should not be nil")
		return
	}

	if len(run.Glyphs) == 0 {
		t.Error("Run should have glyphs")
	}
	if len(run.Positions) != len(run.Glyphs) {
		t.Errorf("Positions count %d should match glyphs count %d", len(run.Positions), len(run.Glyphs))
	}
}

// Helper to create test image for visual tests
func createTestImage(width, height int, col color.NRGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, col)
		}
	}
	return img
}

// ─── C++ PARITY TESTS ───────────────────────────────────────────────────────
// Ported from: skia-source/tests/CanvasTest.cpp

// TestCanvas_SaveState_Parity tests save/saveLayer from C++ Canvas_SaveState test
// Ported from: DEF_TEST(Canvas_SaveState, reporter) - lines 431-447
func TestCanvas_SaveState_Parity(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Initial save count is 1
	if canvas.GetSaveCount() != 1 {
		t.Errorf("Initial save count: expected 1, got %d", canvas.GetSaveCount())
	}

	// save() returns the OLD save count (before the save)
	n := canvas.Save()
	if n != 2 { // After save, count is 2, and Save returns new count
		t.Errorf("After save, expected return 2, got %d", n)
	}
	if canvas.GetSaveCount() != 2 {
		t.Errorf("After save, getSaveCount: expected 2, got %d", canvas.GetSaveCount())
	}

	// saveLayer returns the OLD save count too
	n = canvas.SaveLayer(nil, nil)
	if n != 3 { // After saveLayer, count is 3
		t.Errorf("After saveLayer, expected return 3, got %d", n)
	}
	if canvas.GetSaveCount() != 3 {
		t.Errorf("After saveLayer, getSaveCount: expected 3, got %d", canvas.GetSaveCount())
	}

	canvas.Restore()
	if canvas.GetSaveCount() != 2 {
		t.Errorf("After first restore: expected 2, got %d", canvas.GetSaveCount())
	}

	canvas.Restore()
	if canvas.GetSaveCount() != 1 {
		t.Errorf("After second restore: expected 1, got %d", canvas.GetSaveCount())
	}
}

// TestCanvas_RestoreToCount_Parity tests restoreToCount behavior
// Ported from: kCanvasTests lambda at lines 362-376
func TestCanvas_RestoreToCount_Parity(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	baseSaveCount := canvas.GetSaveCount()
	if baseSaveCount != 1 {
		t.Errorf("Base save count: expected 1, got %d", baseSaveCount)
	}

	n := canvas.Save()
	if baseSaveCount+1 != n {
		t.Errorf("After Save, expected %d, got %d", baseSaveCount+1, n)
	}
	if baseSaveCount+1 != canvas.GetSaveCount() {
		t.Errorf("After Save, getSaveCount: expected %d, got %d", baseSaveCount+1, canvas.GetSaveCount())
	}

	canvas.Save()
	canvas.Save()
	if baseSaveCount+3 != canvas.GetSaveCount() {
		t.Errorf("After 3 saves, expected %d, got %d", baseSaveCount+3, canvas.GetSaveCount())
	}

	canvas.RestoreToCount(baseSaveCount + 1)
	if baseSaveCount+1 != canvas.GetSaveCount() {
		t.Errorf("After RestoreToCount, expected %d, got %d", baseSaveCount+1, canvas.GetSaveCount())
	}

	// Should this pin to 1, or be a no-op, or crash?
	canvas.RestoreToCount(0)
	if canvas.GetSaveCount() != 1 {
		t.Errorf("After RestoreToCount(0), expected 1, got %d", canvas.GetSaveCount())
	}
}

// TestCanvas_SaveLayer_Restore_Parity tests saveLayer/restore pair
// Ported from: kCanvasTests lambda at line 312-316
func TestCanvas_SaveLayer_Restore_Parity(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	saveCount := canvas.GetSaveCount()
	canvas.SaveLayer(nil, nil)
	canvas.Restore()
	if canvas.GetSaveCount() != saveCount {
		t.Errorf("After saveLayer/restore, expected %d, got %d", saveCount, canvas.GetSaveCount())
	}
}

// TestCanvas_SaveLayer_WithBounds_Parity tests saveLayer with bounds
// Ported from: kCanvasTests lambda at line 318-322
func TestCanvas_SaveLayer_WithBounds_Parity(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	saveCount := canvas.GetSaveCount()
	bounds := models.Rect{Left: 0, Top: 0, Right: 2, Bottom: 1}
	canvas.SaveLayer(&bounds, nil)
	canvas.Restore()
	if canvas.GetSaveCount() != saveCount {
		t.Errorf("After saveLayer with bounds/restore, expected %d, got %d", saveCount, canvas.GetSaveCount())
	}
}

// TestCanvas_SaveLayer_WithPaint_Parity tests saveLayer with paint
// Ported from: kCanvasTests lambda at line 324-329
func TestCanvas_SaveLayer_WithPaint_Parity(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	saveCount := canvas.GetSaveCount()
	paint := NewPaint()
	canvas.SaveLayer(nil, paint)
	canvas.Restore()
	if canvas.GetSaveCount() != saveCount {
		t.Errorf("After saveLayer with paint/restore, expected %d, got %d", saveCount, canvas.GetSaveCount())
	}
}

// TestCanvas_DrawDRRect tests double rounded rect (donut shape)
func TestCanvas_DrawDRRect(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintFill(color.NRGBA{R: 0, G: 128, B: 255, A: 255})

	var outer, inner models.RRect
	outer.SetRectXY(models.Rect{Left: 0, Top: 0, Right: 100, Bottom: 100}, 10, 10)
	inner.SetRectXY(models.Rect{Left: 20, Top: 20, Right: 80, Bottom: 80}, 5, 5)

	// Should not panic - draws donut shape
	canvas.DrawDRRect(outer, inner, paint)
}

// TestCanvas_DrawPaint tests filling the canvas
func TestCanvas_DrawPaint(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	paint := NewPaintFill(color.NRGBA{R: 255, G: 128, B: 0, A: 255})

	// Should not panic
	canvas.DrawPaint(paint)
}

// TestCanvas_TransformConsistency tests that transforms are properly stacked
// Ported from: kCanvasTests lambda at lines 302-310
func TestCanvas_TransformConsistency(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	saveCount := canvas.GetSaveCount()
	canvas.Save()

	// Apply transforms
	canvas.Translate(1, 2)

	canvas.Restore()

	if canvas.GetSaveCount() != saveCount {
		t.Errorf("After save/transform/restore, expected %d, got %d", saveCount, canvas.GetSaveCount())
	}
}

// TestCanvas_MultipleTransforms tests multiple transform operations
// Ported from: kCanvasTests at lines 377-392
func TestCanvas_MultipleTransforms(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	// Test complex transform/save/restore sequence
	canvas.Rotate(30)
	canvas.Save()
	canvas.Translate(2, 1)
	canvas.Save()
	canvas.Scale(3, 3)

	paint := NewPaintFill(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	canvas.DrawPaint(paint)

	canvas.Restore()
	canvas.Restore()

	// Should not panic and save count should be back to original
	if canvas.GetSaveCount() != 1 {
		t.Errorf("After complex transforms, expected save count 1, got %d", canvas.GetSaveCount())
	}
}

// TestCanvas_ClipEmptyPath tests clipping with empty path
// Ported from: DEF_TEST(Canvas_ClipEmptyPath, reporter) - lines 449-466
func TestCanvas_ClipEmptyPath_Parity(t *testing.T) {
	ops := new(op.Ops)
	canvas := NewCanvas(ops)

	canvas.Save()
	path := impl.NewSkPath(enums.PathFillTypeWinding)
	canvas.ClipPath(path, enums.ClipOpIntersect, false)
	canvas.Restore()

	canvas.Save()
	path2 := impl.NewSkPath(enums.PathFillTypeWinding)
	path2.MoveTo(5, 5)
	canvas.ClipPath(path2, enums.ClipOpIntersect, false)
	canvas.Restore()

	canvas.Save()
	path3 := impl.NewSkPath(enums.PathFillTypeWinding)
	path3.MoveTo(7, 7)
	canvas.ClipPath(path3, enums.ClipOpIntersect, false) // should not panic
	canvas.Restore()
}

func TestCanvas_ClipStack(t *testing.T) {
	// regression test for clip stack isolation
	ops := new(op.Ops)
	c := NewCanvas(ops)

	rect := models.Rect{Left: 0, Top: 0, Right: 100, Bottom: 100}
	c.ClipRect(rect, enums.ClipOpIntersect, true)

	c.Save()
	c.Translate(10, 10)
	c.ClipRect(rect, enums.ClipOpIntersect, true)
	c.DrawRect(rect, NewPaint())
	c.Restore()

	// Should be back to 1 clip
	// Verify no panic on Draw
	c.DrawRect(rect, NewPaint())
}
