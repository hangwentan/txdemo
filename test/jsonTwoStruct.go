package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	List []List `json:"list"`
}

type List struct {
	Day int `json:"day"`
}

func jsonTwoStruct(str string) {
	obj := new(User)
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*obj)
}

func jsonTwoStructSlice(str string) {
	obj := make([]User, 0)
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
}

func jsonTwoMap(str string) {
	obj := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
}

func jsonTwoMapSlice(str string) {
	obj := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
}

func StructTwoJson(sut User) {
	str, ok := json.Marshal(sut)
	if ok != nil {
		fmt.Printf("%s\n", ok)
		return
	}
	fmt.Printf("%s\n", str)
}

func StructSliceTwoJson(sut []User) {
	str, ok := json.Marshal(sut)
	if ok != nil {
		fmt.Printf("%s\n", ok)
		return
	}
	fmt.Printf("%s\n", str)
}

// bytes与string互转
func string2Bytes(s string) []byte {
	var x = (*[2]uintptr)(unsafe.Pointer(&s))
	var h = [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	jsonTwoStruct(`{"id":1,"name":"aaa","list":[{"day":3},{"day":4}]}`)
	jsonTwoStructSlice(`[{"id":1,"name":"aaa","list":[{"day":3},{"day":4}]},{"id":2,"name":"bbb","list":[{"day":5},{"day":6}]}]`)

	jsonTwoMap(`{"id":1,"name":"aaa","list":[{"day":3},{"day":4}]}`)
	jsonTwoMapSlice(`[{"id":1,"name":"aaa","list":[{"day":3},{"day":4}]},{"id":2,"name":"bbb","list":[{"day":5},{"day":6}]}]`)

	var uu User
	uu.List = append(uu.List, List{Day: 1})
	uu.List = append(uu.List, List{Day: 2})
	StructTwoJson(User{Id: 2, Name: "bbb", List: uu.List})

	var mapSlice []User
	var u1 User
	u1.List = append(u1.List, List{Day: 1}, List{Day: 2})
	mapSlice = append(mapSlice, User{Id: 1, Name: "aaa", List: u1.List})
	u1.List = append(u1.List[:1], u1.List[2:len(u1.List)]...)
	u1.List = append(u1.List, List{Day: 3}, List{Day: 4})
	mapSlice = append(mapSlice, User{Id: 2, Name: "bbb", List: u1.List})
	StructSliceTwoJson(mapSlice)
}
