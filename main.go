package main

import (
	"fmt"
	"leetcode/divideConquer"
)

func main() {
	var nums []int
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		nums = append(nums, num)
	}
	fmt.Println(divideConquer.ReverseOrder(nums))
	fmt.Println(divideConquer.InversePairs(nums))
}
