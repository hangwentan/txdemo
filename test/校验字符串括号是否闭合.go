package main

import (
	"container/list"
	"fmt"
)

func is_valid_v1(str string) bool {
	strLen := len(str)
	//判断数据合法性，基本的校验
	if strLen == 0 {
		return true
	}
	if strLen%2 == 1 {
		return false
	}
	//判断标识位从根据右边获取左边的
	pairsMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	//判断标识位从根据左边边获取右边的
	//pairsMap= map[byte]byte{'(':')','[':']','{':'}'}
	for i := 0; i < strLen/2; i++ {
		indexValue := str[i]
		lastValue := str[strLen-1-i]
		//因map是右边获取左边，则用lastValue 对比indexValue
		if pairsMap[lastValue] != indexValue {
			//如果上述用pairsMap= map[byte]byte{'(':')','[':']','{':'}'}  这判断用indexValue获取对比lastValue
			//if pairsMap[indexValue]!=lastValue{
			return false
		}
	}
	return true
}

func is_valid(str string) bool {
	strLen := len(str)
	//判断数据合法性，基本的校验
	if strLen == 0 {
		return true
	}
	if strLen%2 == 1 {
		return false
	}
	//利用go里面的list 模拟stack 主要用到front 、remove 方法来模拟堆栈的POP、Push方法
	stack := list.New()
	pairsMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < strLen; i++ {
		value := str[i]
		//碰见左边的入栈，右边则不入栈，并获取上一个元素与该元素匹配是否正确
		if pairsMap[value] == 0 {
			//入栈
			stack.PushFront(value)
		} else {
			//出栈、通过右边获取map中的值，与list front相对比
			if stack.Len() == 0 || stack.Front().Value != pairsMap[value] {
				return false
			}
			//fmt.Printf("front value:%v; value:%v \n",stack.Front().Value,pairsMap[value])
			//匹配成功则删除、左右删除，知道最后
			stack.Remove(stack.Front())
		}
	}
	//最后判断list的len大小
	return stack.Len() == 0
}

func main() {
	fmt.Println("string:", is_valid("([]){}"))
	fmt.Println("string:", is_valid("()[]{}"))
	fmt.Println("string:", is_valid("({[]})"))
	fmt.Println("string:", is_valid("){[]}("))
	fmt.Println("string:", is_valid("){[]})"))

	fmt.Println("string:", is_valid_v1("([]){}"))
}
