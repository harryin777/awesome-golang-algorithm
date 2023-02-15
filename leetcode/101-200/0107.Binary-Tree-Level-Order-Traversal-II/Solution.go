package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		temp := make([]*TreeNode, 0)
		tempVals := make([]int, 0)
		for len(stack) > 0 {
			node := stack[0]
			tempVals = append(tempVals, node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
			stack = stack[1:]
		}
		stack = append(stack, temp...)
		res = append(res, tempVals)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

var depth int

var res int

func maxDepth(root *TreeNode) int {
	pre(root)
	return res
}

func pre(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	depth++
	order := []int{root.Val}
	if root.Left == nil && root.Right == nil {
		res = max(res, depth)
	}
	order = append(order, pre(root.Left)...)
	order = append(order, pre(root.Right)...)
	depth--
	return order
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
