package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTreeNode(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTreeNode(t1.Left, t2.Left)
	t1.Right = mergeTreeNode(t1.Right, t2.Right)
	return t1
}

func InTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	InTraversal(node.Left)
	fmt.Printf("%d ", node.Val)
	InTraversal(node.Right)
}

func main() {
	t1 := &TreeNode{Val: 1}
	t1.Left = &TreeNode{Val: 3}
	t1.Right = &TreeNode{Val: 2}
	t1.Left.Left = &TreeNode{Val: 5}

	t2 := &TreeNode{Val: 2}
	t2.Left = &TreeNode{Val: 1}
	t2.Right = &TreeNode{Val: 3}
	t2.Left.Right = &TreeNode{Val: 4}
	t2.Right.Right = &TreeNode{Val: 7}

	treeNode := mergeTreeNode(t1, t2)
	InTraversal(treeNode)
}
