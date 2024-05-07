package main

import "fmt"

type treeNode struct {
	Val   int
	Left  *treeNode
	right *treeNode
}

func ProOrderTraversal(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	ProOrderTraversal(node.Left)
	ProOrderTraversal(node.right)
}

func InOrderTraversal(node *treeNode) {
	if node == nil {
		return
	}
	ProOrderTraversal(node.Left)
	fmt.Println(node.Val)
	ProOrderTraversal(node.right)
}

func PostOrderTraversal(node *treeNode) {
	if node == nil {
		return
	}
	ProOrderTraversal(node.Left)
	ProOrderTraversal(node.right)
	fmt.Println(node.Val)
}

func main() {
	TreeNode := &treeNode{Val: 1}
	TreeNode.Left = &treeNode{Val: 2}
	TreeNode.right = &treeNode{Val: 3}
	TreeNode.Left.Left = &treeNode{Val: 4}
	TreeNode.Left.right = &treeNode{Val: 5}
	TreeNode.right.Left = &treeNode{Val: 6}
	TreeNode.right.right = &treeNode{Val: 7}

	// 前序遍历
	fmt.Println("前序遍历")
	ProOrderTraversal(TreeNode)

	// 中序排序
	fmt.Println("中序排序")
	InOrderTraversal(TreeNode)

	// 后序排序
	fmt.Println("前序遍历")
	PostOrderTraversal(TreeNode)
}
