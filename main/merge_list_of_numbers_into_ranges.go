package main


import (
	"fmt"
	"strconv"
)

/**
https://leetcode.com/problems/summary-ranges/

Given a *sorted* list of *unique* numbers, return a list of strings that represent all of the consecutive numbers.

Example:
Input: [0, 1, 2, 5, 7, 8, 9, 10, 11, 15]
Output: ['0->2', '5', '7->11', '15']

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", mergeListOfNumIntoRanges([]int{})) //[]
	fmt.Println("Answer is:", mergeListOfNumIntoRanges([]int{0, 1, 2, 5, 7, 8, 9, 10, 11, 15})) //['0->2', '5', '7->11', '15']
	fmt.Println("Answer is:", mergeListOfNumIntoRanges([]int{0, 1, 2, 5, 7, 8, 9 })) //['0->2', '5', '7->9']
	fmt.Println("Answer is:", mergeListOfNumIntoRanges([]int{1, 3, 5, 7, 9 })) //['1', '3', '5', '7', '9']
	fmt.Println("Answer is:", mergeListOfNumIntoRanges([]int{1, 2, 3, 4, 5 })) //['1->5']
}

func mergeListOfNumIntoRanges(sortedNums []int) []string {
	var result []string
	var start, previous int

	for i, num := range sortedNums {
		if i == 0 {
			start = num
			previous = num
		} else {
			if num != previous + 1 {
				result = append(result, getString(start, previous))
				start = num
				previous = num
			} else {
				previous = num
			}
		}
		//handle the ending
		if i == len(sortedNums) - 1 {
			result = append(result, getString(start, previous))
		}
	}

	return result
}

func getString(start int, previous int) string {
	if start == previous {
		return strconv.Itoa(start)
	} else {
		return strconv.Itoa(start)+"->"+strconv.Itoa(previous)
	}
}
