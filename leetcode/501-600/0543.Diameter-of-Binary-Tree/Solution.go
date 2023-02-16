package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func diameterOfBinaryTree_1(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return ans
}

func LevelTraverse(root *TreeNode) []int {

	var arr []int
	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root)
	for len(queue) != 0 {
		for _, val := range queue {
			arr = append(arr, val.Val)
			queue = append(queue, val.Left, val.Right)
		}
		queue = queue[1:]
	}

	return arr
}

func largestValues(root *TreeNode) []int {
	var res []int
	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root)
	for len(queue) != 0 {
		var maxmax int
		count := 0
		for _, q := range queue {
			maxmax = max(maxmax, q.Val)
			queue = queue[1:]
			if q.Left != nil {
				queue = append(queue, q.Left)
				count++
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
				count++
			}
		}
		res = append(res, maxmax)

	}

	return res
}
