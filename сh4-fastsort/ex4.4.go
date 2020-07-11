package ch4_fastsort

func binarySearch(arr []int, element int) int {
	return bSearch(arr, element, 0, len(arr)-1)
}

func bSearch(arr []int, item, low, high int) int {

	if len(arr) == 0 {
		return -1
	}

	mid := (high - low) / 2
	guess := arr[mid]

	switch {
	case guess == item:
		return mid
	case guess > item:
		high = mid - 1
		return bSearch(arr, item, low, high)
	default:
		low := mid + 1
		return bSearch(arr, item, low, high)
	}
}
