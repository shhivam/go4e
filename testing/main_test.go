package main

import "testing"

func TestAverage(t *testing.T) {
	type args struct {
		lst []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Something", args: args{[]float64{4.0, 6}}, want: 5.0},
		{name: "Something hi", args: args{[]float64{1.2, 1.3}}, want: 1.25},
		{name: "Something hi 2", args: args{[]float64{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.lst); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}
