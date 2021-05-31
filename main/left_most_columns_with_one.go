package main

import "fmt"

/**
https://leetcode.com/problems/leftmost-column-with-at-least-a-one/

A row-sorted binary matrix means that all elements are 0 or 1 and each row of the matrix is sorted in non-decreasing order.

Given a row-sorted binary matrix binaryMatrix, return the index (0-indexed) of the leftmost column with a 1 in it.
If such an index does not exist, return -1.

You can't access the Binary Matrix directly. You may only access the matrix using a BinaryMatrix interface:

BinaryMatrix.get(row, col) returns the element of the matrix at index (row, col) (0-indexed).
BinaryMatrix.dimensions() returns the dimensions of the matrix as a list of 2 elements [rows, cols], which means the matrix is rows x cols.
Submissions making more than 1000 calls to BinaryMatrix.get will be judged Wrong Answer.
Also, any solutions that attempt to circumvent the judge will result in disqualification.

For custom testing purposes, the input will be the entire binary matrix mat. You will not have access to the binary matrix directly.

Example 1:
Input: mat = [
	[0,0],
	[1,1]
]
Output: 0

Example 2:
Input: mat = [
	[0,0],
	[0,1]
]
Output: 1

Example 3:
Input: mat = [
	[0,0],
	[0,0]
]
Output: -1

Example 4:
Input: mat = [
	[0,0,0,1],
	[0,0,1,1],
	[0,1,1,1]
]
Output: 1


Constraints:
rows == mat.length
cols == mat[i].length
1 <= rows, cols <= 100
mat[i][j] is either 0 or 1.
mat[i] is sorted in non-decreasing order.

*/

func main() {
	fmt.Println("initializing test data...")

	matrix := BinaryMatrix {
		Get: func(i int,j int) int {
			data := [][]int{
				{0,0,0,1},
				{0,0,1,1},
				{0,1,1,1},
			}
			return data[i][j]
		},
		Dimensions: func() []int {
			return []int{3,4}
		},
	}
	fmt.Println("Answer is:", leftMostColumnWithOne(matrix)) //1
}

type BinaryMatrix struct {
	Get func(int, int) int
	Dimensions func() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()
	m, n := dimensions[0], dimensions[1]

	leftMostColumn := -1
	for i := 0; i < m; i ++ {
		if binaryMatrix.Get(i, n-1) == 0 { //no 1s for sure in this row
			continue
		} else {
			result := binarySearchOne(binaryMatrix, i, 0, n-1)
			if leftMostColumn == -1 {
				leftMostColumn= result
			} else if result != -1 && result < leftMostColumn {
				leftMostColumn = result
			}

		}
	}
	return leftMostColumn
}

func binarySearchOne(matrix BinaryMatrix, m int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if matrix.Get(m, mid) == 0 {
		return binarySearchOne(matrix, m, mid+1, end)
	} else {
		betterResult := binarySearchOne(matrix, m, start, mid-1)
		if betterResult >= 0 {
			return betterResult
		} else {
			return mid
		}
	}
}