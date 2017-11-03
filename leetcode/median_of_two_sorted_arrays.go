// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var lArray, sArray []int
	if len(nums1) > len(nums2) {
		lArray = nums1
		sArray = nums2
	} else {
		lArray = nums2
		sArray = nums1
	}
	var i, j int
	var result []int
	for {
		if j == len(sArray) {
			result = append(result, lArray[i:]...)
			break
		}
		if i == len(lArray) {
			result = append(result, sArray[j:]...)
			break
		}
		if sArray[j] > lArray[i] {
			result = append(result, lArray[i])
			i++
		} else {
			result = append(result, sArray[j])
			j++
		}
	}
	var median float64
	if len(result)%2 == 1 {
		median = float64(result[len(result)/2])
	} else {
		median = float64(result[len(result)/2]+result[len(result)/2-1]) / 2
	}
	return median
}
