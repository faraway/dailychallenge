package main

import (
	. "dailychallenge/utils"
	"fmt"
	"strconv"
	"strings"
)

/**
https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

Serialization is the process of converting a data structure or object into a sequence of bits so that
it can be stored in a file or memory buffer, or transmitted across a network connection link to be
reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree.
There is no restriction on how your serialization/deserialization algorithm should work.
You just need to ensure that a binary tree can be serialized to a string and
this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree.
You do not necessarily need to follow this format, so please be creative and come up with different
approaches yourself.

*/

func main() {
	fmt.Println("initializing test data...")
	/**
				1
			   / \
			  2    3
			 / \  / \
			4  5 6   7
					/ \
	               8   9
	*/
	root2 := &TreeNode{Value: 1}
	root2.Left = &TreeNode{Value: 2}
	root2.Right = &TreeNode{Value: 3}

	root2.Left.Left = &TreeNode{Value: 4}
	root2.Left.Right = &TreeNode{Value: 5}

	root2.Right.Left = &TreeNode{Value: 6}
	root2.Right.Right = &TreeNode{Value: 7}

	root2.Right.Right.Left = &TreeNode{Value: 8}
	root2.Right.Right.Right = &TreeNode{Value: 9}

	ser := Constructor()
	deser := Constructor()
	fmt.Println("Answer is:", ser.serialize(root2)) //1,2,4,#,#,5,#,#,3,6,#,#,7,8,#,#,9,#,#
	//this should reveal "1,2,4,#,#,5,#,#,3,6,#,#,7,8,#,#,9,#,#" as well
	fmt.Println("Answer is:", ser.serialize(deser.deserialize(ser.serialize(root2))))
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

/**
	Serializes a tree to a single string.
	DFS - pre-order traversal;
	we can use BFS as well - layer by layer traversal the tree, which seems to be the leetcode approach
 */

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Value)+","+this.serialize(root.Left)+","+this.serialize(root.Right)
}

func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	root,_ := deserializeInternal(nodes, 0)
	return root
}

func deserializeInternal(nodes []string, rootIndex int) (*TreeNode, int) {
	rootNodeStr := nodes[rootIndex]
	if rootNodeStr == "#" {
		return nil, 1
	}
	nodeVal, _ := strconv.Atoi(rootNodeStr)
	root := &TreeNode{Value: nodeVal}
	var leftCount, rightCount int
	root.Left, leftCount = deserializeInternal(nodes, rootIndex+1)
	root.Right, rightCount = deserializeInternal(nodes, rootIndex+1+leftCount)
	return root, leftCount+rightCount+1
}
