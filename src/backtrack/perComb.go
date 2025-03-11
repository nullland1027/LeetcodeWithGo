package backtrack

import (
	"fmt"
	"strconv"
)

var mapping = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// digits: 用户输入的数字字符串
// index:  表明当前递归到哪个字符了
func backTrackForComb(digits string, index int, leafString string, finalAns []string) []string {
	// 终止条件
	if index == len(digits) { // 已经超出了数字串，表明递归结束
		finalAns = append(finalAns, leafString)
		return finalAns
	}
	// 开始for循环遍历每个数字字符的选择
	currentNum, _ := strconv.Atoi(string(digits[index]))
	optionLetters := mapping[currentNum]
	for _, letter := range optionLetters {
		leafString = leafString + string(letter)
		finalAns = backTrackForComb(digits, index+1, leafString, finalAns)
		leafString = leafString[:len(leafString)-1]
	}
	return finalAns
}

func letterCombinations(digits string) []string {
	leafString := ""              // 叶子节点的字符串，需要添加到最终结果中
	finalAns := make([]string, 0) // 最终的答案，需要返回，用递归函数来修改
	index := 0
	return backTrackForComb(digits, index, leafString, finalAns)
}

func backTrackForPer(nums []int, startIdx int, leafAns []int, finalAns [][]int) [][]int {
	// 基本条件
	if len(leafAns) == len(nums) {
		finalAns = append(finalAns, leafAns)
		return finalAns
	}
	// 开始递归
	for i := startIdx; i < len(nums); i++ {
		leafAns = append(leafAns, nums[i])
		finalAns = backTrackForPer(nums, startIdx+1, leafAns, finalAns)
		leafAns = leafAns[:len(leafAns)-1]
	}
	fmt.Println(finalAns)
	return finalAns
}

func Permute(nums []int) [][]int {
	startIdx := 0
	leafAns := make([]int, 0)
	finalAns := make([][]int, 0)
	return backTrackForPer(nums, startIdx, leafAns, finalAns)
}
