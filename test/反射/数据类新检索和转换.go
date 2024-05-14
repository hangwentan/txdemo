package main

import (
	"fmt"
	"reflect"
)

func main() {
	var X interface{} = 7

	if reflect.TypeOf(X).Kind() == reflect.Int {
		fmt.Println("该数据是 int 类型")
	}

	y := reflect.ValueOf(X).Interface().(int)
	fmt.Println("Value of y:", y)
}
