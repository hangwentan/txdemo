package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if err := valid(input.Text()); err {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}

}

func valid(s string) bool {

	if len(s) <= 8 {
		return false
	}

	counts := [4]int{0}

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			counts[0]++
		} else if v >= 'A' && v <= 'Z' {
			counts[1]++
		} else if v >= '0' && v <= '9' {
			counts[2]++
		} else {
			counts[3]++
		}
	}
	diffType := 0
	for _, v := range counts {
		if v != 0 {
			diffType++
		}
	}
	if diffType < 3 {
		return false
	}
	m := make(map[string]bool)
	for i := 0; i < len(s)-2; i++ {
		if _, ok := m[s[i:i+3]]; !ok {
			m[s[i:i+3]] = true
		} else {
			return false
		}
	}
	return true
}
