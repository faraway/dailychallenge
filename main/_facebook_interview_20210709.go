package main

/**
Hello

The input is a string which contains open and/or close parentheses.
Find the minimum number of parentheses needed to be added in order to balance the input string.

Examples:

Input: ((( Output: 3
Input: ()) Output: 1
Input: (()) Output: 0
Input: )( Output: 2
**/
func minNumberOfParenthesesToBalance (input string) int {
	var stack Stack //push,pop,peak  //leftP := 0

	result := 0

	for _, char := range input {
		parentheses := string(char)
		if parentheses == "(" {
			stack.Push(parentheses)  //leftP++
		} else {
			if stack.size() > 0 {    //leftP > 0
				stack.Pop()          //leftP --
			} else {
				result++
			}
		}
	}

	return result + stack.Size() //result + leftP
}

// test cases:
// ""    => 0
// "(((" => 3
// "))"  => 2



/**
Return dot product of two sparse vectors.
Given two vectors a and b, the dot product is a1*b1 + a2*b2 + ... + aN*bN
A sparse vector is a vector with most elements equal to zero - imagine a vector with millions of elements

sparseVector = [1,0,0,0,0,2,0,0,3]

Compress vector:
1. using hashmap
2. [ [0,1], [5,2], [8,3] ]
**/

func dotProduct(vector1 []int, vector2 []int) int {
	result := 0

	c1 := compressVector(vector1)
	c2 := compressVector(vector2)

	p1, p2 := 0, 0

	for p1 < len(c1) && p2 < len(c2) { //for = while
		if c1[p1][0] == c2[p2][0] {
			result += c1[p1][1] * c2[p2][1]
			p1++
			p2++
		} else if c1[p1][0] > c2[p2][0] {
			p2++
		} else {
			p1++
		}
	}

	return result
}

func compressVector(vector []int) [][]int {
	var result [][]int

	for index, val := range vector {
		if val != 0 {
			result = append(result, []int{index,val})
		}
	}
	return result
}

//test cases:
// v1 = [] v2 = []
// v1 = [0,0,0] v2 = []
// v1 = [1,0,0] v2 = []
// v1 = [0,0,3] v2 = [0,3,0]  [2,3] [1,3]
// v1 = [0,0,3] v2 = [0,0,3]  => 9



