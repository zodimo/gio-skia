package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/models"
	"github.com/zodimo/gio-skia/skia"
)

// This example demonstrates canonical Skia matrix operations.
// Matrices are fundamental to Skia's transformation system, enabling
// complex geometric transformations including translation, rotation, scaling,
// skewing, and perspective transformations.

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
			ops.Reset()

			// Dark background
			paint.ColorOp{Color: color.NRGBA{R: 20, G: 20, B: 30, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			spacing := float32(200)
			startX, startY := spacing, spacing

			// Example 1: Matrix concatenation (pre vs post)
			c.Save()
			c.TranslateFloat32(startX, startY)

			// Base shape
			basePath := skia.NewPath()
			skia.PathAddRect(basePath, -20, -20, 40, 40)

			// Pre-concatenation: translate then rotate
			c.Save()
			c.TranslateFloat32(-30, 0)
			c.RotateFloat32(45)
			prePaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 100, A: 255})
			c.DrawPath(basePath, prePaint)
			c.Restore()

			// Post-concatenation: rotate then translate
			c.Save()
			c.RotateFloat32(45)
			c.TranslateFloat32(30, 0)
			postPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 100, A: 255})
			c.DrawPath(basePath, postPaint)
			c.Restore()

			c.Restore()

			// Example 2: Matrix multiplication
			c.Save()
			c.TranslateFloat32(startX+spacing, startY)

			// Create matrices and multiply them
			translateMatrix := impl.NewMatrixTranslate(models.Scalar(30), models.Scalar(0))
			rotateMatrix := impl.NewMatrixRotate(45)
			scaleMatrix := impl.NewMatrixScale(models.Scalar(1.5), models.Scalar(1.5))

			// Multiply: scale * rotate * translate
			resultMatrix := impl.NewMatrixIdentity()
			resultMatrix.SetConcat(scaleMatrix, rotateMatrix)
			resultMatrix.SetConcat(resultMatrix, translateMatrix)

			c.Concat(resultMatrix)
			multPath := skia.NewPath()
			skia.PathAddRect(multPath, -15, -15, 30, 30)
			multPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 100, B: 255, A: 255})
			c.DrawPath(multPath, multPaint)

			c.Restore()

			// Example 3: Matrix inversion
			c.Save()
			c.TranslateFloat32(startX, startY+spacing)

			// Create a transformation
			transformMatrix := impl.NewMatrixRotate(30)
			transformMatrix.PreTranslate(models.Scalar(40), models.Scalar(0))
			transformMatrix.PreScale(models.Scalar(1.5), models.Scalar(1.5))

			// Apply transformation
			c.Save()
			c.Concat(transformMatrix)
			transformPath := skia.NewPath()
			skia.PathAddRect(transformPath, -15, -15, 30, 30)
			transformPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
			c.DrawPath(transformPath, transformPaint)
			c.Restore()

			// Invert and apply
			invertedMatrix, ok := transformMatrix.Invert()
			if ok {
				c.Save()
				c.Concat(invertedMatrix)
				invertedPath := skia.NewPath()
				skia.PathAddRect(invertedPath, -15, -15, 30, 30)
				invertedPaint := skia.NewPaintFill(color.NRGBA{R: 200, G: 100, B: 255, A: 255})
				c.DrawPath(invertedPath, invertedPaint)
				c.Restore()
			}

			c.Restore()

			// Example 4: Point mapping
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing)

			// Create transformation matrix
			mapMatrix := impl.NewMatrixRotate(60)
			mapMatrix.PreScale(models.Scalar(1.2), models.Scalar(1.2))
			mapMatrix.PreTranslate(models.Scalar(20), models.Scalar(10))

			// Map points
			originalPoints := []models.Point{
				{X: models.Scalar(-20), Y: models.Scalar(-20)},
				{X: models.Scalar(20), Y: models.Scalar(-20)},
				{X: models.Scalar(20), Y: models.Scalar(20)},
				{X: models.Scalar(-20), Y: models.Scalar(20)},
			}

			// Draw original points
			for _, pt := range originalPoints {
				p := skia.NewPath()
				skia.PathAddCircle(p, float32(pt.X), float32(pt.Y), 3)
				originalPointPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 0, B: 0, A: 255})
				c.DrawPath(p, originalPointPaint)
			}

			// Draw mapped points
			for _, pt := range originalPoints {
				mappedX, mappedY := mapMatrix.MapXY(pt.X, pt.Y)
				p := skia.NewPath()
				skia.PathAddCircle(p, float32(mappedX), float32(mappedY), 3)
				mappedPointPaint := skia.NewPaintFill(color.NRGBA{R: 0, G: 255, B: 0, A: 255})
				c.DrawPath(p, mappedPointPaint)
			}

			c.Restore()

			// Example 5: Matrix type classification
			c.Save()
			c.TranslateFloat32(startX, startY+spacing*2)

			// Identity matrix
			identityMatrix := impl.NewMatrixIdentity()
			c.Save()
			c.Concat(identityMatrix)
			identityPath := skia.NewPath()
			skia.PathAddRect(identityPath, -15, -15, 30, 30)
			identityPaint := skia.NewPaintFill(color.NRGBA{R: 200, G: 200, B: 200, A: 255})
			c.DrawPath(identityPath, identityPaint)
			c.Restore()

			// Translate-only matrix
			translateOnlyMatrix := impl.NewMatrixTranslate(models.Scalar(40), models.Scalar(0))
			c.Save()
			c.Concat(translateOnlyMatrix)
			translateOnlyPath := skia.NewPath()
			skia.PathAddRect(translateOnlyPath, -15, -15, 30, 30)
			translateOnlyPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 150, B: 150, A: 255})
			c.DrawPath(translateOnlyPath, translateOnlyPaint)
			c.Restore()

			c.Restore()

			// Example 6: Complex matrix chain
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing*2)

			// Build a complex transformation chain
			chainMatrix := impl.NewMatrixIdentity()
			chainMatrix.SetConcat(impl.NewMatrixTranslate(models.Scalar(30), models.Scalar(0)), chainMatrix)
			chainMatrix.SetConcat(impl.NewMatrixRotate(45), chainMatrix)
			chainMatrix.SetConcat(impl.NewMatrixScale(models.Scalar(1.3), models.Scalar(1.3)), chainMatrix)
			chainMatrix.SetConcat(impl.NewMatrixRotate(-15), chainMatrix)

			c.Concat(chainMatrix)
			chainPath := skia.NewPath()
			skia.PathAddRect(chainPath, -20, -20, 40, 40)
			chainPaint := skia.NewPaintFill(color.NRGBA{R: 150, G: 255, B: 200, A: 255})
			c.DrawPath(chainPath, chainPaint)

			c.Restore()

			// Example 7: Matrix with pivot point rotation
			c.Save()
			c.TranslateFloat32(startX, startY+spacing*3)

			// Rotate around a pivot point
			pivotX, pivotY := models.Scalar(30), models.Scalar(30)
			pivotMatrix := impl.NewMatrixRotateWithPivot(45, pivotX, pivotY)

			c.Concat(pivotMatrix)
			pivotPath := skia.NewPath()
			skia.PathAddRect(pivotPath, 0, 0, 30, 30)
			pivotPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 255, A: 255})
			c.DrawPath(pivotPath, pivotPaint)

			// Draw pivot point
			pivotPointPath := skia.NewPath()
			skia.PathAddCircle(pivotPointPath, float32(pivotX), float32(pivotY), 3)
			pivotPointPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 255, B: 0, A: 255})
			c.DrawPath(pivotPointPath, pivotPointPaint)

			c.Restore()

			// Example 8: Matrix decomposition demonstration
			c.Save()
			c.TranslateFloat32(startX+spacing, startY+spacing*3)

			// Create a complex matrix
			complexMatrix := impl.NewMatrixRotate(30)
			complexMatrix.PreScale(models.Scalar(1.5), models.Scalar(0.8))
			complexMatrix.PreTranslate(models.Scalar(20), models.Scalar(10))

			// Apply matrix
			c.Concat(complexMatrix)
			decompPath := skia.NewPath()
			skia.PathAddRect(decompPath, -15, -15, 30, 30)
			decompPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 200, B: 255, A: 255})
			c.DrawPath(decompPath, decompPaint)

			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}

