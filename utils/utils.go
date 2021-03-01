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
	Val int
	Next *ListNode
}

func TestFunc() {
	fmt.Println("calling package function...")
}
