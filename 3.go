package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[int32]int)
	left := 0
	max := 0
	for i, ch := range s {
		if elem, ok := charMap[ch]; ok && elem >= left {
			left = elem + 1
		}

		charMap[ch] = i

		if i-left+1 > max {
			max = i - left + 1
		}
	}

	return max
}

func main() {
	fmt.Printf("max: %d\n", lengthOfLongestSubstring("abcabcbb"))
}
