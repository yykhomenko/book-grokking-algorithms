package ch4_quick

import (
	"reflect"
	"testing"
)

func TestQsort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "empty", args: []int{}, want: []int{}},
		{name: "one", args: []int{1}, want: []int{1}},
		{name: "two", args: []int{2, 1}, want: []int{1, 2}},
		{name: "some", args: []int{3, 5, 2, 1, 4}, want: []int{1, 2, 3, 4, 5}},
		{name: "more", args: []int{-1, 50, -6, 80, 90, 11, 14, 17, -20, 23, 24}, want: []int{-20, -6, -1, 11, 14, 17, 23, 24, 50, 80, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := qsort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("qsort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQsort(b *testing.B) {
	arr := []int{-1, 50, -6, 80, 90, 11, 14, 17, -20, 23, 24}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		qsort(arr)
	}
	b.StopTimer()
}
