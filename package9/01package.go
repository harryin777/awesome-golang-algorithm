package package9

// package01 01 背包
func package01(weight, num int, values, weights []int) int {
	dp := make([][]int, num+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, weight+1)
	}

	for i := 1; i <= num; i++ {
		for j := 1; j <= weight; j++ {
			dp[i][j] = dp[i-1][j]
			if weights[i-1] < j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}

	return dp[num][weight]
}

/*
package01V2 注意这里为什么内层循环的时候是从后往前。整个优化的过程实际上是源于最开始的二维状态转义方程
dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
注意到，无论是当前物品i的重量是不是大于了当前的背包容积，dp[i][j] 始终和dp[i-1][j] 有关系，重点是这个 i-1，也就是说当前的第i行数组
需要用到上一行也就是i-1行的数组
那么可以优化，也就是我们只需要记录上一行 i-1 行的状态，而不需要记录 i-2， i-3 也就是最外层的一个维度的数组都不用记录，只用记录当前行之前的一行就可以
所以可以用一个一位数组，当 i -1 行所有的位置都算完以后，当前的一维数组记录的是 i -1 行的数据，当需要计算第 i 行的时候，注意
状态转移方程变成了 dp[j] = max(dp[j], dp[j-weights[i]]+values[i])， 在这个一维数组的任意位置j，它只和j-weights[i]这个位置有关系，因为weights中的数据必然大于0
所以 j-weights[i] 必然比 j 小，也就是这个位置一定在j的前面，整体算法我们虽然优化成了一维数组，但是整体的思路还是基于上面的一句话，并不需要记录整个 最外层的i数组
只需要记录 i-1 行的数据，而此时我们真正需要用的是 i-1 行的 j-weights[i] 位置。所以如果按照从小到大的顺序去计算，我们会用i行的j-weights[i]替换掉i-1 行的 j-weights[i]
唯一的办法就是从后往前遍历，这样当前的一维数组，前半部分保留的是i-1行的数据，后半部分保留的是i行的数据
*/
func package01V2(weight, num int, values, weights []int) int {
	dp := make([]int, weight+1)

	for i := 0; i < num; i++ {
		for j := weight; j >= 0; j-- {
			dp[j] = dp[j]
			if j > weights[i] {
				dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
			}
		}
	}

	return dp[weight]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
