package main

import (
	"fmt"
	"leetcode/src/basicSort"
	"math"
	"math/rand"
	"sort"
)

func maxScoreSightseeingPair(values []int) int {
	var maxScore int
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if maxScore < values[i]+values[j]+i-j {
				maxScore = values[i] + values[j] + i - j
			}
		}
	}
	return maxScore
}

func HeapifyOneNodeInLoop(arr []int, startNodeIdx int) []int { // 索引0不存放元素
	currentIdx := startNodeIdx
	childNodeIdx := 2 * currentIdx

	arr[0] = arr[currentIdx]
	for childNodeIdx < len(arr) {
		if childNodeIdx < len(arr) && arr[childNodeIdx] < arr[childNodeIdx+1] { // This judgement finds the max value between left and right
			childNodeIdx++
		}
		if arr[0] < arr[childNodeIdx] { //
			arr[currentIdx] = arr[childNodeIdx]
			currentIdx = childNodeIdx
			childNodeIdx = 2 * currentIdx
		} else {
			break
		}
	}
	arr[currentIdx] = arr[0]
	return arr
}

func minOfSlice(s []int) int {
	minValue := math.MaxInt
	for _, value := range s {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	isExistedMap := make(map[[3]int]struct{})
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r] < -nums[i] { // left pointer should move right
				l++
			} else if nums[l]+nums[r] < -nums[i] { // right pointer should move left
				r--
			} else {
				threePair := [3]int{nums[i], nums[l], nums[r]}
				if _, exist := isExistedMap[threePair]; !exist {
					isExistedMap[threePair] = struct{}{}
					ans = append(ans, threePair[:])
				}
				l++
				r--
			}
		}
		fmt.Println(ans)
	}
	return ans
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	if numRows == 1 {
		return res
	}
	if numRows == 2 {
		res[1] = []int{1, 1}
		return res
	}
	res[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		curRow := make([]int, i+1)
		curRow[0] = 1
		curRow[i] = 1
		for j := 1; j < i; j++ {
			curRow[j] = res[i-1][j-1] + res[i-1][j]
		}

		res = append(res, curRow)
	}
	return res
}

func main() {
	randomSlice := make([]int, 0)
	for i := 0; i < 20; i++ {
		randomSlice = append(randomSlice, rand.Intn(100))
	}
	fmt.Println(randomSlice)
	fmt.Println(basicSort.MergeSort(randomSlice))
}
