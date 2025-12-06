package stroke

import "github.com/andybalholm/stroke"

type CapStyle = stroke.CapStyle

const (
	// FlatCap caps stroked paths with a flat cap, joining the right-hand
	// and left-hand sides of a stroked path with a straight line.
	FlatCap = stroke.FlatCap
	// RoundCap caps stroked paths with a round cap, joining the right-hand and
	// left-hand sides of a stroked path with a half disc of diameter the
	// stroked path's width.
	RoundCap = stroke.RoundCap
	// SquareCap caps stroked paths with a square cap, joining the right-hand
	// and left-hand sides of a stroked path with a half square of length
	// the stroked path's width.
	SquareCap = stroke.SquareCap
	// TriangularCap caps stroked paths with a triangular cap, joining the
	// right-hand and left-hand sides of a stroked path with a triangle
	// with height half of the stroked path's width.
	TriangularCap = stroke.TriangularCap
)

type JoinStyle = stroke.JoinStyle

const (
	// MiterJoin joins path segments with a sharp corner.
	// It falls back to a bevel join if the miter limit is exceeded.
	MiterJoin = stroke.MiterJoin
	// RoundJoin joins path segments with a round segment.
	RoundJoin = stroke.RoundJoin
	// BevelJoin joins path segments with sharp bevels.
	BevelJoin = stroke.BevelJoin
)
