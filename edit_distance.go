package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
Given two strings, determine the edit distance between them.
The edit distance is defined as the minimum number of edits (insertion, deletion, or substitution)
needed to change one string to the other.

For example, "biting" and "sitting" have an edit distance of 2 (substitute b for s, and insert a t).

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", getEditDistance("word", "word")) //0
	fmt.Println("Answer is:", getEditDistance("biting", "sitting")) //2
	fmt.Println("Answer is:", getEditDistance("horse", "ros")) //3
	fmt.Println("Answer is:", getEditDistance("intention", "execution")) //5
	fmt.Println("Answer is:", getEditDistance("apple", "sixty")) //5
	fmt.Println("Answer is:", getEditDistance("java", "javascript")) //6
	fmt.Println("Answer is:", getEditDistance("", "golang")) //6
	fmt.Println("Answer is:", getEditDistance("ruby", "")) //4
	fmt.Println("Answer is:", getEditDistance("sea", "eat")) //2
	fmt.Println("Answer is:", getEditDistance("a", "ab")) //1

}

func getEditDistance(s1 string, s2 string) int {
	m,n := len(s1), len(s2)

	//if either of them is an empty string, edit distance will be the length of the other
	if m*n == 0 { return m|n }

	//initialize a 2D matrix, in Go there is not a more concise way...
	//ref https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	//2D dynamic programming
	for i := 0; i< m; i++ {
		for j := 0; j< n; j++ {
			charMatch := 1
			if s1[i] == s2[j] {
				charMatch = 0
			}
			if i==0 && j==0 {
				matrix[i][j] = charMatch
			} else if i==0 {
				matrix[i][j] = MIN(matrix[i][j-1]+1, j+charMatch)
			} else if j==0 {
				matrix[i][j] = MIN(matrix[i-1][j]+1, i+charMatch)
			} else {
				matrix[i][j] = MIN(matrix[i][j-1]+1, matrix[i-1][j]+1, matrix[i-1][j-1]+charMatch)
			}
			//fmt.Println("Matrix is: ", matrix)
		}
	}
	return matrix[m-1][n-1]
}
