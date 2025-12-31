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
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/interfaces"
	"github.com/zodimo/go-skia-support/skia/models"
	"github.com/zodimo/go-skia-support/skia/shaper"
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
	typeface := impl.NewTypefaceWithTypefaceFace("regular", impl.FontStyle{}, parsedFont)

	// Create fonts
	// SkFont font1(typeface, 64.0f, 1.0f, 0.0f);
	font1 := impl.NewFontWithTypefaceSizeScaleSkew(typeface, 64.0, 1.0, 0.0)
	font1.SetEdging(enums.FontEdgingAntiAlias)

	// SkFont font2(typeface, 64.0f, 1.5f, 0.0f);
	font2 := impl.NewFontWithTypefaceSizeScaleSkew(typeface, 64.0, 1.5, 0.0)
	font2.SetEdging(enums.FontEdgingAntiAlias)

	// Create text blobs
	// sk_sp<SkTextBlob> blob1 = SkTextBlob::MakeFromString("Skia", font1);
	blob1 := makeTextBlob("Skia", font1)
	// sk_sp<SkTextBlob> blob2 = SkTextBlob::MakeFromString("Skia", font2);
	blob2 := makeTextBlob("Skia", font2)

	// Create paints
	// Paint 1: Fill, Blue (0x42, 0x85, 0xF4)
	paint1 := skia.NewPaintFill(color.NRGBA{R: 0x42, G: 0x85, B: 0xF4, A: 0xFF})
	paint1.SetAntiAlias(true)

	// Paint 2: Stroke, Red (0xDB, 0x44, 0x37), Width 3.0
	paint2 := skia.NewPaintStroke(color.NRGBA{R: 0xDB, G: 0x44, B: 0x37, A: 0xFF}, 3.0)
	paint2.SetAntiAlias(true)

	// Paint 3: Fill, Green (0x0F, 0x9D, 0x58)
	paint3 := skia.NewPaintFill(color.NRGBA{R: 0x0F, G: 0x9D, B: 0x58, A: 0xFF})
	paint3.SetAntiAlias(true)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()

			// Clear background to white
			// canvas->clear(SK_ColorWHITE);
			paint.ColorOp{Color: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)

			// canvas->drawTextBlob(blob1.get(), 20.0f, 64.0f, paint1);
			if blob1 != nil {
				c.DrawTextBlob(blob1, 20.0, 64.0, paint1)
			}

			// canvas->drawTextBlob(blob1.get(), 20.0f, 144.0f, paint2);
			if blob1 != nil {
				c.DrawTextBlob(blob1, 20.0, 144.0, paint2)
			}

			// canvas->drawTextBlob(blob2.get(), 20.0f, 224.0f, paint3);
			if blob2 != nil {
				c.DrawTextBlob(blob2, 20.0, 224.0, paint3)
			}

			e.Frame(&ops)
		}
	}
}

// makeTextBlob creates a text blob using the HarfBuzz shaper
func makeTextBlob(text string, font interfaces.SkFont) interfaces.SkTextBlob {
	hbShaper := shaper.NewHarfbuzzShaper()
	handler := shaper.NewTextBlobBuilderRunHandler(text, models.Point{X: 0, Y: 0})
	hbShaper.Shape(text, font, true, 0, handler, nil)
	return handler.MakeBlob()
}
