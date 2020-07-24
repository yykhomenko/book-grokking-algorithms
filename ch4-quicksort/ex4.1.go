package ch4_quick

func sum(arr []int) int {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return arr[0]
	default:
		return arr[0] + sum(arr[1:])
	}
}
