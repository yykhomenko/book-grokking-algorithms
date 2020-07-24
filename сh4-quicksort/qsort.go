package ch4_quick

func qsort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	pivot := slice[0]
	less := make([]int, 0, len(slice))
	greater := make([]int, 0, len(slice))

	for _, i := range slice[1:] {
		if i <= pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}

	result := append(less, pivot)
	return append(result, greater...)
}
