package main

import "fmt"

func longestPalindromeDP(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					start = i
				}
			}
		}
	}

	return s[start : start+maxLen]
}

func main() {
	fmt.Printf(longestPalindromeDP("abababcdefabcbcbcbcbbbb"))
}
