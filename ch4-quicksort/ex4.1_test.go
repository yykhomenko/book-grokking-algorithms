package ch4_quick

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid sum of empty",
			args: struct{ arr []int }{arr: []int{}},
			want: 0,
		},
		{
			name: "valid sum of one element",
			args: struct{ arr []int }{arr: []int{42}},
			want: 42,
		},
		{
			name: "valid sum of several elements",
			args: struct{ arr []int }{arr: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.arr); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
