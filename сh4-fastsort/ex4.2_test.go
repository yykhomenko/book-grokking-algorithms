package ch4_fastsort

import "testing"

func Test_length(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid length of empty",
			args: struct{ arr []int }{arr: []int{}},
			want: 0,
		},
		{
			name: "valid length of one element",
			args: struct{ arr []int }{arr: []int{42}},
			want: 1,
		},
		{
			name: "valid length of several elements",
			args: struct{ arr []int }{arr: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := length(tt.args.arr); got != tt.want {
				t.Errorf("length() = %v, want %v", got, tt.want)
			}
		})
	}
}
