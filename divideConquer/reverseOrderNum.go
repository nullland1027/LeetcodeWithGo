package divideConquer

import "fmt"

// 暴力法验证

func ReverseOrder(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				count++
			}
		}
	}
	return count
}

// 分治法求逆序数
// 将数组分为左右2部分，逆序数和=左部分+右部分+横跨左右部分

func mergeTwoSortedArrAndCountInverse(left, right []int) ([]int, int) {
	l, r := 0, 0
	arr := make([]int, 0)
	count := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			count += len(left) - l
			arr = append(arr, right[r])
			r++
		} else {
			arr = append(arr, left[l])
			l++
		}
	}
	arr = append(arr, left[l:]...)
	arr = append(arr, right[r:]...)
	return arr, count
}

func mergeSortAndCount(nums []int) ([]int, int) {
	if len(nums) < 2 {
		return nums, 0
	}
	mid := len(nums) / 2
	left, right := nums[:mid], nums[mid:]
	leftSorted, leftCount := mergeSortAndCount(left)
	rightSorted, rightCount := mergeSortAndCount(right)
	sorted, mergeCount := mergeTwoSortedArrAndCountInverse(leftSorted, rightSorted)
	return sorted, mergeCount + leftCount + rightCount
}

func InversePairs(nums []int) int {
	sortedNums, count := mergeSortAndCount(nums)
	fmt.Println(sortedNums)
	return count
}
