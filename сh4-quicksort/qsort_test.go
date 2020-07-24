package ch4_quick

import (
	"reflect"
	"testing"
)

func Test_qsort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "valid", args: []int{3, 5, 2, 1, 4}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := qsort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("qsort() = %v, want %v", got, tt.want)
			}
		})
	}
}
