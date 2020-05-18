/*
 * @lc app=leetcode id=665 lang=golang
 *
 * [665] Non-decreasing Array
 */

// @lc code=start
func checkPossibility(nums []int) bool {
    var (
        doubleNums = make([][]int, 0, len(nums))
        dn = make([]int,0, len(nums))
    )
    if len(nums) <= 2 {
        return true
    }
    
    for i:= 0; i < len(nums) -1; i++ {
        if nums[i] > nums[i+1] {
            dn = append(dn, nums[i])
            doubleNums = append(doubleNums, dn)
            dn = make([]int,0, len(nums))
            continue
        }
        dn = append(dn, nums[i])
    }
    doubleNums = append(doubleNums, dn)
    if len(doubleNums) > 2{
        return false
    }
    if len(doubleNums) < 2 {
        return true
    }
    if len(doubleNums[0]) < 2 || len(doubleNums[1]) == 0 {
        return true
    }
    if doubleNums[0][len(doubleNums[0])-2] > doubleNums[1][0] {
        return false
    }
    return true
}
// @lc code=end

