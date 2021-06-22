package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
We are given a binary tree (with root node root), a target node, and an integer value k.

Return a list of the values of all nodes that have a distance k from the target node.
The answer can be returned in any order.

Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
                 3
                / \
               5   1
             / |   | \
            6  2   0  8
			  / \
			 7   4
Output: [7,4,1]
Explanation:
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.


Note that the inputs "root" and "target" are actually TreeNodes.
The descriptions of the inputs above are just serializations of these objects.

Note:
The given tree is non-empty.
Each node in the tree has unique values 0 <= node.val <= 500.
The target node is a node in the tree.
0 <= k <= 1000.
*/

func main() {
	fmt.Println("initializing test data...")
	/**
	                 3
	                / \
	               5   1
	             / |   | \
	            6  2   0  8
				  / \
				 7   4
	 */
	root := &TreeNode{Value: 3}
	root.Left = &TreeNode{Value: 5}
	root.Right = &TreeNode{Value: 1}
	root.Left.Left = &TreeNode{Value: 6}
	root.Left.Right = &TreeNode{Value: 2}
	root.Right.Left = &TreeNode{Value: 0}
	root.Right.Right = &TreeNode{Value: 8}
	root.Left.Right.Left = &TreeNode{Value: 7}
	root.Left.Right.Right = &TreeNode{Value: 4}
	fmt.Println("Answer is:", distanceK(root, root.Left, 2)) //[7,4,1]


	/**
		                 0
		                /
		               1
		             /  \
		            3    2

	 */
	root2 := &TreeNode{Value: 0}
	root2.Left = &TreeNode{Value: 1}
	root2.Left.Left = &TreeNode{Value: 3}
	root2.Left.Right = &TreeNode{Value: 2}
	fmt.Println("Answer is:", distanceK(root2, root2.Left.Right, 1)) //[1]
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	result := []int{}
	//find the child nodes with distance K from target
	findNodeWithDistanceK(target, k, &result)
	//find the parent and their other child nodes with distance K from target
	findTargetNode(root, target, k, &result)

	return result
}

func findTargetNode(root *TreeNode, target *TreeNode, k int, result *[]int) (bool, int){
	if root == nil {
		return false, -1
	}
	if root.Value == target.Value {
		return true, 0
	}
	foundLeft, dl := findTargetNode(root.Left, target, k, result)
	foundRight, dr := findTargetNode(root.Right, target, k, result)
	if foundLeft {
		if k == dl+1 { //this root node to target node is with distance K
			*result = append(*result, root.Value)
		} else if k-dl-2 >= 0 {
			// choose the right child as starting point to find matching nodes; distance will be k - dl - 2 (left child -> root -> right child)
			findNodeWithDistanceK(root.Right, k-dl-2, result)
		}
		return true, dl+1 //distance to left + 1 (this node)
	} else if foundRight {
		if k == dr+1 {
			*result = append(*result, root.Value)
		} else if k-dr-2 >= 0 {
			findNodeWithDistanceK(root.Left, k-dr-2, result)
		}
		return true, dr+1
	} else {
		return false, -1
	}
}

func findNodeWithDistanceK(startingNode *TreeNode, k int, result *[]int) {
	if startingNode == nil {
		return
	}
	if k == 0 {
		*result = append(*result, startingNode.Value)
	} else {
		findNodeWithDistanceK(startingNode.Left, k-1, result)
		findNodeWithDistanceK(startingNode.Right, k-1, result)
	}
}