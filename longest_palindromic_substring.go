package main

import "fmt"

/**
A palindrome is a sequence of characters that reads the same backwards and forwards.
Given a string, s, find the longest palindromic substring in s.

Example:
Input: "banana"
Output: "anana"

Input: "million"
Output: "illi"

//TODO: linear time optimal solution: Manacher's algorithm
*/

func main() {
	fmt.Println("initializing test data...")
	//test case 1
	fmt.Println("Answer is:", longestPalindrome("banana")) //'anana'
	fmt.Println("Answer is:", longestPalindrome("million")) //'illi'
}

//O(n^2) solution
func longestPalindrome(input string) string {
    longestPalindrome := ""
    for index, _ := range input {
		//make current char as center
    	step := 1
    	for index-step>=0 && index+step<len(input) && input[index-step]==input[index+step] {
    		step++
		}
		length := step*2-1
		if length > len(longestPalindrome) {
			longestPalindrome = input[index-step+1:index+step]
		}
		// make the center to be between current and next char; essentially
		// current and next char are 'center' only if they are the same
		step = 0
		for index-step>=0 && index+1+step<len(input) && input[index-step]==input[index+1+step] {
			step++
		}
		length = step*2
		if length > len(longestPalindrome) {
			longestPalindrome = input[index-step+1:index+step+1]
		}
	}
	return longestPalindrome
}
