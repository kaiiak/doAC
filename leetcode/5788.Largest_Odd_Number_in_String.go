// https://leetcode-cn.com/problems/largest-odd-number-in-string/
func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i]%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
