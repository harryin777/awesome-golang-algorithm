package Solution

func Solution(flowerbed []int, n int) bool {
	l, count := len(flowerbed), 0
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == l-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[i+1] == 0 && flowerbed[i] != 1 {
				n--
				flowerbed[i] = 1
			}
			continue
		}
		if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] != 1 {
				n--
				flowerbed[i-1] = 1
			}
			continue
		}
		if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 && flowerbed[i] != 1 {
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}
