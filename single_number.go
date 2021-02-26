package main

import "fmt"
/**
Given a list of numbers, where every number shows up twice except for one number, find that one number.

Example:
Input: [4, 3, 2, 4, 1, 3, 2]
Output: 1

Challenge: Find a way to do this using O(1) memory.
*/
func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", solution([]int{4,3,2,1,3,2,4})) //1
	fmt.Println("Answer is:", solution([]int{7, 3, 5, 4, 5, 3, 4})) //7
}

/**
The best solution is to use XOR.
XOR of all array elements gives us the number with a single occurrence. The idea is based on the following two facts.
a) XOR of a number with itself is 0.
b) XOR of a number with 0 is number itself.

e.g.
res = 7 ^ 3 ^ 5 ^ 4 ^ 5 ^ 3 ^ 4

Since XOR is associative and commutative, above
expression can be written as:
res = 7 ^ (3 ^ 3) ^ (4 ^ 4) ^ (5 ^ 5)
    = 7 ^ 0 ^ 0 ^ 0
    = 7 ^ 0
    = 7
*/
func solution(nums []int) int {
	result := 0
	for _,num := range nums {
		result ^= num
	}
	return result
}
