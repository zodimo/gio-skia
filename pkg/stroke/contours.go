package stroke

import (
	"math"

	"gioui.org/f32"
	"github.com/andybalholm/stroke"
	andyStroke "github.com/andybalholm/stroke"
)

// strokedContours computes stroked path segments
func StrokedContours(s Path, opts StrokeOpts) [][]andyStroke.Segment {
	var path [][]andyStroke.Segment
	var contour []andyStroke.Segment
	var pen f32.Point

	for _, seg := range s.Segments {
		switch seg.Op {
		case segOpMoveTo:
			if len(contour) > 0 {
				path = append(path, contour)
				contour = nil
			}
			pen = seg.Args[0]
		case segOpLineTo:
			contour = append(contour, andyStroke.LinearSegment(
				andyStroke.Point(pen), andyStroke.Point(seg.Args[0]),
			))
			pen = seg.Args[0]
		case segOpQuadTo:
			contour = append(contour, andyStroke.QuadraticSegment(
				andyStroke.Point(pen), andyStroke.Point(seg.Args[0]), andyStroke.Point(seg.Args[1]),
			))
			pen = seg.Args[1]
		case segOpCubeTo:
			contour = append(contour, andyStroke.Segment{
				Start: andyStroke.Point(pen),
				CP1:   andyStroke.Point(seg.Args[0]),
				CP2:   andyStroke.Point(seg.Args[1]),
				End:   andyStroke.Point(seg.Args[2]),
			})
			pen = seg.Args[2]
		case segOpArcTo:
			var (
				start  = andyStroke.Point(pen)
				center = andyStroke.Point(seg.Args[0])
				angle  = seg.Args[1].X
			)
			if absF32(angle) > math.Pi {
				contour = andyStroke.AppendArc(contour, start, center, angle)
				pen = f32.Point(contour[len(contour)-1].End)
			} else {
				out := stroke.ArcSegment(start, center, angle)
				contour = append(contour, out)
				pen = f32.Point(out.End)
			}
		}
	}
	if len(contour) > 0 {
		path = append(path, contour)
	}

	if len(opts.Dash) > 0 {
		path = stroke.Dash(path, opts.Dash, opts.Dash0)
	}

	var opt stroke.Options
	opt.Width = opts.Width
	opt.MiterLimit = opts.Miter
	opt.Cap = opts.Cap
	opt.Join = opts.Join

	return stroke.Stroke(path, opt)
}
