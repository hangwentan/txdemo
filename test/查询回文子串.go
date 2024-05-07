package main

import "fmt"

func CallBackString(str string) bool {
	return true
}

func CallBackStringNum(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		fmt.Printf("l=%v r=%v\n", l, r)
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println("Num>>", CallBackStringNum("abcdefga"))
}
