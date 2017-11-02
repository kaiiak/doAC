func hammingDistance(x int, y int) int {
	z := x ^ y
	var count int
	for _, v := range strconv.FormatInt(int64(z), 2) {
		if v == rune('1') {
			count++
		}
	}
	return count
}