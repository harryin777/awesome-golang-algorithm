package Solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0, 10)
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := make([]*Node, 0, 10)
		lll := len(queue)
		for i := 0; i < lll; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		for i := 0; i < len(tmp)-1; i++ {
			tmp[i].Next = tmp[i+1]
		}
		tmp[len(tmp)-1].Next = nil
	}

	return root
}
