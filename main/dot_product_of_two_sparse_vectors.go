package main

import "fmt"

/**
https://leetcode.com/problems/dot-product-of-two-sparse-vectors/

Given two sparse vectors, compute their dot product.

Implement class SparseVector:

SparseVector(nums) Initializes the object with the vector nums
dotProduct(vec) Compute the dot product between the instance of SparseVector and vec
A sparse vector is a vector that has mostly zero values,
you should store the sparse vector efficiently and compute the dot product between two SparseVector.

Follow up: What if only one of the vectors is sparse?

Example 1:
Input: nums1 = [1,0,0,2,3], nums2 = [0,3,0,4,0]
Output: 8
Explanation: v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
v1.dotProduct(v2) = 1*0 + 0*3 + 0*0 + 2*4 + 3*0 = 8

Example 2:
Input: nums1 = [0,1,0,0,0], nums2 = [0,0,0,0,2]
Output: 0
Explanation: v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
v1.dotProduct(v2) = 0*0 + 1*0 + 0*0 + 0*0 + 0*2 = 0

Example 3:
Input: nums1 = [0,1,0,0,2,0,0], nums2 = [1,0,0,0,3,0,4]
Output: 6
*/

func main() {
	fmt.Println("initializing test data...")
	v1 := NewSparseVector([]int{1,0,0,2,3})
	v2 := NewSparseVector([]int{0,3,0,4,0})
	fmt.Println("Answer is:", v1.dotProduct(v2)) //8
}

type SparseVector struct {
	val map[int]int
}

func NewSparseVector(nums []int) SparseVector {
	val := make(map[int]int)
	for idx, num := range nums {
		if num != 0 {
			val[idx] = num
		}
	}
	return SparseVector{val: val}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	base := this
	target := &vec
	result := 0
	//take care of the follow up question, in case one is much more sparse than the other
	if len(vec.val) < len(this.val) {
		base = &vec
		target = this
	}
	for idx, val := range base.val {
		num, exists := target.val[idx]
		if exists {
			result += val * num
		}
	}
	return result
}
