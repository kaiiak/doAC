// https://leetcode.com/problems/climbing-stairs/

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    s := []int{1,2}
    for i := 3; i <= n; i++ {
        next := s[i-2]+s[i-3]
        s = append(s, next)
    }
    return s[n-1]
}