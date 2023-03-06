package slices

func Index[T comparable](elems []T, v T) int {
	for index, s := range elems {
		if v == s {
			return index
		}
	}
	return -1
}

func Contains[T comparable](elems []T, v T) bool {
	return Index(elems, v) >= 0
}
