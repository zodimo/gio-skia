package stroke

import andyStroke "github.com/andybalholm/stroke"

type CapStyle = andyStroke.CapStyle

const (
	// FlatCap caps stroked paths with a flat cap, joining the right-hand
	// and left-hand sides of a stroked path with a straight line.
	FlatCap = andyStroke.FlatCap
	// RoundCap caps stroked paths with a round cap, joining the right-hand and
	// left-hand sides of a stroked path with a half disc of diameter the
	// stroked path's width.
	RoundCap = andyStroke.RoundCap
	// SquareCap caps stroked paths with a square cap, joining the right-hand
	// and left-hand sides of a stroked path with a half square of length
	// the stroked path's width.
	SquareCap = andyStroke.SquareCap
	// TriangularCap caps stroked paths with a triangular cap, joining the
	// right-hand and left-hand sides of a stroked path with a triangle
	// with height half of the stroked path's width.
	TriangularCap = andyStroke.TriangularCap
)

type JoinStyle = andyStroke.JoinStyle

const (
	// MiterJoin joins path segments with a sharp corner.
	// It falls back to a bevel join if the miter limit is exceeded.
	MiterJoin = andyStroke.MiterJoin
	// RoundJoin joins path segments with a round segment.
	RoundJoin = andyStroke.RoundJoin
	// BevelJoin joins path segments with sharp bevels.
	BevelJoin = andyStroke.BevelJoin
)
