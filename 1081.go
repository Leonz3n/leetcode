package main

import (
	"fmt"
)

func smallestSubsequence(s string) string {
	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i, _ := range s {
		if inStack[s[i]-'a'] == false {
			for length := len(stack); length > 0 && s[i] < stack[length-1]; length-- {
				if count[stack[length-1]-'a'] == 0 {
					break
				}
				inStack[stack[length-1]-'a'] = false
				stack = stack[:length-1]
			}
			stack = append(stack, s[i])
			inStack[s[i]-'a'] = true
		}
		count[s[i]-'a']--
	}

	return string(stack)
}

func main() {
	fmt.Printf("%v\n", smallestSubsequence("cbacdcbc"))
}
