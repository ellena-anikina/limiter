package filter

func filter(predicate func(int) bool, iterable []int) []int {
	var result []int

	for _, v := range iterable {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
