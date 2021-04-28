package main

import (
	. "dailychallenge/utils"
	"fmt"
	"strconv"
)

/**
https://leetcode.com/problems/decode-string/

Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being
repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for
those repeat numbers, k. For example, there won't be input like 3a or 2[4].

Example 1:
Input: s = "3[a]2[bc]"
Output: "aaabcbc"

Example 2:
Input: s = "3[a2[c]]"
Output: "accaccacc"

Example 3:
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"

Example 4:
Input: s = "abc3[cd]xyz"
Output: "abccdcdcdxyz"
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", decodeString("3[a]2[bc]")) //aaabcbc
	fmt.Println("Answer is:", decodeString("3[a2[c]]")) //accaccacc
	fmt.Println("Answer is:", decodeString("2[abc]3[cd]ef")) //abcabccdcdcdef
	fmt.Println("Answer is:", decodeString("abc3[cd]xyz")) //abccdcdcdxyz
}

func decodeString(s string) string {
	var result string
	var stack Stack

	for _, char := range s {
		token := string(char)
		switch  {
		case (char>=65 && char<=90) || (char>=97 && char<=122): //letters
			if stack.IsEmpty() {
				result += token
			} else {
				stack.Push(token)
			}
		case char >= 48 && char <= 57: //number
			stack.Push(token)
		case char == 91: //[
			stack.Push(token)
		case char == 93: //]
			var tempString string
			next, _ := stack.Pop()
			for next.(string) != "[" {
				tempString = next.(string) + tempString
				next, _ = stack.Pop()
			}
			next, _ = stack.Pop()
			count, _ := strconv.Atoi(next.(string)) //TODO: need better handling here, this only works for single digit number
			decodedString := repeatString(tempString, count)
			if (stack.IsEmpty()) {
				result += decodedString
 			} else { //nested "[" and "]" exists; put it back to stack and wait for further resolution until next "]"
				stack.Push(decodedString)
			}
		}
	}
	return result
}

func repeatString(s string, times int) string {
	var result string
	for times > 0 {
		result += s
		times --
	}
	return result
}