package main

import (
	"bytes"
	"image/color"
	"log"
	"math"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"github.com/go-text/typesetting/font"
	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
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

			canvas := skia.NewCanvas(&ops)

			// Clear background to white
			canvas.Clear(skia.ColorToColor4f(color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}))

			draw(canvas, typeface)

			e.Frame(&ops)
		}
	}
}

// draw matches the C++ fiddle draw function
func draw(canvas skia.Canvas, typeface interfaces.SkTypeface) {
	const iterations = 26
	transforms := make([]models.RSXform, iterations)
	text := make([]byte, iterations)
	var angle float32 = 0
	var scale float32 = 1

	for i := 0; i < iterations; i++ {
		s := float32(math.Sin(float64(angle))) * scale
		c := float32(math.Cos(float64(angle))) * scale
		transforms[i] = models.MakeRSXform(-c, -s, -s*16, c*16)
		angle += .45
		scale += .2
		text[i] = 'A' + byte(i)
	}

	paint := impl.NewPaint()
	font := impl.NewFontWithTypefaceAndSize(typeface, 20)

	// MakeTextBlobFromRSXform usually takes a font, checks for encoding, checks glyphs.
	// We are passing simple ASCII bytes which map 1:1 to glyphs for this font (usually).
	spiral := impl.MakeTextBlobFromRSXform(text, enums.TextEncodingUTF8, transforms, font)

	if spiral != nil {
		canvas.DrawTextBlob(spiral, 110, 138, paint)
	}
}
