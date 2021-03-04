package main

import (
	"fmt"
	"sort"
)

/**
Given a list of numbers, find if there exists a pythagorean triplet in that list.
A pythagorean triplet is 3 variables a, b, c where a^2 + b^2 = c^2

Example:
Input: [3, 5, 12, 5, 13]
Output: True
Here, 5^2 + 12^2 = 13^2.
 */
func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", findPythagoreanTriplets([]int{3,5,12,5,13})) //True
	fmt.Println("Answer is:", findPythagoreanTriplets([]int{3,1,4,6,5})) //True
	fmt.Println("Answer is:", findPythagoreanTriplets([]int{10,4,6,12,5})) //False
}

func findPythagoreanTriplets(input []int) bool {
	//convert the array to the array of num^2
	for i, num := range input {
		input[i] = num * num //assuming it won't overflow
	}
	sort.Ints(input)
	//now it's a problem similar to 3-sum.
	//3sum can be easily solved in O(n^2); but fancier algorithm exist that pushes the limits a little bit
	//https://en.wikipedia.org/wiki/3SUM
	for i := len(input)-1; i>=2; i-- {
		if checkTwoSum(input[:i], input[i]) {
			return true
		}
	}

	return false
}

func checkTwoSum(array []int, target int) bool {
	checkMap := make(map[int]bool)
	for _,num := range array {
		if checkMap[target-num] {
			return true
		} else {
			checkMap[num] = true
		}
	}
	return false
}