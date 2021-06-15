func makeEqual(words []string) bool {
	bytesCount := map[byte]int{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			bytesCount[words[i][j]]++
		}
	}
	for _, count := range bytesCount {
		if count%len(words) != 0 {
			return false
		}
	}
	return true
}