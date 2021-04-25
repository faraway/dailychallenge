package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.com/problems/reverse-integer/
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21

Example 4:
Input: x = 0
Output: 0
 */
func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", reverseInteger(123)) //321
	fmt.Println("Answer is:", reverseInteger(-123)) //-321
	fmt.Println("Answer is:", reverseInteger(120)) //21
	fmt.Println("Answer is:", reverseInteger(0)) //0
	fmt.Println("Answer is:", reverseInteger(1534236469)) //0  - Note: because this overflows int32
}

func reverseInteger(num int) int {
	//NOTE: It's possible to combine the two for loops below to make it more elegant
	//NOTE: These digits are with signs. e.g. '1', '2', '3' or '-1', '-2', '-3' (as "%" operation takes care of that)
	//So there is no need for any extra handling for negative numbers.
	var digits []int
	for num != 0 {
		digits = append(digits, num%10)
		num = num/10
	}

	result := 0
	for i:=0; i<len(digits); i++ {
		// int32 min: -2147483648
		// int32 max: 2147483647
		// hence the check
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digits[i] > 7)||
			result < math.MinInt32/10 || (result == math.MinInt32/10 && digits[i] < -8) {
			return 0
		}
		result = result * 10
		result += digits[i]
	}
	//This is kinda cheating :)
	//if result < math.MinInt32 || result > math.MaxInt32 {
	//	return 0
	//}
	return result
}
