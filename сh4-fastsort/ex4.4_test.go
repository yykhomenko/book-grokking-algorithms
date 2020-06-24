package ch4_fastsort

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid search of empty",
			args: struct {
				arr  []int
				item int
			}{
				arr:  []int{},
				item: 0,
			},
			want: -1,
		},
		{
			name: "valid search of one",
			args: struct {
				arr  []int
				item int
			}{
				arr:  []int{42},
				item: 42,
			},
			want: 0,
		},
		{
			name: "valid search of several",
			args: struct {
				arr  []int
				item int
			}{
				arr:  []int{1, 2, 3, 4, 5},
				item: 3,
			},
			want: 2,
		},
		// {
		// 	name: "valid search of several - item not exist",
		// 	args: struct {
		// 		arr  []int
		// 		item int
		// 	}{
		// 		arr:  []int{1, 2, 3, 4, 5},
		// 		item: 42,
		// 	},
		// 	want: -1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.item); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
