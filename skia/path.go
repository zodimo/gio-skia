package skia

type path struct {
	verbs []pathOp
}

type pathOp struct {
	verb uint8
	pts  [6]float32
}

// ── Path construction ─────────────────────────────────────────────────

func NewPath() Path {
	return &path{}
}

func (p *path) MoveTo(x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 0, pts: [6]float32{x, y}})
}

func (p *path) LineTo(x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 1, pts: [6]float32{x, y}})
}

func (p *path) QuadTo(cx, cy, x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 2, pts: [6]float32{cx, cy, x, y}})
}

func (p *path) CubeTo(cx1, cy1, cx2, cy2, x, y float32) {
	p.verbs = append(p.verbs, pathOp{verb: 3, pts: [6]float32{cx1, cy1, cx2, cy2, x, y}})
}

func (p *path) Close() {
	p.verbs = append(p.verbs, pathOp{verb: 4})
}

func (p *path) AddRect(x, y, w, h float32) {
	p.MoveTo(x, y)
	p.LineTo(x+w, y)
	p.LineTo(x+w, y+h)
	p.LineTo(x, y+h)
	p.Close()
}

func (p *path) AddCircle(cx, cy, r float32) {
	const k = 0.5522848
	p.MoveTo(cx+r, cy)
	p.CubeTo(cx+r, cy+k*r, cx+k*r, cy+r, cx, cy+r)
	p.CubeTo(cx-k*r, cy+r, cx-r, cy+k*r, cx-r, cy)
	p.CubeTo(cx-r, cy-k*r, cx-k*r, cy-r, cx, cy-r)
	p.CubeTo(cx+k*r, cy-r, cx+r, cy-k*r, cx+r, cy)
}

func (p *path) unwrap() []pathOp {
	return p.verbs
}
