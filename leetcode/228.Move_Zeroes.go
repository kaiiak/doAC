// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		cur, next := 0, 1
		for {
			if next >= len(nums) {
				break
			}
			if nums[cur] == 0 && nums[next] != 0 {
				nums[cur], nums[next] = nums[next], nums[cur]
			}
			cur++
			next++
		}
	}
}

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// [0,1,0,3,12]
// [0,0,1]