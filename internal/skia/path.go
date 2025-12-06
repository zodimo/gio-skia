package skia

// Path contains geometry. Path may be empty, or contain one or more verbs that
// outline a figure. Path always starts with a move verb to a Cartesian coordinate,
// and may be followed by additional verbs that add lines or curves.
type Path interface {
	// FillType returns the fill type used to determine which parts are inside.
	FillType() PathFillType

	// SetFillType sets the fill type used to determine which parts are inside.
	SetFillType(fillType PathFillType)

	// IsInverseFillType returns true if the fill type is inverse.
	IsInverseFillType() bool

	// ToggleInverseFillType toggles between inverse and non-inverse fill types.
	ToggleInverseFillType()

	// Convexity returns the convexity type of the path.
	Convexity() PathConvexity

	// IsConvex returns true if the path is convex.
	IsConvex() bool

	// Reset clears the path, removing all verbs, points, and conic weights.
	Reset()

	// IsEmpty returns true if the path has no verbs.
	IsEmpty() bool

	// IsFinite returns true if all points in the path are finite.
	IsFinite() bool

	// IsLine returns true if the path contains only one line.
	IsLine() bool

	// CountPoints returns the number of points in the path.
	CountPoints() int

	// Point returns the point at the specified index.
	Point(index int) Point

	// GetPoints copies all points from the path into the provided slice.
	GetPoints(points []Point) int

	// CountVerbs returns the number of verbs in the path.
	CountVerbs() int

	// GetVerbs copies all verbs from the path into the provided slice.
	GetVerbs(verbs []PathVerb) int

	// Bounds returns the bounding box of the path.
	Bounds() Rect

	// UpdateBoundsCache updates the cached bounds of the path.
	UpdateBoundsCache()

	// ComputeTightBounds returns a tight bounding box of the path.
	ComputeTightBounds() Rect

	// MoveTo starts a new contour at the specified point.
	MoveTo(x, y Scalar)

	// MoveToPoint starts a new contour at the specified point.
	MoveToPoint(p Point)

	// LineTo adds a line from the last point to the specified point.
	LineTo(x, y Scalar)

	// LineToPoint adds a line from the last point to the specified point.
	LineToPoint(p Point)

	// QuadTo adds a quadratic bezier from the last point to the specified point.
	QuadTo(cx, cy, x, y Scalar)

	// QuadToPoint adds a quadratic bezier from the last point to the specified point.
	QuadToPoint(c, p Point)

	// ConicTo adds a conic bezier from the last point to the specified point.
	ConicTo(cx, cy, x, y Scalar, w Scalar)

	// ConicToPoint adds a conic bezier from the last point to the specified point.
	ConicToPoint(c, p Point, w Scalar)

	// CubicTo adds a cubic bezier from the last point to the specified point.
	CubicTo(cx1, cy1, cx2, cy2, x, y Scalar)

	// CubicToPoint adds a cubic bezier from the last point to the specified point.
	CubicToPoint(c1, c2, p Point)

	// Close closes the current contour.
	Close()

	// AddRect adds a rectangle to the path.
	AddRect(rect Rect, dir PathDirection, startIndex uint)

	// AddOval adds an oval to the path.
	AddOval(rect Rect, dir PathDirection)

	// AddCircle adds a circle to the path.
	AddCircle(cx, cy, radius Scalar, dir PathDirection)

	// AddRRect adds a rounded rectangle to the path.
	AddRRect(rrect RRect, dir PathDirection)

	// AddPath adds another path to this path.
	AddPath(path Path, dx, dy Scalar, addMode AddPathMode)

	// Transform applies a matrix transformation to the path.
	Transform(matrix Matrix)

	// Offset translates the path by the specified offset.
	Offset(dx, dy Scalar)
}

// PathConvexity describes the convexity of a path.
type PathConvexity uint8

const (
	PathConvexityUnknown PathConvexity = iota
	PathConvexityConvex
	PathConvexityConcave
)

// AddPathMode specifies how paths are combined.
type AddPathMode uint8

const (
	AddPathModeAppend AddPathMode = iota
	AddPathModeExtend
)

// RRect represents a rounded rectangle.
type RRect interface {
	// Type returns the type of the rounded rectangle.
	Type() RRectType

	// Rect returns the bounding rectangle.
	Rect() Rect

	// Width returns the width of the rounded rectangle.
	Width() Scalar

	// Height returns the height of the rounded rectangle.
	Height() Scalar

	// GetRadii returns the radii for the specified corner.
	GetRadii(corner RRectCorner) Point

	// SetRect sets the rounded rectangle to a rectangle.
	SetRect(rect Rect)

	// SetOval sets the rounded rectangle to an oval.
	SetOval(rect Rect)

	// SetRectXY sets the rounded rectangle with uniform radii.
	SetRectXY(rect Rect, rx, ry Scalar)

	// SetNinePatch sets the rounded rectangle with different radii for each corner.
	SetNinePatch(rect Rect, rx1, ry1, rx2, ry2, rx3, ry3, rx4, ry4 Scalar)

	// SetRectRadii sets the rounded rectangle with radii for each corner.
	SetRectRadii(rect Rect, radii [4]Point)

	// IsEmpty returns true if the rounded rectangle is empty.
	IsEmpty() bool

	// IsRect returns true if the rounded rectangle is a rectangle.
	IsRect() bool

	// IsOval returns true if the rounded rectangle is an oval.
	IsOval() bool

	// IsSimple returns true if the rounded rectangle is simple.
	IsSimple() bool

	// IsNinePatch returns true if the rounded rectangle is a nine-patch.
	IsNinePatch() bool

	// IsComplex returns true if the rounded rectangle is complex.
	IsComplex() bool
}

// RRectType describes the type of a rounded rectangle.
type RRectType uint8

const (
	RRectTypeEmpty RRectType = iota
	RRectTypeRect
	RRectTypeOval
	RRectTypeSimple
	RRectTypeNinePatch
	RRectTypeComplex
)

// RRectCorner specifies which corner of a rounded rectangle.
type RRectCorner uint8

const (
	RRectCornerUpperLeft RRectCorner = iota
	RRectCornerUpperRight
	RRectCornerLowerRight
	RRectCornerLowerLeft
)

