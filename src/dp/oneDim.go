package dp

import (
	"math"
)

func maxOfSlice(s []int) int {
	maxValue := 0
	for _, value := range s {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
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

func MaxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		if prices[i] > minOfSlice(prices[0:i]) {
			dp[i] = prices[i] - minOfSlice(prices[0:i])
		} else {
			dp[i] = 0
		}
	}
	return maxOfSlice(dp)
}
