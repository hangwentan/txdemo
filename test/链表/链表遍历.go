package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func ListRecursive(list *listNode) {
	if list == nil {
		fmt.Println("nil")
		return
	}
	fmt.Printf("%d =>", list.Val)
	ListRecursive(list.Next)
}

func ListInteractive(list *listNode) {
	head := list
	for head != nil {
		fmt.Printf("%d ->", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	list := &listNode{Val: 1}
	list.Next = &listNode{Val: 2}
	list.Next.Next = &listNode{Val: 3}
	list.Next.Next.Next = &listNode{Val: 4}
	list.Next.Next.Next.Next = &listNode{Val: 5}

	// 循环遍历
	ListInteractive(list)

	// 递归遍历
	ListRecursive(list)
}
