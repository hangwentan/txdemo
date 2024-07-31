package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 1.直接使用运算符
func BenchmarkAddStringWithOperator() {
	hello := "hello"
	world := "world"
	hello += "," + world
	fmt.Println(hello)
}

// 2.fmt.Sprintf()
func BenchmarkAddStringWithSprintf() {
	hello := "hello"
	world := "world"
	hello = fmt.Sprintf("%s,%s", hello, world)
	fmt.Println(hello)
}

// 3.strings.Join()
func BenchmarkAddStringWithJoin() {
	hello := "hello"
	world := "world"
	hello = strings.Join([]string{hello, world}, ",")
	fmt.Println(hello)
}

// 4.buffer.WriteString()
func BenchmarkAddStringWithBuffer() {
	hello := "hello"
	world := "world"
	var buffer bytes.Buffer
	buffer.WriteString(hello)
	buffer.WriteString(",")
	buffer.WriteString(world)
	hello = buffer.String()
	fmt.Println(hello)
}

func main() {
	BenchmarkAddStringWithOperator()
	BenchmarkAddStringWithSprintf()
	BenchmarkAddStringWithJoin()
	BenchmarkAddStringWithBuffer()
}
