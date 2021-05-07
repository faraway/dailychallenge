package main

import (
	"fmt"
)

/**

Coding interview question when I was interviewing with c3.ai
(similar to https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/, but simpler)

               1 (0)
       2 (-1)         3 (1)
  4 (-2)      9(0) 6 (0)   7 (2)

print out the values of the tree nodes in a column-level order
group tree nodes into columns
assign 0 as the coordinate to the root
cooridnate of the left child = cooridnate of the parent -  1
cooridnate of the right child = cooridnate of the parent +  1

group nodes with the same coordinate in the same column

print out nodes in the leftmost column first, within the same column,
print out values on the top first,
then print out nodes in the second leftmost column, ... until the rightmost column

[4, 2, 1, 6, 9, 3, 7] or [4, 2, 1, 9, 6, 3, 7]

*/
func main() {
	fmt.Println("initializing test data...")

	root := &CNode{val:1}
	root.left = &CNode{val:2}
	root.left.left = &CNode{val:4}
	root.left.right = &CNode{val:9}
	root.right = &CNode{val:3}
	root.right.left = &CNode{val:6}
	root.right.right = &CNode{val:7}

	fmt.Println("Answer is:", verticalTraversal(root)) //[4, 2, 1, 9, 6, 3, 7]
}

type CNode struct {
	left *CNode
	right *CNode
	val int
	coordinator int
}

func verticalTraversal(root *CNode) []int {
	var result []int

	columns := make(map[int]*[]int)

	minCoordinator := 0
	bfs := []*CNode{root}
	for len(bfs) > 0 {
		head := bfs[0]
		if head.coordinator < minCoordinator {
			minCoordinator = head.coordinator
		}
		//each column is under the same "coordinator" number
		column := columns[head.coordinator]
		if column == nil {
			column = &[]int{}
			columns[head.coordinator] = column
		}
		//since we use BFS, it's guaranteed top-down ordered within each column
		*column = append(*column, head.val)
		if head.left != nil {
			head.left.coordinator = head.coordinator - 1
			bfs = append(bfs, head.left)
		}
		if head.right != nil {
			head.right.coordinator = head.coordinator + 1
			bfs = append(bfs, head.right)
		}
		bfs = bfs[1:] //remove the head
	}

	//we can certainly sort the keys of the hashmap...but below is a smarter way using the 'minCoordinator' we get above
	for columns[minCoordinator] != nil {
		for _, num := range *columns[minCoordinator] {
			result = append(result, num)
		}
		minCoordinator++
	}

	return result
}