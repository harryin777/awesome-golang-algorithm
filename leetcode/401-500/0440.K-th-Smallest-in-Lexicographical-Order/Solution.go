package Solution

import (
	"strconv"
)

func findKthNumber(n int, k int) int {
	cur := 1
	k--

	for k > 0 {
		count := prefix(cur, n)

		if count <= k {
			cur++
			k -= count
		} else {
			k--
			cur *= 10
		}
	}

	return cur
}

func prefix(prefix, n int) int {
	first := prefix
	next := prefix + 1
	res := 0
	for first <= n {
		res += min(n+1, next) - first
		first *= 10
		next *= 10
	}

	return res
}

// 可以，但是超时了
func findKthNumber2(n int, k int) int {
	arr := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > n-k; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr[0]

}

func heapify(arr []int, n, i int) {
	min := i
	lson := 2*i + 1
	rson := 2*i + 2

	minStr := strconv.Itoa(arr[min])
	if lson < n {
		lsonStr := strconv.Itoa(arr[lson])
		if lsonStr < minStr {
			min = lson
			minStr = lsonStr
		}
	}
	if rson < n {
		rsonStr := strconv.Itoa(arr[rson])
		if rsonStr < minStr {
			min = rson
		}

	}
	if min != i {
		arr[i], arr[min] = arr[min], arr[i]
		heapify(arr, n, min)
	}

}

// 包含了当前节点在内的，以 cur 为根节点的字典树到 n 还有多少个元素
func getSteps(cur, n int) (steps int) {
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber3(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		// 如果以 cur 为根节点的字典树都不够抵达 k ，因为cur是从1开始，那么肯定是前 steps个最小的，那么就到下一个字典树
		if steps <= k {
			k -= steps
			cur++
		} else {
			// 如果可以抵达 k，那么就往下一次走
			cur *= 10
			k--
		}
	}
	return cur
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
