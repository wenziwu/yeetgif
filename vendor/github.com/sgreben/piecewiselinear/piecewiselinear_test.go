package piecewiselinear

import (
	"fmt"
	"testing"
)

func Example() {
	f := Function{Y: []float64{0, 1, 0}} // range: "hat" function
	f.X = Span(0, 1, len(f.Y))           // domain: equidistant points along X axis
	fmt.Println(
		f.At(0), // f.At(x) evaluates f at x
		f.At(0.25),
		f.At(0.5),
		f.At(0.75),
		f.At(1.0),
	)
	// Output:
	// 0 0.5 1 0.5 0
}

func TestFunction_Area(t *testing.T) {
	tests := []struct {
		name string
		X    []float64
		Y    []float64
		want float64
	}{
		{
			name: "empty",
			want: 0.0,
		},
		{
			name: "simple",
			X:    []float64{0, 1},
			Y:    []float64{0, 1},
			want: 0.5,
		},
		{
			name: "simple+1",
			X:    []float64{0, 1},
			Y:    []float64{1, 2},
			want: 1.5,
		},
		{
			name: "two segments",
			X:    []float64{0, 1, 2},
			Y:    []float64{1, 2, 2},
			want: 3.5,
		},
		{
			name: "three segments",
			X:    []float64{0, 1, 2, 3},
			Y:    []float64{1, 2, 2, 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				X: tt.X,
				Y: tt.Y,
			}
			if got := f.Area(); got != tt.want {
				t.Errorf("Function.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_At(t *testing.T) {
	tests := []struct {
		name string
		X    []float64
		Y    []float64
		x    float64
		want float64
	}{
		{
			name: "simple",
			X:    []float64{0, 1},
			Y:    []float64{0, 1},
			x:    0.5,
			want: 0.5,
		},
		{
			name: "saw(0.25)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.25,
			want: -1,
		},
		{
			name: "saw(0.125)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.125,
			want: -0.5,
		},
		{
			name: "saw(1.0)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    1.0,
			want: 0.0,
		},
		{
			name: "saw(0.0)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.0,
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				X: tt.X,
				Y: tt.Y,
			}
			if got := f.At(tt.x); got != tt.want {
				t.Errorf("Function.At() = %v, want %v", got, tt.want)
			}
		})
	}
}
