package ch4_fastsort

import (
	"math"
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "valid max of empty", args: struct{ arr []int }{arr: []int{}}, want: math.MinInt64},
		{name: "valid max of one", args: struct{ arr []int }{arr: []int{42}}, want: 42},
		{name: "valid max of several", args: struct{ arr []int }{arr: []int{1, 4, 2, 9, 2}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.arr); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
