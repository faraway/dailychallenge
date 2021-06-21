package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
https://leetcode.com/problems/diameter-of-binary-tree/ [EASY]

Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.



Example 1:
Input: root = [1,2,3,4,5]
         1
        /  \
       2    3
      / \
     4   5
Output: 3
Explanation: 3is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:
Input: root = [1,2]
Output: 1


Constraints:
The number of nodes in the tree is in the range [1, 104].
-100 <= Node.val <= 100

*/

func main() {
	fmt.Println("initializing test data...")
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}

	fmt.Println("Answer is:", diameterOfBinaryTree(root)) //3

	root2 := &TreeNode{Value: 1}
	root2.Left = &TreeNode{Value: 2}
	fmt.Println("Answer is:", diameterOfBinaryTree(root2)) //1
}


func diameterOfBinaryTree(root *TreeNode) int {
	pathMax, totalMax := diameterOfBinaryTreeRecursive(root)
	if pathMax-1 > totalMax { //need to do "-1" because "pathMax" assumes "root" has an edge to a higher level node
		return pathMax-1
	} else {
		return totalMax
	}
}

func diameterOfBinaryTreeRecursive(root *TreeNode) (pathMax int, totalMax int) {
	if root == nil {
		return 0, 0
	}
	leftPathMax, leftTotalMax := diameterOfBinaryTreeRecursive(root.Left)
	rightPathMax, rightTotalMax := diameterOfBinaryTreeRecursive(root.Right)

	if leftPathMax > rightPathMax {
		pathMax = leftPathMax
	} else {
		pathMax = rightPathMax
	}

	if leftTotalMax > rightTotalMax {
		totalMax = leftTotalMax
	} else {
		totalMax = rightTotalMax
	}
	if leftPathMax + rightPathMax > totalMax {
		totalMax = leftPathMax + rightPathMax
	}
	//increment pathMax by 1 because of this node
	return pathMax+1, totalMax
}