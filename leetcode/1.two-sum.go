// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	indexMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if p, ok := indexMap[target-nums[i]]; ok {
			return []int{p, i}
		}
		indexMap[nums[i]] = i
	}
	return nil
}
