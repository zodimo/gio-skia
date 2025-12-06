package stroke

import (
	"math"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	andyStroke "github.com/andybalholm/stroke"
)

// expandStroke converts a stroke.Path to a clip.PathSpec representing the stroked outline.
// This is used internally by canvas.DrawPath for stroke operations.
func ExpandStroke(s Path, width float32, join JoinStyle, cap CapStyle,
	miter float32, dash []float32, dash0 float32) clip.PathSpec {
	// Use the andybalholm/stroke library to expand the path
	var opt andyStroke.Options
	opt.Width = width
	opt.MiterLimit = miter
	opt.Cap = cap
	opt.Join = join

	// Convert stroke.Path to [][]stroke.Segment
	var path [][]andyStroke.Segment
	var contour []andyStroke.Segment
	var pen f32.Point

	for _, seg := range s.Segments {
		switch seg.op {
		case segOpMoveTo:
			if len(contour) > 0 {
				path = append(path, contour)
				contour = nil
			}
			pen = seg.args[0]
		case segOpLineTo:
			contour = append(contour, andyStroke.LinearSegment(
				andyStroke.Point(pen),
				andyStroke.Point(seg.args[0]),
			))
			pen = seg.args[0]
		case segOpQuadTo:
			contour = append(contour, andyStroke.QuadraticSegment(
				andyStroke.Point(pen),
				andyStroke.Point(seg.args[0]),
				andyStroke.Point(seg.args[1]),
			))
			pen = seg.args[1]
		case segOpCubeTo:
			contour = append(contour, andyStroke.Segment{
				Start: andyStroke.Point(pen),
				CP1:   andyStroke.Point(seg.args[0]),
				CP2:   andyStroke.Point(seg.args[1]),
				End:   andyStroke.Point(seg.args[2]),
			})
			pen = seg.args[2]
		case segOpArcTo:
			// ArcTo is not used in the basic Skia API, but handle it for completeness
			var (
				start  = andyStroke.Point(pen)
				center = andyStroke.Point(seg.args[0])
				angle  = seg.args[1].X
			)
			if absF32(angle) > math.Pi {
				contour = andyStroke.AppendArc(contour, start, center, angle)
				pen = f32.Point(contour[len(contour)-1].End)
			} else {
				out := andyStroke.ArcSegment(start, center, angle)
				contour = append(contour, out)
				pen = f32.Point(out.End)
			}
		}
	}
	if len(contour) > 0 {
		path = append(path, contour)
	}

	// Apply dashing if provided
	if len(dash) > 0 {
		path = andyStroke.Dash(path, dash, dash0)
	}

	// Stroke the path
	stroked := andyStroke.Stroke(path, opt)

	// Convert back to clip.Path
	var ops op.Ops
	var outline clip.Path
	outline.Begin(&ops)

	for _, contour := range stroked {
		for i, seg := range contour {
			if i == 0 {
				outline.MoveTo(f32.Point(seg.Start))
			}
			outline.CubeTo(
				f32.Point(seg.CP1),
				f32.Point(seg.CP2),
				f32.Point(seg.End),
			)
		}
	}

	return outline.End()
}

func absF32(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}
