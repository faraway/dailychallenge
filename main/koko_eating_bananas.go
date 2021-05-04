package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.com/problems/koko-eating-bananas/

Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k.
Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas,
she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
Return the minimum integer k such that she can eat all the bananas within h hours.

Example 1:
Input: piles = [3,6,7,11], h = 8
Output: 4

Example 2:
Input: piles = [30,11,23,4,20], h = 5
Output: 30

Example 3:
Input: piles = [30,11,23,4,20], h = 6
Output: 23

Constraints:

1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109

*/

func main() {
	fmt.Println("initializing test data...")
	fmt.Println("Answer is:", minEatingSpeed([]int{3,6,7,11}, 8)) //4
	fmt.Println("Answer is:", minEatingSpeed([]int{30,11,23,4,20}, 5)) //30
	fmt.Println("Answer is:", minEatingSpeed([]int{30,11,23,4,20},  6)) //23
}

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	//the maximum speed koko ever need is 'max', the minimum speed is essentially 1.
	//now do the binary search between 1 and max
	return binarySearch(1, max, piles, h)
}

func binarySearch(start int, end int, piles []int, h int) int {
	//base case
	if start > end {
		return -1
	} else if start == end {
		if isValid(piles, h, start) {
			return start
		} else {
			return -1
		}
	}
    //recursive case
	mid := (start+end)/2
	currentResult := isValid(piles, h, mid)
	if !currentResult { //result must be in the right side
		return binarySearch(mid+1, end, piles, h)
	} else {
		potentialResult := binarySearch(start, mid-1, piles, h)
		if potentialResult > 0 {
			return potentialResult
		} else {
			return mid
		}
	}
}

func isValid(piles []int, h int, k int) bool {
	hours := 0
	for _, pile := range piles {
		hours += int(math.Ceil(float64(pile)/float64(k)))
	}
	return hours <= h
}