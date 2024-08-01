package main

import "fmt"

func stringReverse() {
	var str = "hello"

	var bytes []byte = []byte(str)
	for i := 0; i < len(bytes); i++ {

		tmp := str[len(bytes)-i-1]

		bytes[len(bytes)-i-1] = str[i]
		bytes[i] = tmp
	}
	fmt.Println(string(bytes))
}

func stringReverse2() {
	var str = "hello"
	var bytes []byte = []byte(str)
	for i := 0; i < len(bytes); i++ {
		bytes[len(str)-i-1], bytes[i] = str[i], str[len(bytes)-i-1]
	}
	fmt.Println(string(bytes))
}

func main() {
	stringReverse()
	stringReverse2()
}
