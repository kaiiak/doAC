/*
 * @lc app=leetcode id=1431 lang=golang
 *
 * [1431] Kids With the Greatest Number of Candies
 */

 func kidsWithCandies(candies []int, extraCandies int) []bool {
	var (
		result   = make([]bool, 0, len(candies))
		maxCandy int
	)
	for i := 0; i < len(candies); i++ {
		if maxCandy < candies[i] {
			maxCandy = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		result = append(result, candies[i]+extraCandies >= maxCandy)
	}
	return result
}