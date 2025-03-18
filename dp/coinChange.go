package dp

import "fmt"

func CoinChange(coins []int, amount int) int {
	Max := amount + 1
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = Max
	}
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
