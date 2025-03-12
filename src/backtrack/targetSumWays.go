package backtrack

// 回溯法解决
func FindTargetSumWays(nums []int, target int) int {
	count := 0
	var dfs func(i, runningSum int)
	dfs = func(i, runningSum int) {
		if i >= len(nums) {
			if runningSum == target {
				count++
			}
			return
		}
		originSum := runningSum
		runningSum += nums[i]
		dfs(i+1, runningSum)
		originSum += -nums[i]
		dfs(i+1, originSum)
	}
	dfs(0, 0)
	return count
}
