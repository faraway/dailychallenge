package main

import "fmt"
/**
Given a list of numbers with only 3 unique numbers (1, 2, 3), sort the list in O(n) time.

Example 1:
Input: [3, 3, 2, 1, 3, 2, 1]
Output: [1, 1, 2, 2, 3, 3, 3]
def sortNums(nums):
  # Fill this in.

print sortNums([3, 3, 2, 1, 3, 2, 1])
# [1, 1, 2, 2, 3, 3, 3]

Challenge: Try sorting the list using constant space.
 */

func main() {
	fmt.Println("initializing test data...")
	fmt.Println("Answer is:", sortThreeNumbers([]int{3,3,2,1,3,2,1})) //1,1,2,2,3,3,3
	fmt.Println("Answer is:", sortThreeNumbers([]int{1,1,2,2,2,2,3,3,3})) //1,1,2,2,2,2,3,3,3
	fmt.Println("Answer is:", sortThreeNumbers([]int{2,3,3,2,3,2,3})) //2,2,2,3,3,3,3
	fmt.Println("Answer is:", sortThreeNumbers([]int{2,1,1,1,2,1,2})) //1,1,1,1,2,2,2
	fmt.Println("Answer is:", sortThreeNumbers([]int{1,3,3,1,3,1,1})) //1,1,1,1,3,3,3
}

func sortThreeNumbers(input []int) []int {
	pointerOne := 0
	pointerTwo := 0
	pointerThree := len(input) - 1

	//initialize these pointers
	for input[pointerOne] == 1 {
		pointerOne ++
	}
	pointerTwo = pointerOne
	for input[pointerThree] == 3 {
		pointerThree --
	}

	for pointerTwo < pointerThree {
		if input[pointerTwo] == 2 {
			pointerTwo ++
		} else if input[pointerTwo] == 3 {
			//swap values of pointer2 and pointer3
			temp := input[pointerThree]
			input[pointerThree] = input[pointerTwo]
			input[pointerTwo] = temp
			pointerThree --
		} else if input[pointerTwo] == 1 {
			//swap values of pointer2 and pinter1
			temp := input[pointerOne]
			input[pointerOne] = 1
			input[pointerTwo] = temp
			pointerOne ++
			pointerTwo ++
		} else {
			panic("This should not happen")
		}
	}

	return input
}
