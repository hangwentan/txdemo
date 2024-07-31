package main

import "fmt"

func f() {
	defer catch("f")

	g()
}

func catch(funcname string) {
	if r := recover(); r != nil {
		fmt.Println(funcname, "recover:", r)
	}
}

func g() {
	defer m()

	panic("g panic")
}

func m() {
	defer catch("m")

	panic("m panic")
}

func main() {
	f()
}
