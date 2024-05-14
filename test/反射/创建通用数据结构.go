package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 创建通用数据结构
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	// 使用反射处理通用数据结构
	for s, i := range data {
		fmt.Printf("%s:%v\n", s, reflect.TypeOf(i))
	}
}
