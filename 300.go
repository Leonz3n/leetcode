package main

import "sort"

func lengthOfLIS(nums []int) int {
	f, length := make([]int, 0, len(nums)), 1
	if cap(f) == 0 {
		return 0
	}
	f = append(f, nums[0])
	for _, v := range nums {
		if v > f[length-1] {
			f = append(f, v)
			length++
		} else {
			insert := sort.Search(len(f), func(i int) bool {
				return f[i] >= v
			})

			if insert < len(f) {
				f[insert] = v
			}
		}
	}
	return length
}

func main() {

}
