func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        if nums[abs(nums[i])-1] > 0 {
            nums[abs(nums[i])-1] *= -1
        }
    }
    disappeared := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            disappeared = append(disappeared, i+1)
        }
    }
    return disappeared
}

func abs(i int) int {
    if i < 0 {
        return i*-1
    }
    return i
}