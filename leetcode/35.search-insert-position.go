// https://leetcode.cn/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	var start, end, mid = 0, len(nums) - 1, len(nums) / 2

	for start <= end { // 一个元素时，start == end
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// 没有查到,此时start >= end
	// 小于全部元素时,start = 0,end = -1
	// 大于全部元素时,start = len(nums), end = len(nums)=-1

	return start
}
