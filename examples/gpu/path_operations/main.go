package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/zodimo/gio-skia/skia"
	"github.com/zodimo/go-skia-support/skia/enums"
	"github.com/zodimo/go-skia-support/skia/impl"
	"github.com/zodimo/go-skia-support/skia/models"
)

// This example demonstrates canonical Skia path operations.
// Paths can be combined, transformed, and manipulated in various ways.
// This showcases advanced path construction techniques commonly used in Skia.

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

			// Light gray background
			paint.ColorOp{Color: color.NRGBA{R: 245, G: 245, B: 250, A: 255}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			c := skia.NewCanvas(&ops)
			spacing := float32(200)
			startX, startY := spacing, spacing

			// Example 1: Path addition (combining paths)
			c.Save()
			c.Translate(startX, startY)

			// Create a path by adding multiple shapes
			combinedPath := impl.NewSkPath(enums.PathFillTypeWinding)
			// Add a rectangle
			combinedPath.AddRect(
				models.Rect{Left: -30, Top: -30, Right: 30, Bottom: 30},
				enums.PathDirectionCW, 0)
			// Add a circle
			combinedPath.AddCircle(models.Scalar(0), models.Scalar(0), models.Scalar(25), enums.PathDirectionCW)

			combinedPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 150, B: 255, A: 255})
			c.DrawPath(combinedPath, combinedPaint)

			c.Restore()

			// Example 2: Path transformation
			c.Save()
			c.Translate(startX+spacing, startY)

			// Create base path
			basePath := impl.NewSkPath(enums.PathFillTypeWinding)
			basePath.AddRect(
				models.Rect{Left: -20, Top: -20, Right: 20, Bottom: 20},
				enums.PathDirectionCW, 0)

			// Transform the path using matrix
			transformMatrix := impl.NewMatrixRotate(45)
			transformMatrix.PreTranslate(models.Scalar(0), models.Scalar(0))
			transformedPath := impl.NewSkPath(enums.PathFillTypeWinding)
			transformedPath.AddPathMatrix(basePath, transformMatrix, enums.AddPathModeAppend)

			transformedPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 150, B: 100, A: 255})
			c.DrawPath(transformedPath, transformedPaint)

			c.Restore()

			// Example 3: Path offset
			c.Save()
			c.Translate(startX, startY+spacing)

			// Create original path
			originalPath := impl.NewSkPath(enums.PathFillTypeWinding)
			originalPath.AddCircle(models.Scalar(0), models.Scalar(0), models.Scalar(20), enums.PathDirectionCW)

			// Create offset paths
			offsetPath1 := impl.NewSkPath(enums.PathFillTypeWinding)
			offsetPath1.AddPath(originalPath, models.Scalar(-25), models.Scalar(0), enums.AddPathModeAppend)

			offsetPath2 := impl.NewSkPath(enums.PathFillTypeWinding)
			offsetPath2.AddPath(originalPath, models.Scalar(25), models.Scalar(0), enums.AddPathModeAppend)

			// Draw original
			originalPaint := skia.NewPaintFill(color.NRGBA{R: 200, G: 200, B: 200, A: 255})
			c.DrawPath(originalPath, originalPaint)

			// Draw offsets
			offsetPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 150, A: 255})
			c.DrawPath(offsetPath1, offsetPaint)
			c.DrawPath(offsetPath2, offsetPaint)

			c.Restore()

			// Example 4: Complex path construction with arcs
			c.Save()
			c.Translate(startX+spacing, startY+spacing)

			// Build a complex path using multiple operations
			complexPath := impl.NewSkPath(enums.PathFillTypeWinding)
			// Start with a rectangle
			complexPath.AddRect(
				models.Rect{Left: -30, Top: -30, Right: 30, Bottom: 30},
				enums.PathDirectionCW, 0)
			// Add circles at corners
			complexPath.AddCircle(models.Scalar(-20), models.Scalar(-20), models.Scalar(10), enums.PathDirectionCW)
			complexPath.AddCircle(models.Scalar(20), models.Scalar(-20), models.Scalar(10), enums.PathDirectionCW)
			complexPath.AddCircle(models.Scalar(-20), models.Scalar(20), models.Scalar(10), enums.PathDirectionCW)
			complexPath.AddCircle(models.Scalar(20), models.Scalar(20), models.Scalar(10), enums.PathDirectionCW)

			complexPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 100, B: 200, A: 255})
			c.DrawPath(complexPath, complexPaint)

			c.Restore()

			// Example 5: Path with holes using fill rules
			c.Save()
			c.Translate(startX, startY+spacing*2)

			// Create a path with a hole using winding rule
			holePath := impl.NewSkPath(enums.PathFillTypeWinding)
			// Outer shape (clockwise)
			holePath.AddRect(
				models.Rect{Left: -40, Top: -40, Right: 40, Bottom: 40},
				enums.PathDirectionCW, 0)
			// Inner hole (counter-clockwise)
			holePath.AddCircle(models.Scalar(0), models.Scalar(0), models.Scalar(20), enums.PathDirectionCCW)

			holePaint := skia.NewPaintFill(color.NRGBA{R: 150, G: 200, B: 255, A: 255})
			c.DrawPath(holePath, holePaint)

			c.Restore()

			// Example 6: Path with multiple contours
			c.Save()
			c.Translate(startX+spacing, startY+spacing*2)

			// Create path with multiple separate contours
			multiContourPath := impl.NewSkPath(enums.PathFillTypeWinding)
			// First contour: triangle
			multiContourPath.MoveTo(models.Scalar(-30), models.Scalar(20))
			multiContourPath.LineTo(models.Scalar(0), models.Scalar(-20))
			multiContourPath.LineTo(models.Scalar(30), models.Scalar(20))
			multiContourPath.Close()
			// Second contour: circle
			multiContourPath.AddCircle(models.Scalar(0), models.Scalar(0), models.Scalar(15), enums.PathDirectionCW)

			multiContourPaint := skia.NewPaintFill(color.NRGBA{R: 255, G: 200, B: 100, A: 255})
			c.DrawPath(multiContourPath, multiContourPaint)

			c.Restore()

			// Example 7: Path bounds demonstration
			c.Save()
			c.Translate(startX, startY+spacing*3)

			// Create a path
			boundsPath := impl.NewSkPath(enums.PathFillTypeWinding)
			boundsPath.AddCircle(models.Scalar(0), models.Scalar(0), models.Scalar(30), enums.PathDirectionCW)

			// Get bounds
			bounds := boundsPath.Bounds()

			// Draw the path
			boundsPathPaint := skia.NewPaintFill(color.NRGBA{R: 100, G: 255, B: 200, A: 255})
			c.DrawPath(boundsPath, boundsPathPaint)

			// Draw bounds rectangle
			boundsRect := skia.NewPath()
			skia.PathAddRect(boundsRect,
				float32(bounds.Left), float32(bounds.Top),
				float32(bounds.Right-bounds.Left), float32(bounds.Bottom-bounds.Top))
			boundsPaint := skia.NewPaintStroke(color.NRGBA{R: 255, G: 0, B: 0, A: 255}, 2)
			c.DrawPath(boundsRect, boundsPaint)

			c.Restore()

			// Example 8: Path with Bézier curves
			c.Save()
			c.Translate(startX+spacing, startY+spacing*3)

			// Create a path using Bézier curves
			bezierPath := impl.NewSkPath(enums.PathFillTypeWinding)
			bezierPath.MoveTo(models.Scalar(-30), models.Scalar(0))
			// Cubic Bézier curve
			bezierPath.CubicTo(
				models.Scalar(-20), models.Scalar(-20),
				models.Scalar(20), models.Scalar(-20),
				models.Scalar(30), models.Scalar(0))
			bezierPath.CubicTo(
				models.Scalar(20), models.Scalar(20),
				models.Scalar(-20), models.Scalar(20),
				models.Scalar(-30), models.Scalar(0))
			bezierPath.Close()

			bezierPaint := skia.NewPaintFill(color.NRGBA{R: 200, G: 100, B: 255, A: 255})
			c.DrawPath(bezierPath, bezierPaint)

			c.Restore()

			frameEvent.Frame(&ops)
		}
	}
}
