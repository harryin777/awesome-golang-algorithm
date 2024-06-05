package Solution

import "awesome-golang-algorithm/utils"

func rob(root *utils.TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *utils.TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func rob2(root *utils.TreeNode) int {
	return getAmt(root)
}

func getAmt(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	leftAmt := getAmt(root.Left)
	rightAmt := getAmt(root.Right)

	if root.Val > (leftAmt + rightAmt) {
		return root.Val
	}

	return leftAmt + rightAmt
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
