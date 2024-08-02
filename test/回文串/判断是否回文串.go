package main

import (
	"fmt"
	"strings"
)

func CheckString(s string) bool {

	var str string
	for i := 0; i < len(s); i++ {
		if ('A' <= s[i] && 'Z' >= s[i]) || ('a' <= s[i] && 'z' >= s[i]) || ('0' <= s[i] && '9' >= s[i]) {
			str += string(s[i])
		}
	}
	lenght := len(str)
	str = strings.ToLower(str)
	for i := 0; i < lenght/2; i++ {
		l, r := i, lenght-i-1
		if str[l] == str[r] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckString("A man, a plan, a canal: Panama"))
}
