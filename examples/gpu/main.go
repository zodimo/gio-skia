package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/skia"
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
	for {
		switch frameEvent := window.Event().(type) {
		case app.DestroyEvent:
			return frameEvent.Err
		case app.FrameEvent:
			// Reset ops for new frame
			ops.Reset()

			// White background — no path, no clip
			paint.ColorOp{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			// Draw test rectangle
			c := skia.NewCanvas(&ops)
			p := skia.NewPath()
			p.AddRect(10, 10, 100, 50)
			c.SetStroke(skia.StrokeOpts{Width: 3, Miter: 4})
			c.SetColor(color.NRGBA{R: 255, A: 255})
			c.DrawPath(p)

			frameEvent.Frame(&ops)
		}
	}
}
