package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	var ans string
	remain := len(num) - k
	for _, v := range num {
		for length := len(ans); k > 0 && length > 0 && ans[length-1] > uint8(v); k-- {
			length--
			ans = ans[:length]
		}
		ans = ans + string(v)
	}

	return strings.TrimLeft(ans[:remain], "0")
}

func main() {

	fmt.Printf("%v\n", removeKdigits("10200", 1))
}
