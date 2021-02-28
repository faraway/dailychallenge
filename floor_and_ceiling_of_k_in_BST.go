package main

import "fmt"

/**
Given an integer k and a binary search tree,
find the floor (less than or equal to) of k,
and the ceiling (larger than or equal to) of k.
If either does not exist, then print them as None.
 */
type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func main() {
	fmt.Println("initializing test data...")
	root := &TreeNode{value: 8}
	root.left = &TreeNode{value: 4}
	root.right = &TreeNode{value: 12}

	root.left.left = &TreeNode{value: 2}
	root.left.right = &TreeNode{value: 6}

	root.right.left = &TreeNode{value: 10}
	root.right.right = &TreeNode{value: 14}

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
		//node value can be any number, "-1" doesn't really represent "NotFound" information.
		//so we use 'floorFound' and 'ceilingFound' to indicate that
		return -1, -1, false, false
	}

	if root.value < k {
		//root can be the 'floor' candidate; both floor and ceiling can be further searched in *right* subtree
		childFloor, childCeiling, childFloorFound, childCeilingFound := findFloorAndCeiling(root.right, k)
		if !childFloorFound {
			childFloor = root.value
		}
		return childFloor, childCeiling, true, childCeilingFound
	} else if root.value > k {
		//root can be the 'ceiling' candidate; both floor and ceiling can be further searched in *left* subtree
		childFloor, childCeiling, childFloorFound, childCeilingFound := findFloorAndCeiling(root.left, k)
		if !childCeilingFound {
			childCeiling = root.value
		}
		return childFloor, childCeiling, childFloorFound, true
	} else { // root.value == k -> floor and ceiling IS k by definition
		return k, k, true, true
	}

}
