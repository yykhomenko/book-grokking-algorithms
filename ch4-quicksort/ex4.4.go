package ch4_quick

func binarySearch(arr []int, item int) int {
	return bSearch(arr, item, 0, len(arr)-1)
}

func bSearch(arr []int, item, lo, hi int) int {
	if hi < lo {
		return -1
	}
	mid := lo + (hi-lo)/2
	switch {
	case arr[mid] > item:
		return bSearch(arr, item, lo, mid-1)
	case arr[mid] < item:
		return bSearch(arr, item, mid+1, hi)
	default:
		return mid
	}
}
