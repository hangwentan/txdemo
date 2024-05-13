package main

import "fmt"

func main() {

	s := "ababababcdcdcdcdcdc"

	if len(s) <= 1 {
		fmt.Println("字符串长度小于1")
		return
	}

	pd := make([][]bool, len(s))

	for i := range pd {
		pd[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		pd[i][i] = true
	}

	var start, lenmax = 0, 1

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {

			if s[j] == s[i] && (j-i <= 2 || pd[i+1][j-1]) {
				pd[i][j] = true
				if j-i+1 > lenmax {
					start = i
					lenmax = j - i + 1
				}
			}
		}
	}
	fmt.Println("获取最长回文子串")
	fmt.Printf("方法1 动态规划方法 结果： %+v\n", s[start:start+lenmax])
	fmt.Printf("方法2 中心扩展法 结果： %+v\n", longsCenter(s))
	fmt.Println()
	fmt.Printf("获取所有回文串 结果： %+v\n", longsCenterStrings(s))
	fmt.Println()
	fmt.Printf("检测是否回文串 结果： %+v\n", checkLogsCenter("aba"))
}

// 中心扩展法
func longsCenter(s string) string {

	n := len(s)
	// 判断长度是否为空
	if n < 1 {
		return s
	}
	// 声明起始位和结束为
	var start, end int

	for i := 0; i < n; i++ {

		len1 := center(s, i, i)
		len2 := center(s, i, i+1)
		maxLen := len1
		if len1 < len2 {
			maxLen = len2
		}

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func center(string2 string, left int, right int) int {

	for left >= 0 && right < len(string2) && string2[left] == string2[right] {
		left--
		right++
	}
	return right - left - 1
}

// 获取所有回文子串

func longsCenterStrings(s string) []string {

	n := len(s)
	if n <= 1 {
		return []string{s}
	}

	var result []string

	for i := 0; i < n; i++ {

		result = centerString(s, i, i, result)
		result = centerString(s, i, i+1, result)

	}
	return result
}

func centerString(s string, left, right int, result []string) []string {
	n := len(s)
	for left >= 0 && right < n && s[left] == s[right] {
		result = append(result, s[left:right+1])
		left--
		right++
	}
	return result
}

func checkLogsCenter(s string) bool {

	n := len(s)
	for i := 0; i < n/2; i++ {

		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
