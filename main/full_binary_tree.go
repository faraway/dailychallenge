package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
Given a binary tree, remove the nodes in which there is only 1 child, so that the binary tree is a full binary tree.

So leaf nodes with no children should be kept, and nodes with 2 children should be kept as well.

# Given this tree:
#     1
#    / \
#   2   3
#  /   / \
# 0   9   4

# We want a tree like:
#     1
#    / \
#   0   3
#      / \
#     9   4

*/
func main() {
	fmt.Println("initializing test data...")
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}

	root.Left.Left = &TreeNode{Value: 0}
	root.Left.Left.Right = &TreeNode{Value: 5}

	root.Right.Left = &TreeNode{Value: 9}
	root.Right.Right = &TreeNode{Value: 4}

	shapeToBinaryTree(root).PrintPreorder() //1 5 3 9 4
	fmt.Println("\n---------")
	root2 := &TreeNode{Value: 1}
	root2.Left = &TreeNode{Value: 2}
	root2.Right = &TreeNode{Value: 3}

	root2.Left.Left = &TreeNode{Value: 0}
	root2.Left.Left.Left = &TreeNode{Value: 7}
	root2.Left.Left.Right = &TreeNode{Value: 5}

	root2.Right.Left = &TreeNode{Value: 9}
	root2.Right.Right = &TreeNode{Value: 4}

	shapeToBinaryTree(root2).PrintPreorder() //1 0 7 5 3 9 4

}

/**
#     1                   1
#    / \                 / \
#   2   3               5   3
#  /   / \                 / \
# 0   9   4    =>         9   4      (remove node 2 and 0)
   \
#   5

#     1                   1
#    / \                 / \
#   2   3               0   3
#  /   / \            / \  / \
# 0   9   4    =>    7   5 9  4      (remove node 2)
 / \
7   5
 */
func shapeToBinaryTree(root *TreeNode) *TreeNode {
	//leaf node
	if root.Right == nil && root.Left == nil {
		return root
	}
	//full node, just need to adjust children nodes
	if root.Right != nil && root.Left != nil {
		root.Left = shapeToBinaryTree(root.Left)
		root.Right = shapeToBinaryTree(root.Right)
		return root
	}
	//otherwise this node is 1-child node, and needs to be removed
	if root.Left == nil {
		return shapeToBinaryTree(root.Right)
	} else{
		return shapeToBinaryTree(root.Left)
	}
}



