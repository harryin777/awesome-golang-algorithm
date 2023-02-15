package utils

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func MakePreTreeNode(arr []int) *TreeNode {
	var t *TreeNode
	for i := 0; i < len(arr); i++ {
		t = insert(t, arr[i])
	}
	return t
}
func insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{nil, v, nil}
	}
	if v == t.Val {
		return t
	}
	if v < t.Val {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}
