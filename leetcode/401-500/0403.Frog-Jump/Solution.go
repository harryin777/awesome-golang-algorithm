package Solution

func canCross(stones []int) bool {
	l := len(stones)
	if l == 1 {
		return true
	}

	dp := make([][]bool, l)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, l)
	}
	dp[0][0] = true

	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == l-1 && dp[i][k] {
				return true
			}
		}
	}

	// for _, val := range dp[l-1]{
	//     if val{
	//         return true
	//     }
	// }

	return false
}
