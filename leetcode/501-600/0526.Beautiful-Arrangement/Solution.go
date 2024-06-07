package Solution

func countArrangement(n int) int {
	arr := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	res := 0
	usedMap := make(map[int]struct{})
	var dfs func([]int)
	dfs = func(currArr []int) {
		if len(currArr) == n {
			res++
		}
		if len(currArr) > n {
			return
		}
		for i := 0; i < len(arr); i++ {
			if _, ok := usedMap[arr[i]]; ok {
				continue
			}
			if (arr[i]%(len(currArr)+1)) != 0 && (len(currArr)+1)%arr[i] != 0 {
				continue
			}
			usedMap[arr[i]] = struct{}{}
			currArr = append(currArr, arr[i])
			dfs(currArr)
			currArr = currArr[0 : len(currArr)-1]
			delete(usedMap, arr[i])
		}
	}
	dfs([]int{})

	return res
}
