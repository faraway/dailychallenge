package main

import "fmt"

/**
Given a sorted array, A, with possibly duplicated elements,
find the indices of the first and last occurrences of a target element, x. Return -1 if the target is not found.

Example:
Input: A = [1,3,3,5,7,8,9,9,9,15], target = 9
Output: [6,8]

Input: A = [100, 150, 150, 153], target = 150
Output: [1,2]

Input: A = [1,2,3,4,5,6,10], target = 9
Output: [-1, -1]
*/

func main() {
	fmt.Println("initializing test data...")
	//test case 1
	fmt.Println(getRange([]int{1,3,3,5,7,8,9,9,9,15}, 9)) //[6,8]
	fmt.Println(getRange([]int{100, 150, 150, 153}, 150)) //[1,2]
	fmt.Println(getRange([]int{1,2,3,4,5,6,10}, 9)) //[-1,-1]
}


func getRange(array []int, target int) (first int, last int) {
	return findRange(array, target, 0, len(array) -1)
}

func findRange(array []int, target int, targetStart int, targetEnd int) (first int, last int) {
	if targetStart > targetEnd {
		return -1, -1
	}
	if targetStart == targetEnd {
		if array[targetEnd] == target {
			return targetStart, targetEnd
		} else {
			return -1, -1
		}
	}
	mid := (targetStart + targetEnd) / 2
	if array[mid]<target {
		return findRange(array, target, mid+1, targetEnd)
	} else if array[mid]>target {
		return findRange(array, target, targetStart, mid-1)
	} else { // array[mid] == target
		leftStart, _ := findRange(array, target, targetStart, mid-1)
		_, rightEnd := findRange(array, target, mid+1, targetEnd)
		if leftStart == -1 {leftStart = mid}
		if rightEnd == -1 {rightEnd = mid}
		return leftStart, rightEnd
	}
}
