func getConcatenation(nums []int) []int {
	n := len(nums)
	result := make([]int, n*2)

	for i := 0; i < n; i++ {
		result[i] = nums[i]
		result[i+n] = nums[i]
	}
	return result
}