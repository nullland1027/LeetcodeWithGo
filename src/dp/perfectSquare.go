package dp

import (
	"math"
)

func isPerfectSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	nearestInt := int(math.Round(sqrt))
	return nearestInt*nearestInt == n
}

func NumSquares(n int) int {
	dp := make([]int, n+1)
	// 根据4完全平方数定理，将所有值更新为4
	for i := 0; i < n+1; i++ {
		dp[i] = 4
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		if isPerfectSquare(i) {
			dp[i] = 1
		} else {
			minValue := 4
			for s := 1; s < i; s++ {
				if dp[i-s] < minValue && isPerfectSquare(s) {
					minValue = dp[i-s]
				}
			}
			dp[i] = min(1+minValue, dp[i])
		}
	}
	return dp[n]
}
