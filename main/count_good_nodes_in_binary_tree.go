package main

import (
	. "dailychallenge/utils"
	"fmt"
	"math"
)

/**
https://leetcode.com/problems/count-good-nodes-in-binary-tree/

Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
Return the number of good nodes in the binary tree.



Example 1:
		3
       / \
      1    4
     /    / \
    3    1   5
Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.

Example 2:
Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.

Example 3:
Input: root = [1]
Output: 1
Explanation: Root is considered as good.


Constraints:
The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].
*/

func main() {
	fmt.Println("initializing test data...")

	root := &TreeNode{Value: 3}
	root.Left = &TreeNode{Value: 1}
	root.Right = &TreeNode{Value: 4}

	root.Left.Left = &TreeNode{Value: 3}

	root.Right.Left = &TreeNode{Value: 1}
	root.Right.Right = &TreeNode{Value: 5}
	fmt.Println("Answer is:", goodNodes(root)) //4

}

var GOOD_NODE_COUNT int = 0

func goodNodes(root *TreeNode) int {
	GOOD_NODE_COUNT = 0
	findGoodNodes(root, math.MinInt32)
	return GOOD_NODE_COUNT
}

func findGoodNodes(root *TreeNode, pathMax int) {
	if root == nil {
		return
	}

	if root.Value >= pathMax {
		GOOD_NODE_COUNT++
		pathMax = root.Value
	}
	findGoodNodes(root.Left, pathMax)
	findGoodNodes(root.Right, pathMax)
}
