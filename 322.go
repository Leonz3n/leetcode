package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			if dp[i-coin] == -1 {
				continue
			}

			dp[i] = min(dp[i], 1+dp[i-coin])
		}

		if dp[i] == amount+1 {
			dp[i] = -1
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%v\n", coinChange([]int{1, 2, 5}, 11))
}
