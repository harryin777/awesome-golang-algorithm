package Solution

import "fmt"

//	直接扫描
func searchRange(nums []int, target int) []int {
	targetRange := []int{-1, -1}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			targetRange[0] = i
			break
		}
	}
	if targetRange[0] == -1 {
		return targetRange
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == target {
			targetRange[1] = j
			break
		}
	}
	return targetRange
}

func searchRange2(nums []int, target int) []int {
	targetRange := []int{-1, -1}
	leftIndex := extremeInsertionIndex(nums, target, true)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return targetRange
	}

	targetRange[0] = leftIndex
	targetRange[1] = extremeInsertionIndex(nums, target, false) - 1

	return targetRange
}

func extremeInsertionIndex(nums []int, target int, left bool) int {
	lo := 0
	hi := len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > target || left && target == nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// 这个答案没问题，请举出超时了
func searchRange3(nums []int, target int) []int {
	minVal, maxVal := int(^uint(0)>>1), -1
	var search func(int, int)
	search = func(left int, right int) {
		if left > right {
			return
		}
		point := (right-left)/2 + left
		fmt.Println(point)
		if nums[point] == target {
			minVal = min(minVal, point)
			maxVal = max(maxVal, point)

		}
		search(0, point-1)
		search(point+1, right)
	}

	search(0, len(nums)-1)

	return []int{minVal, maxVal}
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func searchRange4(nums []int, target int) []int {
	minVal, maxVal := int(^uint(0)>>1), -1
	left, right := 0, len(nums)-1
	for left > right {
		if left > right {
			break
		}
		point := (right-left)/2 + left
		fmt.Println(point)
		if nums[point] == target {
			minVal = min(minVal, point)
			maxVal = max(maxVal, point)
		}
		left = point
		right--
	}

	return []int{minVal, maxVal}
}

func searchRange6(nums []int, target int) []int {
	minOne := BinarySearch(nums, target)
	maxOne := BinarySearch(nums, target+1) - 1

	if minOne == len(nums)-1 || nums[minOne] != target {
		return []int{-1, -1}
	}

	return []int{minOne, maxOne}

}

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		//mid := left + (right-left)/2
		mid := (right + left) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = left + 1
		}
	}

	return left
}

// LeftBound 左边界找不到会返回这个数字应该出现的索引位置
func LeftBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// RightBound 这里如果是len - 1 ， 下面吗for 不加=号，意味着循环的终止条件是 left = right， 也就是是left = right不会进入到运算中，因为这个不符合循环条件
// 但是实际上 left = right = len（nums） -1 这种情况需要扫描，但是如果是for <= 的情况，数组里没有target时候会一直死循环
func RightBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	lIndex := BinarySearchLeft(nums, target)
	rIndex := BinarySearchRight(nums, target)

	if lIndex == len(nums) || nums[lIndex] != target {
		return 0
	}

	fmt.Println(lIndex)
	fmt.Println(rIndex)

	return rIndex - lIndex + 1
}

func BinarySearchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right + left) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func BinarySearchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right + left) >> 1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}
