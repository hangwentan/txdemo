package main

import "fmt"

func Do(s string) {

	res := make([]rune, 0)
	a := [26]int{0}

	for _, i2 := range s {
		a[i2-'a']++
	}

	min := len(s) + 1
	for _, i2 := range s {
		if a[i2-'a'] < min {
			min = a[i2-'a']
		}
	}

	for _, i2 := range s {
		if a[i2-'a'] > min {
			res = append(res, i2)
		}
	}
	fmt.Println(string(res))
}
