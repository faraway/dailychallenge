package utils

import "fmt"

//package init function is optional, no parameter, no return value
func init() {
	fmt.Println("utils package initialized")
}

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Value int
	Next *ListNode
}

//Not that I don't know generics..Just for backwards compatibility
//(so that I don't have to refactor all existing code)
type TreeNodeStr struct {
	Value string
	Left *TreeNodeStr
	Right *TreeNodeStr
}

//Print the tree with in "pre order"
func (node *TreeNodeStr) PrintPreorder() {
	if node == nil { return }
	fmt.Print(node.Value," ")
	node.Left.PrintPreorder()
	node.Right.PrintPreorder()
}

func TestFunc() {
	fmt.Println("calling package function...")
}
