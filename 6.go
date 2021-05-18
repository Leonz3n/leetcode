package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ret := ""
	cycleLen := numRows*2 - 2
	length := len(s)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j += cycleLen {
			ret += string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < length {
				ret += string(s[j+cycleLen-i])
			}
		}
	}

	ans := make([]string, numRows)
	curRow := 0
	goingDown := false
	for _, ch := range s {
		ans[curRow] = ans[curRow] + string(ch)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	for _, str := range ans {
		ret += str
	}

	return ret
}

func main() {

}
