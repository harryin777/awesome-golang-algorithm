package Solution

var (
	ans   [][]int
	visit []bool
)

func permute(nums []int) [][]int {
	ans = [][]int{}
	visit = make([]bool, len(nums))
	dfs(nums, []int{})
	return ans
}

func dfs(nums []int, curr []int) {
	if len(curr) == len(nums) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		ans = append(ans, temp)
		return
	}

	for idx, num := range nums {
		if !visit[idx] {
			curr = append(curr, num)
			visit[idx] = true

			dfs(nums, curr)

			visit[idx] = false
			curr = curr[:len(curr)-1]
		}
	}
}

func DFS(nums []int) [][]int {
	var res [][]int
	for _, val := range nums {
		data := DFS2(nums, []int{val})
		res = append(res, data...)
	}

	return res
}

func DFS2(nums []int, curr []int) (res [][]int) {
	if len(nums) == len(curr) {

		return [][]int{curr}
	}

	var tmp []int
	for _, val := range nums {
		exists := false
		for _, innerVal := range curr {
			if val == innerVal {
				exists = true
			}
		}
		if !exists {
			tmp = append(tmp, val)
		}
	}

	for _, val := range tmp {
		curr = append(curr, val)
		data := DFS2(nums, curr)
		res = append(res, data...)
		curr = curr[:len(curr)-1]
	}

	return
}
