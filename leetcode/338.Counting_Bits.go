// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {
    var nums = make([]int, 0, n)
    for i := 0; i <= n; i++ {
        nums = append(nums, bitsCount(i))
    }
    return nums
}

func bitsCount(n int) (count int) {
	for n > 0 {
		if n%2 == 1 {
			count++
			n -= 1
		}
		n /= 2
	}
	return
}