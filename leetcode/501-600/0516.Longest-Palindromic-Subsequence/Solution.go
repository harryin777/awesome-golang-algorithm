package Solution

import (
	"awesome-golang-algorithm/utils"
	"fmt"
)

func Solution(x bool) bool {
	return x
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		fmt.Printf("i :%v  ", i)
		for j := i + 1; j < n; j++ {
			fmt.Printf("j :%v ", j)
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
		fmt.Println()
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {

	return nil
}

// 前序遍历
func pre(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	order := []int{root.Val}
	if root.Left != nil {
		order = append(order, pre(root.Left)...)
	}
	if root.Right != nil {
		order = append(order, pre(root.Right)...)
	}

	return order
}

func Mid(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	order := make([]int, 0, 10)
	if root.Left != nil {
		order = append(order, pre(root.Left)...)
	}
	order = append(order, root.Val)
	if root.Right != nil {
		order = append(order, pre(root.Right)...)
	}

	return order
}
