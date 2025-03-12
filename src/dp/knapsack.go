package dp

// ZeroOneKnapsack 0/1背包问题
/**
1. 明确dp数组的含义
行：按次序依次考虑的物品序号。例如，从考虑0个物品开始，再到考虑第1个物品，考虑第2个物品...一直到考虑所有物品
列：从背包的容量为0开始，逐渐往上递增。0, 1, 2, ..., 一直到cap
*/
func ZeroOneKnapsack(capacity int, values, weights []int) int {
	// 定义dp数组
	dp := make([][]int, len(values)+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	for item := 1; item < len(values)+1; item++ { // 外层遍历：对于所有的物品序号，考虑0个物品；考虑前1个物品...
		for capa := 1; capa < capacity+1; capa++ { // 内层遍历：对于从0开始直到背包容量的上限
			// 需要判断当前物品的重量是否超过了当前容量（重要）
			if weights[item-1] > capa { // 当前物品无法放入背包，太重了
				dp[item][capa] = dp[item-1][capa]
			} else {
				// 选择当前物品的递推关系：回到背包容量减去当前物品容量的上一步，那时的最优解再加上当前物品的价值，即为现在的最优解
				selectCurrentItem := dp[item-1][capa-weights[item-1]] + values[item-1]
				// 不选择当前物品：即价值和考虑前一个物品的那一轮时和当前的背包容量一样时的那个值
				notSelectCurrentItem := dp[item-1][capa]
				// 取二者的最大值
				dp[item][capa] = max(selectCurrentItem, notSelectCurrentItem)
			}
		}
	}
	// 打印dp数组
	//for _, arr := range dp {
	//	for _, val := range arr {
	//		fmt.Print(val, " ")
	//	}
	//	fmt.Println()
	//}
	return dp[len(values)][capacity]
}
