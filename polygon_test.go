package pip

import (
	"testing"
)

func TestPolygon_IsWithin(t *testing.T) {
	type fields struct {
		points []Point
	}

	type testcase struct {
		Point   Point
		Matches bool
	}

	type testgroup struct {
		name   string
		fields fields
		tests  []testcase
	}

	groups := []testgroup{
		{
			name: "square",
			fields: fields{
				points: []Point{
					MakePoint(1, 1),
					MakePoint(1, 5),
					MakePoint(5, 5),
					MakePoint(5, 1),
				},
			},
			tests: []testcase{
				// Latitude checks
				{Point: MakePoint(0, 0), Matches: false},
				{Point: MakePoint(0, 1), Matches: false},
				{Point: MakePoint(0, 2), Matches: false},
				{Point: MakePoint(0, 3), Matches: false},
				{Point: MakePoint(0, 4), Matches: false},
				{Point: MakePoint(0, 5), Matches: false},
				{Point: MakePoint(0, 0), Matches: false},

				{Point: MakePoint(1, 0), Matches: false},
				{Point: MakePoint(1, 1), Matches: true},
				{Point: MakePoint(1, 2), Matches: true},
				{Point: MakePoint(1, 3), Matches: true},
				{Point: MakePoint(1, 4), Matches: true},
				{Point: MakePoint(1, 5), Matches: true},
				{Point: MakePoint(1, 0), Matches: false},

				{Point: MakePoint(2, 0), Matches: false},
				{Point: MakePoint(2, 1), Matches: true},
				{Point: MakePoint(2, 2), Matches: true},
				{Point: MakePoint(2, 3), Matches: true},
				{Point: MakePoint(2, 4), Matches: true},
				{Point: MakePoint(2, 5), Matches: true},
				{Point: MakePoint(2, 0), Matches: false},

				{Point: MakePoint(3, 0), Matches: false},
				{Point: MakePoint(3, 1), Matches: true},
				{Point: MakePoint(3, 2), Matches: true},
				{Point: MakePoint(3, 3), Matches: true},
				{Point: MakePoint(3, 4), Matches: true},
				{Point: MakePoint(3, 5), Matches: true},
				{Point: MakePoint(3, 0), Matches: false},

				{Point: MakePoint(4, 0), Matches: false},
				{Point: MakePoint(4, 1), Matches: true},
				{Point: MakePoint(4, 2), Matches: true},
				{Point: MakePoint(4, 3), Matches: true},
				{Point: MakePoint(4, 4), Matches: true},
				{Point: MakePoint(4, 5), Matches: true},
				{Point: MakePoint(4, 0), Matches: false},

				{Point: MakePoint(5, 0), Matches: false},
				{Point: MakePoint(5, 1), Matches: true},
				{Point: MakePoint(5, 2), Matches: true},
				{Point: MakePoint(5, 3), Matches: true},
				{Point: MakePoint(5, 4), Matches: true},
				{Point: MakePoint(5, 5), Matches: true},
				{Point: MakePoint(5, 0), Matches: false},

				{Point: MakePoint(6, 0), Matches: false},
				{Point: MakePoint(6, 1), Matches: false},
				{Point: MakePoint(6, 2), Matches: false},
				{Point: MakePoint(6, 3), Matches: false},
				{Point: MakePoint(6, 4), Matches: false},
				{Point: MakePoint(6, 5), Matches: false},
				{Point: MakePoint(6, 0), Matches: false},

				// Longitude checks
				{Point: MakePoint(0, 0), Matches: false},
				{Point: MakePoint(1, 0), Matches: false},
				{Point: MakePoint(2, 0), Matches: false},
				{Point: MakePoint(3, 0), Matches: false},
				{Point: MakePoint(4, 0), Matches: false},
				{Point: MakePoint(5, 0), Matches: false},
				{Point: MakePoint(6, 0), Matches: false},

				{Point: MakePoint(0, 1), Matches: false},
				{Point: MakePoint(1, 1), Matches: true},
				{Point: MakePoint(2, 1), Matches: true},
				{Point: MakePoint(3, 1), Matches: true},
				{Point: MakePoint(4, 1), Matches: true},
				{Point: MakePoint(5, 1), Matches: true},
				{Point: MakePoint(6, 1), Matches: false},

				{Point: MakePoint(0, 2), Matches: false},
				{Point: MakePoint(1, 2), Matches: true},
				{Point: MakePoint(2, 2), Matches: true},
				{Point: MakePoint(3, 2), Matches: true},
				{Point: MakePoint(4, 2), Matches: true},
				{Point: MakePoint(5, 2), Matches: true},
				{Point: MakePoint(6, 2), Matches: false},

				{Point: MakePoint(0, 3), Matches: false},
				{Point: MakePoint(1, 3), Matches: true},
				{Point: MakePoint(2, 3), Matches: true},
				{Point: MakePoint(3, 3), Matches: true},
				{Point: MakePoint(4, 3), Matches: true},
				{Point: MakePoint(5, 3), Matches: true},
				{Point: MakePoint(6, 3), Matches: false},

				{Point: MakePoint(0, 4), Matches: false},
				{Point: MakePoint(1, 4), Matches: true},
				{Point: MakePoint(2, 4), Matches: true},
				{Point: MakePoint(3, 4), Matches: true},
				{Point: MakePoint(4, 4), Matches: true},
				{Point: MakePoint(5, 4), Matches: true},
				{Point: MakePoint(6, 4), Matches: false},

				{Point: MakePoint(0, 5), Matches: false},
				{Point: MakePoint(1, 5), Matches: true},
				{Point: MakePoint(2, 5), Matches: true},
				{Point: MakePoint(3, 5), Matches: true},
				{Point: MakePoint(4, 5), Matches: true},
				{Point: MakePoint(5, 5), Matches: true},
				{Point: MakePoint(6, 5), Matches: false},

				{Point: MakePoint(0, 6), Matches: false},
				{Point: MakePoint(1, 6), Matches: false},
				{Point: MakePoint(2, 6), Matches: false},
				{Point: MakePoint(3, 6), Matches: false},
				{Point: MakePoint(4, 6), Matches: false},
				{Point: MakePoint(5, 6), Matches: false},
				{Point: MakePoint(6, 6), Matches: false},
			},
		},
		{
			name: "octogon",
			fields: fields{
				points: []Point{
					MakePoint(4, 0),
					MakePoint(0, 4),
					MakePoint(0, 8),
					MakePoint(4, 12),
					MakePoint(8, 12),
					MakePoint(12, 8),
					MakePoint(12, 4),
					MakePoint(8, 0),
				},
			},
			tests: []testcase{
				{Point: MakePoint(0, 0), Matches: false},
				{Point: MakePoint(0, 2), Matches: false},
				{Point: MakePoint(0, 4), Matches: true},
				{Point: MakePoint(0, 6), Matches: true},
				{Point: MakePoint(0, 8), Matches: true},
				{Point: MakePoint(0, 10), Matches: false},
				{Point: MakePoint(0, 12), Matches: false},

				{Point: MakePoint(2, 0), Matches: false},
				{Point: MakePoint(2, 2), Matches: true},
				{Point: MakePoint(2, 4), Matches: true},
				{Point: MakePoint(2, 6), Matches: true},
				{Point: MakePoint(2, 8), Matches: true},
				{Point: MakePoint(2, 10), Matches: true},
				{Point: MakePoint(2, 12), Matches: false},

				{Point: MakePoint(4, 0), Matches: true},
				{Point: MakePoint(4, 2), Matches: true},
				{Point: MakePoint(4, 4), Matches: true},
				{Point: MakePoint(4, 6), Matches: true},
				{Point: MakePoint(4, 8), Matches: true},
				{Point: MakePoint(4, 10), Matches: true},
				{Point: MakePoint(4, 12), Matches: true},

				{Point: MakePoint(6, 0), Matches: true},
				{Point: MakePoint(6, 2), Matches: true},
				{Point: MakePoint(6, 4), Matches: true},
				{Point: MakePoint(6, 6), Matches: true},
				{Point: MakePoint(6, 8), Matches: true},
				{Point: MakePoint(6, 10), Matches: true},
				{Point: MakePoint(6, 12), Matches: true},

				{Point: MakePoint(8, 0), Matches: true},
				{Point: MakePoint(8, 2), Matches: true},
				{Point: MakePoint(8, 4), Matches: true},
				{Point: MakePoint(8, 6), Matches: true},
				{Point: MakePoint(8, 8), Matches: true},
				{Point: MakePoint(8, 10), Matches: true},
				{Point: MakePoint(8, 12), Matches: true},

				{Point: MakePoint(12, 0), Matches: false},
				{Point: MakePoint(12, 2), Matches: false},
				{Point: MakePoint(12, 4), Matches: true},
				{Point: MakePoint(12, 6), Matches: true},
				{Point: MakePoint(12, 8), Matches: true},
				{Point: MakePoint(12, 10), Matches: false},
				{Point: MakePoint(12, 12), Matches: false},
			},
		},
	}

	for _, group := range groups {
		t.Run(group.name, func(t *testing.T) {
			for _, tt := range group.tests {
				p := New(group.fields.points)
				if gotWithin := p.IsWithin(tt.Point); gotWithin != tt.Matches {
					t.Errorf("Polygon.IsWithin() = %v, want %v", gotWithin, tt.Matches)
				}
			}
		})
	}
}
