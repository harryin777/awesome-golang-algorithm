package Solution

func intToRoman(num int) string {
	i := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	x := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	c := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	m := [4]string{"", "M", "MM", "MMM"}
	return m[num/1000] + c[num%1000/100] + x[num%1000%100/10] + i[num%1000%100%10]
}

func intToRoman2(num int) string {
	var res string
	for num != 0 {
		var str string
		num, str = cal(num)
		res += str
	}

	return res
}

func cal(num int) (int, string) {
	if num >= 1000 {
		return num - 1000, "M"
	}
	if num >= 900 {
		return num - 900, "CM"
	}
	if num > 500 {
		return num - 500, "D"
	}
	if num >= 400 {
		return num - 400, "CD"
	}
	if num >= 100 {
		return num - 100, "C"
	}
	if num >= 90 {
		return num - 90, "XC"
	}
	if num >= 50 {
		return num - 50, "L"
	}
	if num >= 40 {
		return num - 40, "XL"
	}
	if num >= 10 {
		return num - 10, "X"
	}
	if num >= 9 {
		return num - 9, "IX"
	}
	if num >= 5 {
		return num - 5, "V"
	}
	if num >= 4 {
		return num - 4, "IV"
	}
	if num >= 1 {
		return num - 1, "I"
	}

	return 0, ""
}
