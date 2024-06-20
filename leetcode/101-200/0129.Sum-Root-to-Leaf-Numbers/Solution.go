package Solution

import (
	"fmt"
	"strconv"
)

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

// 这个题很有意思，在递归返回的时候会自动删除数组的最后一个元素
// gpt 是这么解释的 在递归返回时，虽然没有显式地删除最后一个元素，但返回后的curStr是传递到当前递归调用时的切片引用，它的长度和内容是进入递归调用前的状态。
// 我觉得有点道理，从打印的地址信息看，如果最开始调用dfs的时候传递一个空slice，那么第一次递归和后序在左右子树的slice不是同一个
// 但是如果提前分配了slice的容量，那么就会是一个slice，并且右子树遍历的时候会重置掉左子树追加进去的元素，这里是写法问题，应该在最后加入到strArr的时候做一个深拷贝

func sumNumbers2(root *TreeNode) int {
	strArr := make([][]int, 0, 10)
	var dfs func(*TreeNode, []int)
	dfs = func(root *TreeNode, curStr []int) {
		if root == nil {
			return
		}
		curStr = append(curStr, root.Val)
		fmt.Printf("add : %p \n", curStr)

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
	curStr := make([]int, 0, 10)
	dfs(root, curStr)
	res := 0
	for _, str := range strArr {
		valStr := ""
		for i := 0; i < len(str); i++ {
			valStr += strconv.Itoa(str[i])
		}
		val, _ := strconv.ParseInt(valStr, 10, 64)
		res += int(val)
	}

	return res
}
