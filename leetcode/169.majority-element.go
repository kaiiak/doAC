// https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
    counts := map[int]int{}
    for i := 0; i < len(nums); i ++ {
        c := counts[nums[i]]
        c++
        counts[nums[i]] = c
    }
    max, count := nums[0], 0
    for v, c := range counts {
        if c > count {
            max = v
            count = c
        }
    }
    return max
}

// 摩尔投票法
func majorityElement(nums []int) int {
    var (
        major, count int
    )
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            major = nums[i]
        }
        if major == nums[i] {
            count++
        } else {
            count--
        }
    }
    return major
}