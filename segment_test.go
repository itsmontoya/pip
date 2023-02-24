package pip

import (
	"math"
	"reflect"
	"testing"
)

func TestSegment_GetStart(t *testing.T) {
	type fields struct {
		p1 Point
		p2 Point
	}

	tests := []struct {
		name   string
		fields fields
		want   Point
	}{
		{
			name: "basic",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(5, 6),
			},
			want: MakePoint(3, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSegment(tt.fields.p1, tt.fields.p2)
			if got := s.GetStart(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Segment.GetStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegment_GetEnd(t *testing.T) {
	type fields struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name   string
		fields fields
		want   Point
	}{
		{
			name: "basic",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(5, 6),
			},
			want: MakePoint(5, 6),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSegment(tt.fields.p1, tt.fields.p2)
			if got := s.GetEnd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Segment.GetEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegment_GetX(t *testing.T) {
	type fields struct {
		p1 Point
		p2 Point
	}

	type args struct {
		y float64
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantX  float64
	}{
		{
			name: "default - start",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:  args{y: 2},
			wantX: 2,
		},
		{
			name: "default - middle",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:  args{y: 3},
			wantX: 3,
		},
		{
			name: "default - end",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:  args{y: 4},
			wantX: 4,
		},
		{
			name: "horizontal",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args:  args{y: 4},
			wantX: math.NaN(),
		},
		{
			name: "vertical",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args:  args{y: 4},
			wantX: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSegment(tt.fields.p1, tt.fields.p2)
			gotX := s.GetX(tt.args.y)
			switch {
			case gotX == tt.wantX:
			case math.IsNaN(gotX) && math.IsNaN(tt.wantX):
			default:
				t.Errorf("Segment.GetX() = %v, want %v", gotX, tt.wantX)
			}
		})
	}
}

func TestSegment_GetY(t *testing.T) {
	type fields struct {
		p1 Point
		p2 Point
	}

	type args struct {
		x float64
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantY  float64
	}{
		{
			name: "default - start",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:  args{x: 2},
			wantY: 2,
		},
		{
			name: "default - middle",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:  args{x: 3},
			wantY: 3,
		},
		{
			name: "default - end",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:  args{x: 4},
			wantY: 4,
		},
		{
			name: "horizontal",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args:  args{x: 4},
			wantY: 4,
		},
		{
			name: "vertical",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args:  args{x: 4},
			wantY: math.NaN(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSegment(tt.fields.p1, tt.fields.p2)
			gotY := s.GetY(tt.args.x)
			switch {
			case gotY == tt.wantY:
			case math.IsNaN(gotY) && math.IsNaN(tt.wantY):
			default:
				t.Errorf("Segment.GetY() = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestSegment_Intersects(t *testing.T) {
	type fields struct {
		p1 Point
		p2 Point
	}

	type args struct {
		x float64
		y float64
	}

	tests := []struct {
		name           string
		fields         fields
		args           args
		wantIntersects bool
		wantContains   bool
	}{
		{
			name: "default - no match",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:           args{x: 6, y: 6},
			wantIntersects: false,
			wantContains:   false,
		},
		{
			name: "default - intersects",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:           args{x: 2, y: 3},
			wantIntersects: true,
			wantContains:   false,
		},

		{
			name: "default - contains (start)",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:           args{x: 2, y: 2},
			wantIntersects: false,
			wantContains:   true,
		},
		{
			name: "default - contains",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args:           args{x: 3, y: 3},
			wantIntersects: false,
			wantContains:   true,
		},

		{
			name: "horizontal - no match",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args:           args{x: 4, y: 6},
			wantIntersects: false,
			wantContains:   false,
		},
		{
			name: "horizontal - intersects",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args:           args{x: 2, y: 4},
			wantIntersects: true,
			wantContains:   false,
		},
		{
			name: "horizontal - contains (start)",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args:           args{x: 3, y: 4},
			wantIntersects: false,
			wantContains:   true,
		},
		{
			name: "horizontal - contains",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args:           args{x: 4, y: 4},
			wantIntersects: false,
			wantContains:   true,
		},

		{
			name: "vertical - no match",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args:           args{x: 4, y: 4},
			wantIntersects: false,
			wantContains:   false,
		},
		{
			name: "vertical - intersects",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args:           args{x: 1, y: 4},
			wantIntersects: true,
			wantContains:   false,
		},

		{
			name: "vertical - contains (start)",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args:           args{x: 2, y: 4},
			wantIntersects: false,
			wantContains:   true,
		},
		{
			name: "vertical - contains",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args:           args{x: 2, y: 6},
			wantIntersects: false,
			wantContains:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSegment(tt.fields.p1, tt.fields.p2)
			gotIntersects, gotContains := s.Intersects(tt.args.x, tt.args.y)
			if gotIntersects != tt.wantIntersects {
				t.Errorf("Segment.Intersects() gotIntersects = %v, want %v", gotIntersects, tt.wantIntersects)
			}
			if gotContains != tt.wantContains {
				t.Errorf("Segment.Intersects() gotContains = %v, want %v", gotContains, tt.wantContains)
			}
		})
	}
}

func TestSegment_HasPoint(t *testing.T) {
	type fields struct {
		p1 Point
		p2 Point
	}

	type args struct {
		x float64
		y float64
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "default - no match",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args: args{x: 6, y: 6},
			want: false,
		},
		{
			name: "default - intersects",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args: args{x: 2, y: 3},
			want: false,
		},

		{
			name: "default - contains (start)",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args: args{x: 2, y: 2},
			want: true,
		},
		{
			name: "default - contains",
			fields: fields{
				p1: MakePoint(2, 2),
				p2: MakePoint(4, 4),
			},
			args: args{x: 3, y: 3},
			want: true,
		},

		{
			name: "horizontal - no match",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args: args{x: 4, y: 6},
			want: false,
		},
		{
			name: "horizontal - intersects",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args: args{x: 2, y: 4},
			want: false,
		},
		{
			name: "horizontal - contains (start)",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args: args{x: 3, y: 4},
			want: true,
		},
		{
			name: "horizontal - contains",
			fields: fields{
				p1: MakePoint(3, 4),
				p2: MakePoint(6, 4),
			},
			args: args{x: 4, y: 4},
			want: true,
		},

		{
			name: "vertical - no match",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args: args{x: 4, y: 4},
			want: false,
		},
		{
			name: "vertical - intersects",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args: args{x: 1, y: 4},
			want: false,
		},

		{
			name: "vertical - contains (start)",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args: args{x: 2, y: 4},
			want: true,
		},
		{
			name: "vertical - contains",
			fields: fields{
				p1: MakePoint(2, 4),
				p2: MakePoint(2, 8),
			},
			args: args{x: 2, y: 6},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSegment(tt.fields.p1, tt.fields.p2)
			if got := s.HasPoint(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Segment.HasPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
