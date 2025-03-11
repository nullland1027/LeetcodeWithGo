package dataStructure

import "fmt"

func GetAnSlice(length, initValue int) []int {
	ls := make([]int, length)
	for i := 0; i < length; i++ {
		ls[i] = initValue
	}
	return ls
}

func Rotate(nums []int, k int) { // Timeout
	k = k % len(nums)
	for i := 0; i < k; i++ { // 执行次数
		// get the last ele
		last := nums[len(nums)-1]
		for j := len(nums) - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = last
	}
	fmt.Println(nums)
}

func Rotate2(nums []int, k int) {
	k = k % len(nums)
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums)-k; i++ {
		newNums[k+i] = nums[i]
	}
	for i := 0; i < k; i++ {
		newNums[i] = nums[len(nums)-k+i]
	}
	fmt.Println(newNums)
}
