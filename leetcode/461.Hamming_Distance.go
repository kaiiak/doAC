// https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
	x ^= y
	count := 0 // 次数
	for x != 0 {
		if x%2 == 1 {
			count++
			x -= 1
		}
		x /= 2
	}
	return count
}

func hammingDistance(x int, y int) int {
	x ^= y
	count := 0 // 次数
	for x != 0 {
		if x%2 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func hammingDistance(x int, y int) int {
	x = x ^ y
	var count = 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}