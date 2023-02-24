package pip

import "math"

// newSegment will return a new Segment from the provided coordinate values
func newSegment(p1, p2 Point) *segment {
	var s segment
	s.p1 = p1
	s.p2 = p2
	s.m, s.b, s.st = getSlopeOffset(s.p1, s.p2)
	return &s
}

// segment is a line segment
type segment struct {
	p1 Point
	p2 Point

	m  float64
	b  float64
	st slope
}

func (s *segment) GetStart() Point {
	return s.p1
}

func (s *segment) GetEnd() Point {
	return s.p2
}

// GetX will get the x value at the given y coordinate
func (s *segment) GetX(y float64) (x float64) {
	switch s.st {
	case slopeHorizontal:
		return math.NaN()

	case slopeVertical:
		return s.p1.X

	default:
		return solveForX(y, s.m, s.b)
	}
}

// GetY will get the y value at the given x coordinate
func (s *segment) GetY(x float64) (y float64) {
	switch s.st {
	case slopeHorizontal:
		return s.p1.Y

	case slopeVertical:
		return math.NaN()

	default:
		return solveForY(x, s.m, s.b)

	}
}

// Intersects tests if a line extending along the x axis (y = b)
func (s *segment) Intersects(x, y float64) (intersects, contains bool) {
	if s.HasPoint(x, y) {
		contains = true
		return
	}

	switch s.st {
	case slopeHorizontal:
		if y != s.p1.Y {
			return
		}

		return x < s.p1.X, false
	case slopeVertical:
		if x > s.p1.X {
			return
		}

		return s.isWithinY(y), false

	default:
		if !s.containsY(y) {
			return
		}

		var nx float64
		if nx = s.GetX(y); math.IsNaN(nx) || nx < x {
			return
		}

		return s.isWithinSegment(nx, y), false
	}
}

// HasPoint will return whether or not a matching point exists within the segment
func (s *segment) HasPoint(x, y float64) bool {
	pt := MakePoint(x, y)
	if s.p1 == pt || s.p2 == pt {
		return true
	}

	switch s.st {
	case slopeHorizontal:
		return pt.Y == s.p1.Y && s.containsX(pt.X)
	case slopeVertical:
		return pt.X == s.p1.X && s.containsY(pt.Y)
	default:
		if !s.containsY(y) {
			return false
		}

		nx := s.GetX(y)
		return !math.IsNaN(nx) && nx == x
	}
}

func (s *segment) containsX(x float64) bool {
	min, max := s.getXVals()
	return x >= min && x <= max
}

func (s *segment) containsY(y float64) bool {
	min, max := s.getYVals()
	return y >= min && y <= max
}

func (s *segment) isWithinX(x float64) bool {
	if s.p1.X > s.p2.X {
		return x > s.p2.X && x <= s.p1.X
	} else {
		return x >= s.p1.X && x < s.p2.X
	}
}

func (s *segment) isWithinY(y float64) bool {
	if s.p1.Y > s.p2.Y {
		return y > s.p2.Y && y <= s.p1.Y
	} else {
		return y >= s.p1.Y && y < s.p2.Y
	}
}

// isWithinSegment will return whether a provided point (x,y) is potentially within the range of the segment
// Note: This does NOT test whether or not the provided coordinate exists within the segment
func (s *segment) isWithinSegment(x, y float64) bool {
	return s.isWithinX(x) && s.isWithinY(y)
}

func (s *segment) getXVals() (min, max float64) {
	if s.p1.X <= s.p2.X {
		return s.p1.X, s.p2.X
	}

	return s.p2.X, s.p1.X
}

func (s *segment) getYVals() (min, max float64) {
	if s.p1.Y <= s.p2.Y {
		return s.p1.Y, s.p2.Y
	}

	return s.p2.Y, s.p1.Y
}
