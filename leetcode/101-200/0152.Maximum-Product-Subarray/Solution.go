package Solution

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxF := make([]int, len(nums))
	minF := make([]int, len(nums))
	copy(maxF, nums)
	copy(minF, nums)

	for i := 1; i < len(nums); i++ {
		maxF[i] = max(maxF[i-1]*nums[i], max(nums[i], minF[i-1]*nums[i]))
		minF[i] = min(minF[i-1]*nums[i], min(nums[i], maxF[i-1]*nums[i]))
	}

	result := maxF[0]
	for _, v := range maxF {
		if v > result {
			result = v
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums)) // Output: 6
}

// func maxProduct(nums []int) int {
//     maxE := make([]int, len(nums))
//     minE := make([]int, len(nums))
//     copy(maxE, nums)
//     copy(minE, nums)
//     for i := 1 ; i < len(nums); i++{
//         maxE[i] = max(maxE[i-1]*nums[i], max(nums[i], minE[i-1]*nums[i]))
//         minE[i] = min(minE[i-1]*nums[i], min(nums[i], maxE[i-1]*nums[i]))
//     }

//     res := maxE[0]
//     for _, val := range maxE{
//         res = max(res, val)
//     }

//     return res

// }

// func min(x, y int)int{
//     if x < y {
//         return x
//     }

//     return y
// }

func max2(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func maxProduct2(nums []int) int {
	var product float64
	product = 1
	var res float64
	res = float64(nums[0])

	for _, val := range nums {
		product = product * float64(val)
		res = max2(res, product)
		if val == 0 {
			product = 1
		}
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		product = product * float64((nums[i]))
		res = max2(res, product)
		if nums[i] == 0 {
			product = 1
		}
	}

	return int(res)
}
