package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
https://leetcode.com/problems/path-sum-iii/

Given the root of a binary tree and an integer targetSum,
return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf,
but it must go downwards (i.e., traveling only from parent nodes to child nodes).

                   10
                /      \
               5        -3
             /   \        \
            3     2        11
           / \     \
          3  -2     1

Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are (5,3)  (5,2,1) and (-3,11)
*/

func main() {
	fmt.Println("initializing test data...")
	/**
   	        10
  	      /      \
  	     5        -3
 	    /   \        \
 	   3     2        11
 	  / \     \
 	 3  -2     1

 	*/
	root := &TreeNode{Value: 10}
	root.Left = &TreeNode{Value: 5}
	root.Right = &TreeNode{Value: -3}

	root.Left.Left = &TreeNode{Value: 3}
	root.Left.Right = &TreeNode{Value: 2}
	root.Right.Right = &TreeNode{Value: 11}

	root.Left.Left.Left = &TreeNode{Value: 3}
	root.Left.Left.Right = &TreeNode{Value: -2}
	root.Left.Right.Right = &TreeNode{Value: 1}

	fmt.Println("Answer is:", pathSum(root, 8)) //3

	/**
	       1
	      / \
	     -2 -3
	 */
	root2 := &TreeNode{Value: 1}
	root2.Left = &TreeNode{Value: -2}
	root2.Right = &TreeNode{Value: -3}
	fmt.Println("Answer is:", pathSum(root2, -1)) //1
}

func pathSum(root *TreeNode, targetSum int) int {
	sumMap := make(map[int]int)
	count := 0
	traverseTree(root, 0, targetSum, &count, sumMap)
	return count
}

func traverseTree(root *TreeNode, prefixSum int, targetSum int, count *int, sumMap map[int]int) {
	if root == nil {
		return
	}

	prefixSum += root.Value
	if prefixSum == targetSum { //path from root
		*count++
	}
	occurrence, exists := sumMap[prefixSum-targetSum]
	if exists { //path from somewhere in the middle
		*count += occurrence
	}
	sumMap[prefixSum]++
	traverseTree(root.Left, prefixSum, targetSum, count, sumMap)
	traverseTree(root.Right, prefixSum, targetSum, count, sumMap)
	//maps are passed by reference, revert to the state as if this node was not visited to make sure
	//we have the right map when visit this node's parents and their children
	sumMap[prefixSum]--
}