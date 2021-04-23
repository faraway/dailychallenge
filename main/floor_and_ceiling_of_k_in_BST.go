package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
Given an integer k and a binary search tree,
find the floor (less than or equal to) of k,
and the ceiling (larger than or equal to) of k.
If either does not exist, then print them as None.
 */

func main() {
	fmt.Println("initializing test data...")
	root := &TreeNode{Value: 8}
	root.Left = &TreeNode{Value: 4}
	root.Right = &TreeNode{Value: 12}

	root.Left.Left = &TreeNode{Value: 2}
	root.Left.Right = &TreeNode{Value: 6}

	root.Right.Left = &TreeNode{Value: 10}
	root.Right.Right = &TreeNode{Value: 14}

	// test binary search tree
	//                 8
	//               /   \
	//              /     \
	//             4       12
	//            / \      / \
	//           /   \    /   \
	//          2     6  10   14

	floor, ceiling, f1, f2 := findFloorAndCeiling(root, 4)  //4 4 true true
	fmt.Println("Answer is:", floor, ceiling, f1, f2)
	floor, ceiling, f1, f2 = findFloorAndCeiling(root, 5)  //4 6 true true
	fmt.Println("Answer is:", floor, ceiling, f1, f2)
	floor, ceiling, f1, f2 = findFloorAndCeiling(root, 9)  //8 10 true true
	fmt.Println("Answer is:", floor, ceiling, f1, f2)
	floor, ceiling, f1, f2 = findFloorAndCeiling(root, 1)  //-1 2 false true
	fmt.Println("Answer is:", floor, ceiling, f1, f2)
	floor, ceiling, f1, f2 = findFloorAndCeiling(root, 99)  //14 -1 true false
	fmt.Println("Answer is:", floor, ceiling, f1, f2)

}

func findFloorAndCeiling(root *TreeNode, k int) (floor int, ceiling int, floorFound bool, ceilingFound bool) {
	if root == nil {
		//node Value can be any number, "-1" doesn't really represent "NotFound" information.
		//so we use 'floorFound' and 'ceilingFound' to indicate that
		return -1, -1, false, false
	}

	if root.Value < k {
		//root can be the 'floor' candidate; both floor and ceiling can be further searched in *Right* subtree
		childFloor, childCeiling, childFloorFound, childCeilingFound := findFloorAndCeiling(root.Right, k)
		if !childFloorFound {
			childFloor = root.Value
		}
		return childFloor, childCeiling, true, childCeilingFound
	} else if root.Value > k {
		//root can be the 'ceiling' candidate; both floor and ceiling can be further searched in *Left* subtree
		childFloor, childCeiling, childFloorFound, childCeilingFound := findFloorAndCeiling(root.Left, k)
		if !childCeilingFound {
			childCeiling = root.Value
		}
		return childFloor, childCeiling, childFloorFound, true
	} else { // root.Value == k -> floor and ceiling IS k by definition
		return k, k, true, true
	}

}
