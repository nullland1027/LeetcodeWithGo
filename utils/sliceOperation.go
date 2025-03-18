package utils

func MaxOfSlice(nums []int) int {
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]
		}
	}
	return maxValue
}
