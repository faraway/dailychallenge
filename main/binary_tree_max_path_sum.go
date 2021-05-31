package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
https://leetcode.com/problems/binary-tree-maximum-path-sum/  [HARD]
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them.
A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any path.

Example 1:
      1
     / \
    2   3
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

Example 2:
          -10
         /   \
        9     20
		 	 /  \
            15   7
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

Constraints:
The number of nodes in the tree is in the range [1, 3 * 10^4].
-1000 <= Node.val <= 1000

 */

func main() {
	fmt.Println("initializing test data...")

	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	fmt.Println("Answer is:", maxPathSum(root)) //6

	root2 := &TreeNode{Value: -10}
	root2.Left = &TreeNode{Value: 9}
	root2.Right = &TreeNode{Value: 20}
	root2.Right.Left = &TreeNode{Value: 15}
	root2.Right.Right = &TreeNode{Value: 7}
	fmt.Println("Answer is:", maxPathSum(root2)) //42

	root3 := &TreeNode{Value: -3}
	fmt.Println("Answer is:", maxPathSum(root3)) //-3
}

func maxPathSum(root *TreeNode) int {
	pathMax, subtreeMax := maxPathSumInternal(root)
	if pathMax > subtreeMax {
		return pathMax
	} else {
		return subtreeMax
	}
}

/**
	returns @pathMax - the max value including this root node (such that it can be used by its parent node to form a even longer path)
	and,    @subtreeMax - the max value (not including the root) we can have with this subtree. But this cannot be used by its parent as it's disconnected
 */
func maxPathSumInternal(root *TreeNode) (pathMax int, subtreeMax int) {
	if root == nil {
		return 0, 0
	}
	leftPathMax, leftSubtreeMax := maxPathSumInternal(root.Left)
	rightPathMax, rightSubtreeMax := maxPathSumInternal(root.Right)

	//calculate subtreeMax
	//it's either leftSubtreeMax | rightSubtreeMax | leftPathMax | rightPathMax | leftPathMax + root.node + rightPathMax
	subtreeMax = root.Value
	if root.Left != nil {
		if leftSubtreeMax > subtreeMax {
			subtreeMax = leftSubtreeMax
		}
		if leftPathMax > subtreeMax {
			subtreeMax = leftPathMax
		}
	}
	if root.Right != nil {
		if rightSubtreeMax > subtreeMax {
			subtreeMax = rightSubtreeMax
		}
		if rightPathMax > subtreeMax {
			subtreeMax = rightPathMax
		}
	}
    //taking both left and right path, going through root. This makes the result subtree max (but not path max) by definition
	if leftPathMax + root.Value + rightPathMax > subtreeMax {
		subtreeMax = leftPathMax + root.Value + rightPathMax
	}

	//calculate pathMax
	//it's either left path + root node OR right path + root node
	leftGain, rightGain := 0, 0
	if leftPathMax > 0 {
		leftGain = leftPathMax
	}
	if rightPathMax > 0 {
		rightGain = rightPathMax
	}
	if leftGain > rightGain {
		pathMax = leftGain + root.Value
	} else {
		pathMax = rightGain + root.Value
	}

	return pathMax, subtreeMax
}