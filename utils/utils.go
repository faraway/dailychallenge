package utils

import "fmt"

//package init function is optional, no parameter, no return value
func init() {
	fmt.Println("utils package initialized")
}

//-------------------------------List Node -----------------------------------
type ListNode struct {
	Value int
	Next *ListNode
}

//-------------------------------Tree Node-----------------------------------
type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
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

//-------------------------------Stack-----------------------------------
/* Definition of stack */
type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		lastIndex := len(*s) - 1
		item := (*s)[lastIndex]
		*s = (*s)[:lastIndex]
		return item, true
	}
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		return (*s)[len(*s) - 1], true
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}








//-------------------------------Test Func-----------------------------------
func TestFunc() {
	fmt.Println("calling package function...")
}
