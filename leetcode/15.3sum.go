// https://leetcode.cn/problems/3sum/

// 双指针写法

func threeSum(nums []int) (arrays [][]int) {
	sort.Ints(nums)
	// fmt.Println(nums)
	for first := 0; first < len(nums)-2; first++ {
		if nums[first] > 0 {
			// 第一个数大于0，后面的数都大于0，不可能和为0
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			//  第一个数重复，会和上一个产生重复的结果
			continue
		}
		for third := len(nums) - 1; third > first+1; third-- {
			// fmt.Printf("first %d, third %d, nums[first] %d, nums[third] %d\n", first, third, nums[first], nums[third])
			if nums[third] < 0 {
				// 第三个数小于0，后面的数都小于0，不可能和为0
				break
			}
			if third < len(nums)-1 && nums[third] == nums[third+1] {
				// 第三个数重复，会和上一个产生重复的结果
				continue
			}
			secord, ok := binarySearch(-nums[first]-nums[third], first+1, third-1, nums)
			if !ok {
				continue
			}
			arrays = append(arrays, []int{nums[first], nums[secord], nums[third]})
		}
	}
	return
}

func binarySearch(target, start, end int, nums []int) (index int, ok bool) {
	for start <= end {
		index = (start + end) / 2
		if nums[index] == target {
			ok = true
			break
		}
		if nums[index] < target {
			start = index + 1
		} else {
			end = index - 1
		}
	}
	return
}
