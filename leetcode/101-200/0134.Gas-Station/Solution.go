package Solution

// 注意: 如果从 i 出发不能走完全程到达了j-1, j<长度,那么下次起点应该从 j 开始而不是 i+1,这样会更省时间
func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; i++ {
		if gas[i] < cost[i] {
			continue
		}
		cumulate := gas[i] - cost[i]
		j := i + 1
		for true {
			if j >= l {
				j = 0
			}
			cumulate += gas[j] - cost[j]
			if cumulate < 0 {
				if i > j-1 {
					return -1
				} else {
					i = j - 1
				}
				break
			}
			if j == i {
				return i
			}
			j++
		}
	}

	return -1
}
