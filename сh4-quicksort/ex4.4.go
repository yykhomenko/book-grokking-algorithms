package ch4_fastsort

func binarySearch(arr []int, element int) int {
	return bSearch(arr, element, 0, len(arr)-1)
}

func bSearch(arr []int, item, lo, hi int) int {
	if hi < lo {
		return -1
	}
	mid := (hi - lo) / 2
	switch {
	case arr[mid] > item:
		return bSearch(arr, item, lo, mid)
	case arr[mid] < item:
		return bSearch(arr, item, mid+1, hi)
	default:
		return mid
	}
}
