package filterOdd

func filterOdd(nums []int) []int {
	var result = make([]int, 0, len(nums))
	for _, num := range nums {
		if num%2 != 0 {
			result = append(result, num)
		}
	}
	return result
}
