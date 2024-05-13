package main

import (
	"fmt"
)

// 35
func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	temp := 0
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			temp++
			a[j][i-j] = temp
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != 0 {
				fmt.Printf("%d ", a[i][j])
			}
		}
		fmt.Println()
	}
}
