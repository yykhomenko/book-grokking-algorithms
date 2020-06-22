package ch4_fastsort

func length(arr []int) int {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return 1 + length(arr[1:])
	}
}
