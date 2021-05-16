package main

import (
	"fmt"
)

/**
https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string s, find the length of the longest *substring* without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Example 4:
Input: s = ""
Output: 0

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

*/

func main() {
	fmt.Println("initializing test data...")
	//test case 1
	fmt.Println("Answer is:", lengthOfLongestSubstring("abrkaabcdefghijjxxx")) //10
	fmt.Println("Answer is:", lengthOfLongestSubstring("aaaaabbbbbbbcccccc")) //2
	fmt.Println("Answer is:", lengthOfLongestSubstring("abcdabcddddddabcefgda")) //7
	fmt.Println("Answer is:", lengthOfLongestSubstring("pwwkew")) //3
	fmt.Println("Answer is:", lengthOfLongestSubstring(" ")) //1
	fmt.Println("Answer is:", lengthOfLongestSubstring("dvdf")) //3
}


func lengthOfLongestSubstring(input string) int {
	checkMap := make(map[int32]bool)
	maxLength, currentLength := 0, 0

	for _, char := range input {
		//fmt.Println(idx, "=>", string(char))
		if !checkMap[char] {
			currentLength ++
			checkMap[char] = true
		} else {
			//clear the map, or maybe just create a new map. i.e. checkMap := make(map[int32]bool)
			for k := range checkMap {
				delete(checkMap, k)
			}
			currentLength = 1
			checkMap[char] = true
		}

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
