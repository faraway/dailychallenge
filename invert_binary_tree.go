package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
You are given the root of a binary tree. Invert the binary tree in place.
That is, all left children should become right children, and all right children should become left children.

Example:

    a
   / \
  b   c
 / \  /
d   e f

The inverted version of this tree is as follows:

  a
 /  \
c    b
 \  / \
  f e  d

*/

func main() {
	fmt.Println("initializing test data...")

	root := &TreeNodeStr{Value: "a"}
	root.Left = &TreeNodeStr{Value: "b"}
	root.Right = &TreeNodeStr{Value: "c"}

	root.Left.Left = &TreeNodeStr{Value: "d"}
	root.Left.Right = &TreeNodeStr{Value: "e"}

	root.Right.Left = &TreeNodeStr{Value: "f"}
	fmt.Println("Before invert:")
	root.PrintPreorder() //a b d e c f
	fmt.Println("\nAfter invert:")
	invertBinaryTree(root)
	root.PrintPreorder() //expect: a c f b e d
}

func invertBinaryTree(root *TreeNodeStr) {
 	if root == nil { return }
 	//invert left subtree
 	invertBinaryTree(root.Left)
 	//invert right subtree
 	invertBinaryTree(root.Right)
	//swap left and right
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}
