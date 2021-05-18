package main

import (
	"fmt"
)

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := make([]int, k)
	for k1 := 0; k1 <= k && k1 <= len(nums1); k1++ {
		k2 := k - k1
		if k2 > len(nums2) {
			continue
		}

		merged := merge(pickMaxNumber(nums1, k1), pickMaxNumber(nums2, k2))
		if lexicographicalLess(ans, merged) {
			ans = merged
		}
	}

	return ans
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func merge(nums1, nums2 []int) []int {
	ans := make([]int, len(nums1)+len(nums2))
	for i, _ := range ans {
		if lexicographicalLess(nums1, nums2) {
			ans[i], nums2 = nums2[0], nums2[1:]
		} else {
			ans[i], nums1 = nums1[0], nums1[1:]
		}
	}

	return ans
}

func pickMaxNumber(nums []int, k int) []int {
	var ans []int
	del := len(nums) - k
	for i, _ := range nums {
		for length := len(ans); del > 0 && length > 0 && ans[length-1] < nums[i]; del-- {
			length--
			ans = ans[:length]
		}
		ans = append(ans, nums[i])
	}

	return ans[:k]
}

func main() {
	nums1 := []int{2, 5, 6, 4, 4, 0}
	nums2 := []int{7, 3, 8, 0, 6, 5, 7, 6, 2}

	picked1 := pickMaxNumber(nums1, 6)
	fmt.Printf("nums1: %v\n", picked1)

	picked2 := pickMaxNumber(nums2, 9)
	fmt.Printf("nums2: %v\n", picked2)

	merged := merge(nums1, nums2)
	fmt.Printf("merged: %v\n", merged)

	fmt.Printf("%v\n", maxNumber(nums1, nums2, 15))
}
