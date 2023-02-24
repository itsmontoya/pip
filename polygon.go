package pip

// NewPolygon will return a new poly target
func New(points []Point) *Polygon {
	var p Polygon
	p.segs = make([]*segment, len(points))

	for i, a := range points {
		j := i + 1
		if j >= len(points) {
			j = 0
		}
		b := points[j]

		p.segs[i] = newSegment(a, b)
	}

	return &p
}

// Polygon is a polygon target
type Polygon struct {
	segs []*segment
}

// IsWithin will return whether ot not a point at a given lat and lon are within a poly target
func (p *Polygon) IsWithin(point Point) (within bool) {
	var n int
	for _, seg := range p.segs {
		intersects, contains := seg.Intersects(point.X, point.Y)
		if contains {
			return true
		}

		if intersects {
			// Point intersects, swap within value
			n++
		}
	}

	return n%2 == 1
}
