package main

import "fmt"

/**
Hacker Rank question - TESTED

In an array, elements at any two indices can be swapped in a single operation called a move.
For example, if the array is arr = [17, 4, 8], swap arr[0] = 17 and arr[2] = 8 to get arr' = [8, 4, 17] in a single move.
Determine the minimum number of moves required to sort an array such that all of the even elements are at the beginning
of the array and all of the odd elements are at the end of the array.

Example
arr = [6, 3, 4, 5]

The following four arrays are valid custom-sorted arrays:
·         a = [6, 4, 3, 5]
·         a = [4, 6, 3, 5]
·         a = [6, 4, 5, 3]
·         a = [4, 6, 5, 3]
The most efficient sorting requires 1 move: swap the 4 and the 3.

Note: The order of the elements within even or odd does not matter.
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", moves([]int32{6,3,4,5})) //1
	fmt.Println("Answer is:", moves([]int32{6,4,1,8,3,5})) //1
	fmt.Println("Answer is:", moves([]int32{9})) //0
	fmt.Println("Answer is:", moves([]int32{6,4,1,3,5})) //0

}

func moves(arr []int32) int32 {

	left, right := 0, len(arr)-1
	var moves int32 = 0

	for left < right {
		for !isOddNumber(arr[left]) {
			left++
		}
		for isOddNumber(arr[right]) {
			right--
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			moves++
		}
	}
	return moves
}

func isOddNumber(num int32) bool {
	return num%2 == 1
}