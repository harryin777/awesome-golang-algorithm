package Solution

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows <= 0 || numRows > len(s) {
		return s
	}
	strArr := make([][]byte, numRows)
	j := 0
	increasing := true

	for i := 0; i < len(s); i++ {
		strArr[j] = append(strArr[j], s[i])

		if increasing {
			j++
			if j == numRows-1 {
				increasing = false
			}
		} else {
			j--
			if j == 0 {
				increasing = true
			}
		}
	}
	fmt.Println(strArr)
	res := make([]byte, 0)

	for _, str := range strArr {
		res = append(res, str...)
	}
	return string(res)
}

func convert2(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	lenth, cycle, n := len(s), 2*numRows-2, 0
	result := make([]byte, lenth)
	for i := 0; i < numRows; i++ {
		for j1 := i; j1 < lenth; j1 = cycle + j1 {

			n += copy(result[n:], string(s[j1]))
			j2 := j1 + cycle - 2*i
			if i != 0 && i != numRows-1 && j2 < lenth {
				n += copy(result[n:], string(s[j2]))
			}

		}
	}
	return string(result)
}

func convert3(s string, numRows int) string {
	buf := bytes.NewBufferString("")
	if numRows == 1 {
		return s
	}

	width := 2 * (numRows - 1)
	for idx := 0; idx < numRows; idx++ {
		step := 2 * idx
		if idx == 0 || idx == numRows-1 {
			step = width
		}
		for walker := idx; walker < len(s); walker += step {
			buf.WriteByte(s[walker])
			if step != width {
				step = width - step
			}
		}
	}

	return buf.String()
}

func convert4(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}
	grid := make([][]string, numRows)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]string, len(s))
	}

	index := 0
	var i, j int
	for index < len(s) {
		for ; i < len(grid); i++ {
			if index == len(s) {
				goto end
			}
			grid[i][j] = string(s[index])
			index++
		}

		j++
		for i = i - 2; i > 0; i-- {
			if index == len(s) {
				goto end
			}
			grid[i][j] = string(s[index])
			j++
			index++
		}
		i = 0
	}

end:
	var ans string
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ans += grid[i][j]
		}
	}

	return ans
}
