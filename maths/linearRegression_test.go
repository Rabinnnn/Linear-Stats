package maths

import (
	"testing"
)

func Test_linearRegression(t *testing.T) {
	type args struct {
		xaxis []float64
		yaxis []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "basic linear relationship",
			args: args{
				xaxis: []float64{0, 1, 2, 3, 4},
				yaxis: []float64{2, 4, 6, 8, 10},
			},
			want:  2, // Expected slope b
			want1: 2, // Expected intercept a
		},
		{
			name: "linear relationship with intercept",
			args: args{
				xaxis: []float64{0, 1, 2, 3, 4},
				yaxis: []float64{3, 5, 7, 9, 11},
			},
			want:  2, // Expected slope b
			want1: 3, // Expected intercept a
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LinearRegression(tt.args.xaxis, tt.args.yaxis)
			if got != tt.want {
				t.Errorf("linearRegression() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("linearRegression() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_sumyax(t *testing.T) {
	type args struct {
		yaxis []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic linear relationship",
			args: args{
				yaxis: []float64{2, 4, 6, 8, 10},
			},
			want: 30, // Expected summation of y axis
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumYaxis(tt.args.yaxis)
			if got != tt.want {
				t.Errorf("linearRegression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumXaxis(t *testing.T) {
	type args struct {
		xaxis []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic linear relationship",
			args: args{
				xaxis: []float64{2, 4, 6, 8, 10},
			},
			want: 30, // Expected summation of x axis
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumXaxis(tt.args.xaxis)
			if got != tt.want {
				t.Errorf("linearRegression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumXSquares(t *testing.T) {
	type args struct {
		xaxis []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic linear relationship",
			args: args{
				xaxis: []float64{2, 4, 6, 8, 10},
			},
			want: 220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumXSquares(tt.args.xaxis)
			if got != tt.want {
				t.Errorf("linearRegression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_XYSumProduct(t *testing.T) {
	type args struct {
		xaxis []float64
		yaxis []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "basic linear relationship",
			args: args{
				xaxis: []float64{1, 2, 3, 4, 5},
				yaxis: []float64{2, 4, 6, 8, 10},
			},
			want: 110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := XYSumProduct(tt.args.xaxis, tt.args.yaxis)
			if got != tt.want {
				t.Errorf("linearRegression() got = %v, want %v", got, tt.want)
			}
		})
	}
}
