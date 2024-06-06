package Solution

import "awesome-golang-algorithm/utils"

func rob(root *utils.TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func rob(root *TreeNode) int {
	robRoot, notRob := getAmt(root)

	return max(robRoot, notRob)
}

func getAmt(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftRob, leftNotRob := getAmt(root.Left)
	rightRob, rightNotRob := getAmt(root.Right)
	robRoot := root.Val + leftNotRob + rightNotRob
	notRob := max(leftRob, leftNotRob) + max(rightRob, rightNotRob)
	return robRoot, notRob
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
