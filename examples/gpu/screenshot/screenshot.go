package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"gioui.org/gpu/headless"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/skia"
)

func main() {
	w, err := headless.NewWindow(1000, 1000)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Release()
	if err := Run(w); err != nil {
		log.Fatal(err)
	}
}

func Run(window *headless.Window) error {
	var ops op.Ops
	// Reset ops for new frame
	ops.Reset()

	// White background — no path, no clip
	paint.ColorOp{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}.Add(&ops)
	paint.PaintOp{}.Add(&ops)

	// Draw test rectangle
	c := skia.NewCanvas(&ops)
	p := skia.NewPath()
	skia.PathAddRect(p, 10, 10, 100, 50)
	paint := skia.NewPaintStroke(color.NRGBA{R: 255, A: 255}, 3)
	paint.SetStrokeMiter(4)
	c.DrawPath(p, paint)

	window.Frame(&ops)
	//save screenshot
	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: 1000, Y: 1000}})
	var buf bytes.Buffer
	err := window.Screenshot(img)
	if err != nil {
		return err
	}
	err = png.Encode(&buf, img)
	if err != nil {
		return err
	}
	err = os.WriteFile("screenshot.png", buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
