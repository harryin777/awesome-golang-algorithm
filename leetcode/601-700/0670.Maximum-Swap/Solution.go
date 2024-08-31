package Solution

import "strconv"

func maximumSwap(num int) int {
	str := strconv.Itoa(num)
	lastOne := -1
	maxOne := -1
	for i := len(str) - 1; i > 0; i-- {
		if (str[i] == '9' && (i < len(str)-1 && str[i+1] != '9')) || str[i] > str[i-1] && int(str[i]-'0') > maxOne {
			lastOne = i
			maxOne = int(str[i] - '0')
		}
	}

	if lastOne == -1 {
		return num
	}

	preOne := -1
	for i := 0; i < len(str); i++ {
		if str[i] < str[lastOne] && i < lastOne {
			preOne = i
			break
		}
	}

	if preOne == -1 {
		return num
	}

	tmp := []byte(str)
	tmp[preOne], tmp[lastOne] = tmp[lastOne], tmp[preOne]
	ans, _ := strconv.Atoi(string(tmp))
	return ans
}
