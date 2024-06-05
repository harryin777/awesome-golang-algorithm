package Solution

func canIWin(maxChoosableInteger, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	dp := make([]int8, 1<<maxChoosableInteger)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(int, int) int8
	dfs = func(usedNum, curTot int) (res int8) {
		dv := &dp[usedNum]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := 0; i < maxChoosableInteger; i++ {
			// 这里检查的是 i+1 位，因为二进制是从0开始的，但是实际上需要检查的数字最低是从1开始的
			// 用位运算去检查是否设置了值，都有这个问题，比如想检查2是否被设置或者被用过，那么需要向右移动1位数字
			// 为什么想要检查2但是向右移动1位呢，因为在二进制表示里，1表示为0，如果想要检查1，那么右移0位，还是原先的数字
			if usedNum>>i&1 == 0 && (curTot+i+1 >= desiredTotal || dfs(usedNum|1<<i, curTot+i+1) == 0) {
				return 1
			}
		}
		return
	}
	return dfs(0, 0) == 1
}
