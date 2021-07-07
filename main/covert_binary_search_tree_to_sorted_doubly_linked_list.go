package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/

Convert a Binary Search Tree to a sorted Circular Doubly-Linked List in place.

You can think of the left and right pointers as synonymous to the predecessor and successor pointers in a doubly-linked list.
For a circular doubly linked list, the predecessor of the first element is the last element, and the successor of the last element is the first element.

We want to do the transformation in place. After the transformation,
the left pointer of the tree node should point to its predecessor, and the right pointer should point to its successor.
You should return the pointer to the smallest element of the linked list.

Example 1:
       4
     /   \
    2     5
   / \
  1   3

Output: [1,2,3,4,5]

       (head)
    ---> 1 --> 2 --> 3 --> 4 --> 5 -----
    | |- 1 <-- 2 <-- 3 <-- 4 <-- 5 <-| |
    | |                              | |
    | |------------------------------| |
    |----------------------------------|

Explanation: The figure below shows the transformed BST. The solid line indicates the successor relationship,
while the dashed line means the predecessor relationship.

*/

func main() {
	fmt.Println("initializing test data...")
	root := &TreeNode{Value: 4}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 5}
	root.Left.Left = &TreeNode{Value: 1}
	root.Left.Right = &TreeNode{Value: 3}
	fmt.Println("Answer is:", treeToDoublyList(root).Value) //1  need to better print this doubly linked list
}


func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	minNode, maxNode := treeToDoublyListRecursive(root)
	minNode.Left = maxNode
	maxNode.Right = minNode
	return minNode
}

func treeToDoublyListRecursive(root *TreeNode) (min *TreeNode, max *TreeNode) {
	if root == nil {
		return nil, nil
	}
	leftmin, leftmax := treeToDoublyListRecursive(root.Left)
	rightmin, rightmax := treeToDoublyListRecursive(root.Right)

	//update min/max node to return
	min, max = root, root
	if leftmin != nil {
		min = leftmin
	}
	if rightmax != nil {
		max = rightmax
	}

	//update links
	if leftmax != nil {
		root.Left = leftmax
		leftmax.Right = root
	}
	if rightmin != nil {
		root.Right = rightmin
		rightmin.Left = root
	}

	return min, max
}
