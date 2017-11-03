// https://leetcode.com/problems/reverse-words-in-a-string-iii


func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i := 0; i < len(strs); i++ {
		strs[i] = reverseString(strs[i])
	}
	return strings.Join(strs, " ")
}

func reverseString(ss string) string {
	str := []byte(ss)
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-1-i] = str[n-i-1], str[i]
	}
	return string(str)
}