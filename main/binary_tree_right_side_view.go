package main

import (
	. "dailychallenge/utils"
	"fmt"
)
/**
https://leetcode.com/problems/binary-tree-right-side-view/

Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.



Example 1:
     1
   /   \
  2     3
   \     \
    5     4
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

Example 3:
Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
 */
func main() {
	fmt.Println("initializing test data...")

	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}

	root.Left.Right = &TreeNode{Value: 5}
	root.Right.Right = &TreeNode{Value: 4}
	fmt.Println("Answer is:", rightSideView(root)) //[1,3,4]
	fmt.Println("Answer is:", rightSideView_BFS_with_two_queues(root)) //[1,3,4]

	var root2 *TreeNode = nil
	fmt.Println("Answer is:", rightSideView(root2)) //[1,3,4]
	fmt.Println("Answer is:", rightSideView_BFS_with_two_queues(root2)) //[1,3,4]
}

/**
	BFS with "level" indicator
 */
func rightSideView(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	nodes := []NodeTuple{{root, 1}}
	for len(nodes) > 0 {
		head := nodes[0]
		if head.node.Left != nil {
			nodes = append(nodes, NodeTuple{head.node.Left, head.level+1})
		}
		if head.node.Right != nil {
			nodes = append(nodes, NodeTuple{head.node.Right, head.level+1})
		}
		if len(nodes) == 1 || nodes[1].level > head.level { //head is the last node of this level
			result = append(result, head.node.Value)
		}
		nodes = nodes[1:]
	}
	return result
}

type NodeTuple struct{
	node *TreeNode
	level int
}

/**
  	BFS with two queues
 */
func rightSideView_BFS_with_two_queues(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	current := []*TreeNode{root}
	next := []*TreeNode{}
	for len(current) > 0 {
		head := current[0]
		if head.Left != nil {
			next = append(next, head.Left)
		}
		if head.Right != nil {
			next = append(next, head.Right)
		}

		current = current[1:]
		if len(current) == 0 { //we are at the end of this level
			//collect result
			result = append(result, head.Value)
			//swap queues
			current, next = next, current
		}
	}
	return result
}

/**
	DFS possible too!!
  	https://leetcode.com/problems/binary-tree-right-side-view/solution/
 */
func rightSideView_DFS(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	//TODO
	return result
}