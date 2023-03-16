package Solution

func Solution(x bool) bool {
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	if root == nil {
		return []int{-1}
	}
	ans := PreOrder(root, voyage)

	for _, val := range ans {
		if val == -1 {
			return []int{-1}
		}
	}

	return ans
}

func PreOrder(root *TreeNode, voyage []int) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0, 10)
	if root.Val != voyage[0] {
		return []int{-1}
	}

	voyage = voyage[1:]
	if root.Left != nil {
		if root.Left.Val == voyage[0] {
			ans = append(ans, PreOrder(root.Left, voyage[1:])...)
		} else if root.Right != nil && root.Right.Val == voyage[0] {
			ans = append(ans, root.Val)
			ans = append(ans, PreOrder(root.Right, voyage[1:])...)
		}
	}
	if root.Right != nil {
		if root.Right.Val == voyage[0] {
			ans = append(ans, PreOrder(root.Right, voyage[1:])...)
		} else if root.Left != nil && root.Left.Val == voyage[0] {
			ans = append(ans, root.Val)
			ans = append(ans, PreOrder(root.Left, voyage[1:])...)
		}
	}

	return ans
}
