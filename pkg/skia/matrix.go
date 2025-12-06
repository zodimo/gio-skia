package skia

// Matrix holds a 3x3 matrix for transforming coordinates.
// This allows mapping Point and vectors with translation, scaling, skewing, rotation, and perspective.
type Matrix interface {
	// Reset sets the matrix to the identity matrix.
	Reset()

	// SetIdentity sets the matrix to the identity matrix.
	SetIdentity()

	// SetScale sets the matrix to scale by (sx, sy).
	SetScale(sx, sy Scalar)

	// SetTranslate sets the matrix to translate by (dx, dy).
	SetTranslate(dx, dy Scalar)

	// SetSkew sets the matrix to skew by (sx, sy).
	SetSkew(sx, sy Scalar)

	// SetRotate sets the matrix to rotate by degrees about a pivot point.
	SetRotate(degrees Scalar, px, py Scalar)

	// SetConcat sets the matrix to the concatenation of a and b.
	SetConcat(a, b Matrix)

	// PreTranslate premultiplies the matrix with a translation.
	PreTranslate(dx, dy Scalar)

	// PreScale premultiplies the matrix with a scale.
	PreScale(sx, sy Scalar)

	// PreSkew premultiplies the matrix with a skew.
	PreSkew(sx, sy Scalar)

	// PreRotate premultiplies the matrix with a rotation.
	PreRotate(degrees Scalar, px, py Scalar)

	// PreConcat premultiplies the matrix with another matrix.
	PreConcat(other Matrix)

	// PostTranslate postmultiplies the matrix with a translation.
	PostTranslate(dx, dy Scalar)

	// PostScale postmultiplies the matrix with a scale.
	PostScale(sx, sy Scalar)

	// PostSkew postmultiplies the matrix with a skew.
	PostSkew(sx, sy Scalar)

	// PostRotate postmultiplies the matrix with a rotation.
	PostRotate(degrees Scalar, px, py Scalar)

	// PostConcat postmultiplies the matrix with another matrix.
	PostConcat(other Matrix)

	// MapPoints applies the matrix transformation to the points.
	MapPoints(dst, src []Point) int

	// MapPoint applies the matrix transformation to a single point.
	MapPoint(p Point) Point

	// MapRect applies the matrix transformation to a rectangle.
	MapRect(rect Rect) Rect

	// MapRectToRect applies the matrix transformation mapping src to dst.
	MapRectToRect(src, dst Rect) bool

	// Invert inverts the matrix if possible.
	Invert() (Matrix, bool)

	// GetScaleX returns the x-axis scale factor.
	GetScaleX() Scalar

	// GetScaleY returns the y-axis scale factor.
	GetScaleY() Scalar

	// GetSkewX returns the x-axis skew factor.
	GetSkewX() Scalar

	// GetSkewY returns the y-axis skew factor.
	GetSkewY() Scalar

	// GetTranslateX returns the x-axis translation.
	GetTranslateX() Scalar

	// GetTranslateY returns the y-axis translation.
	GetTranslateY() Scalar

	// GetPerspX returns the x-axis perspective factor.
	GetPerspX() Scalar

	// GetPerspY returns the y-axis perspective factor.
	GetPerspY() Scalar

	// IsIdentity returns true if the matrix is the identity matrix.
	IsIdentity() bool

	// IsScaleTranslate returns true if the matrix only scales and translates.
	IsScaleTranslate() bool

	// PreservesRightAngles returns true if the matrix preserves right angles.
	PreservesRightAngles() bool

	// HasPerspective returns true if the matrix has perspective.
	HasPerspective() bool

	// RectStaysRect returns true if the matrix maps rectangles to rectangles.
	RectStaysRect() bool

	// GetType returns the type of the matrix.
	GetType() MatrixType
}

// MatrixType describes the type of a matrix.
type MatrixType uint8

const (
	MatrixTypeIdentity MatrixType = iota
	MatrixTypeTranslate
	MatrixTypeScale
	MatrixTypeAffine
	MatrixTypePerspective
)

