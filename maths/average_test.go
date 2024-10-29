package maths

import "testing"

func TestAverage(t *testing.T) {
	type args struct {
		sum    float64
		length float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test with positive integers",
			args: args{
				sum:    20,
				length: 4,
			},
			want: 5,
		},

		{
			name: "Test with negative integers",
			args: args{
				sum:    -20,
				length: 4,
			},
			want: -5,
		},

		{
			name: "Test with one element",
			args: args{
				sum:    17,
				length: 1,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.sum, tt.args.length); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}
