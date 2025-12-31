package main

import (
	"bytes"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/go-text/typesetting/font"
	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func Run(window *app.Window) error {
	var ops op.Ops

	// Load font
	parsedFont, err := font.ParseTTF(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	// Create the typeface
	typeface := impl.NewTypefaceWithTypefaceFace("regular", impl.FontStyle{}, parsedFont)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()

			// Clear background to white
			paint.ColorOp{Color: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			canvas := skia.NewCanvas(&ops)
			draw(canvas, typeface)

			e.Frame(&ops)
		}
	}
}

// draw matches the C++ fiddle draw function
func draw(canvas skia.Canvas, typeface interfaces.SkTypeface) {
	paint := impl.NewPaint()
	paint.SetARGB(0xFF, 0, 0, 0) // Default to black

	// C++: SkFont defaultFont = SkFont(fontMgr->matchFamilyStyle(nullptr, {}));
	defaultFont := impl.NewFontWithTypefaceAndSize(typeface, 12)

	textSizes := []float32{12, 18, 24, 36}

	// First loop: Changing font size
	canvas.Save()
	for _, size := range textSizes {
		defaultFont.SetSize(size)
		// canvas->drawString("Aa", 10, 20, defaultFont, paint);
		canvas.DrawString("Aa", 10, 20, defaultFont, paint)
		// canvas->translate(0, size * 2);
		canvas.Translate(0, size*2)
	}
	canvas.Restore() // Restore to clean state for next loop

	// Reset font to initial state (though we set size in loop anyway)
	defaultFont.SetSize(12)

	yPos := float32(20)

	// Second loop: Scaling
	for _, size := range textSizes {
		scale := size / 12.0

		// canvas->resetMatrix(); equivalent here is essentially starting fresh or restoring
		// Since we want to emulate the C++ loop where each iteration starts "fresh" relative to the canvas origin
		// but accumulates yPos manually, we use Save/Restore per iteration.
		canvas.Save()

		// canvas->translate(100, 0);
		canvas.Translate(100, 0)

		// canvas->scale(scale, scale);
		canvas.Scale(scale, scale)

		// canvas->drawString("Aa", 10 / scale, yPos / scale, defaultFont, paint);
		canvas.DrawString("Aa", 10/scale, yPos/scale, defaultFont, paint)

		canvas.Restore()

		yPos += size * 2
	}
}
