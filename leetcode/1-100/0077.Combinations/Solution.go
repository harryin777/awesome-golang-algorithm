package Solution

func Solution(x bool) bool {
	return x
}

func combine(n int, k int) [][]int {
	var res [][]int
	var dfs func(int, []int, int, map[int]bool)
	used := make(map[int]bool)
	dfs = func(n int, path []int, pos int, used map[int]bool) {
		if len(path) == k {
			data := make([]int, 0, len(path))
			for _, v := range path {
				data = append(data, v)
			}
			res = append(res, data)
			return
		}

		levelMap := make(map[int]bool)
		for i := pos; i <= n; i++ {
			if levelMap[i] || used[i] {
				continue
			}
			used[i] = true
			levelMap[i] = true
			path = append(path, i)
			dfs(n, path, i, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs(n, []int{}, 1, used)

	return res
}
