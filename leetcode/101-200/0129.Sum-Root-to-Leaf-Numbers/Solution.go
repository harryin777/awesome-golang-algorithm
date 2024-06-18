package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func sumNumbers_1(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}
		sum = sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		left := dfs(node.Left, sum)
		right := dfs(node.Right, sum)
		return left + right

	}
	return dfs(root, 0)
}

type pair struct {
	node *TreeNode
	num  int
}

func sumNumbers_2(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	queue := []pair{{root, root.Val}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		left, right, num := p.node.Left, p.node.Right, p.num
		if left == nil && right == nil {
			sum += num
		} else {
			if left != nil {
				queue = append(queue, pair{left, num*10 + left.Val})
			}
			if right != nil {
				queue = append(queue, pair{right, num*10 + right.Val})
			}
		}
	}
	return
}

func sumNumbers2(root *TreeNode) int {
	strArr := make([][]int, 0, 10)
	var dfs func(*TreeNode, []int)
	dfs = func(root *TreeNode, curStr []int) {
		if root == nil {
			return
		}
		curStr = append(curStr, root.Val)

		if root.Left == nil && root.Right == nil {
			strArr = append(strArr, curStr)
			return
		}
		//defer func() {
		//	curStr = curStr[:len(curStr)-1]
		//	fmt.Println(1)
		//}()

		dfs(root.Left, curStr)
		dfs(root.Right, curStr)
	}
	dfs(root, []int{})
	res := 0
	for _, str := range strArr {
		_ = str
	}

	return res
}
