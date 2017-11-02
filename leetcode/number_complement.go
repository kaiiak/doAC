// https://leetcode.com/problems/number-complement

func findComplement(num int) int {
	i := strconv.FormatInt(int64((num)), 2)
	var j string
	for _, v := range i {
		if v == rune('0') {
			j += "1"
		} else {
			j += "0"
		}
	}
	fmt.Println(j)
	c, _ := strconv.ParseInt(j, 2, 32)
	return int(c)
}