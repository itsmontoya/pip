package pip

import "testing"

func Test_getSlope(t *testing.T) {
	type args struct {
		a Point
		b Point
	}

	tests := []struct {
		name  string
		args  args
		wantM float64
	}{
		{
			name: "has slope",
			args: args{
				a: MakePoint(2, 2),
				b: MakePoint(4, 4),
			},
			wantM: 1,
		},
		{
			name: "vertical slope",
			args: args{
				a: MakePoint(2, 2),
				b: MakePoint(2, 4),
			},
			wantM: 0,
		},
		{
			name: "horizontal slope",
			args: args{
				a: MakePoint(2, 2),
				b: MakePoint(4, 2),
			},
			wantM: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := getSlope(tt.args.a, tt.args.b); gotM != tt.wantM {
				t.Errorf("getSlope() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
