package Solution

import "sort"

func nextPermutation(nums []int) {
	i := len(nums) - 2

	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	i, j := start, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 可以但是超出时间限制
func nextPermutation2(nums []int) {
	l := len(nums)
	arr := make([][]int, 0, len(nums))
	var dfs func([]int, map[int]struct{})
	dfs = func(curArr []int, indexMap map[int]struct{}) {
		if len(curArr) == l {
			tmp := make([]int, len(curArr))
			copy(tmp, curArr)
			arr = append(arr, tmp)
			return
		}
		leveMap := make(map[int]struct{})
		for i := 0; i < len(nums); i++ {
			if _, ok := indexMap[i]; ok {
				continue
			}
			if _, ok := leveMap[nums[i]]; ok {
				continue
			}
			leveMap[nums[i]] = struct{}{}
			indexMap[i] = struct{}{}
			curArr = append(curArr, nums[i])
			dfs(curArr, indexMap)
			curArr = curArr[0 : len(curArr)-1]
			delete(indexMap, i)
		}
	}
	indexMap := make(map[int]struct{})
	dfs([]int{}, indexMap)

	sort.Slice(arr, func(i, j int) bool {
		for k := 0; k < len(arr[i]); k++ {
			if arr[i][k] < arr[j][k] {
				return true
			} else if arr[i][k] > arr[j][k] {
				return false
			} else {
				continue
			}
		}

		return true
	})
	for j := 0; j < len(arr); j++ {
		i := 0
		for ; i < l; i++ {
			if arr[j][i] == nums[i] {
				continue
			} else {
				break
			}
		}
		if i == l {
			if j != len(arr)-1 {
				for i2, i3 := range arr[j+1] {
					nums[i2] = i3
				}
				return
			}
		}
	}

	sort.Ints(nums)
	return
}

func nextPermutation3(nums []int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})
	index := 0
	for ; index < len(tmp); index++ {
		if tmp[index] == nums[index] {
			continue
		} else {
			break
		}
	}
	if index == len(nums) {
		sort.Ints(nums)
		return
	}
	i, j := 0, 1
	for j < len(nums) && nums[i] <= nums[j] {
		i++
		j++
	}
	nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
	return

}
