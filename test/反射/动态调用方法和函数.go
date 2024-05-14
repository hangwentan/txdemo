package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct{}

func (m *MyStruct) MyMethod() {
	fmt.Println("Dynamic method call")
}

func main() {
	value := &MyStruct{}

	MethodValue := reflect.ValueOf(value).MethodByName("MyMethod")
	if MethodValue.IsValid() {
		MethodValue.Call(nil)
	}
}
