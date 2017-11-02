// https://leetcode.com/problems/array-partition-i

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var result int
	for i := 0; i < len(nums)/2; i++ {
		result += nums[i*2]
	}
	return result
}
