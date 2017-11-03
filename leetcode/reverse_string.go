// https://leetcode.com/problems/reverse-string

func reverseString(s string) string {
	str := []rune(s)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(s)-1-i] = str[len(str)-i-1], str[i]
	}
	return string(str)
}
