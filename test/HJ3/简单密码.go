package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)

	tmp := 'a' - 'A' + 1
	a := []int{2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9}
	for _, i2 := range str {
		if i2 >= 'A' && i2 <= 'Z' {
			s := i2 + tmp
			if s > 'z' {
				s = 'a'
			}
			fmt.Print(string(s))
		} else if i2 >= '0' && i2 <= '9' {
			fmt.Print(string(i2))
		} else {
			fmt.Print(a[i2-'a'])
		}
	}
}
