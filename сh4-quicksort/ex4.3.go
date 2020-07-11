package ch4_fastsort

import "math"

func max(arr []int) int {
	return maxrec(arr, math.MinInt64)
}

func maxrec(arr []int, curmax int) int {
	switch len(arr) {
	case 0:
		return curmax
	case 1:
		return maxInt(arr[0], curmax)
	default:
		return maxrec(arr[1:], maxInt(arr[0], curmax))
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
