package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 序列化
	person := Person{Name: "Alice", Age: 25}
	jsonData, _ := json.Marshal(person)
	fmt.Println("Serialized data:", string(jsonData))

	// 反序列化
	var newPerson Person
	json.Unmarshal(jsonData, &newPerson)
	fmt.Println("Deserialized data:", newPerson)

	// 使用反射检查数据类型
	fmt.Println("Type of newPerson:", reflect.TypeOf(newPerson))
}
