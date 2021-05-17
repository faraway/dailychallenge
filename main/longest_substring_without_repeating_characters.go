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

/**
Using sliding window solution (typical case for "substring" problems)
See: https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%8A%80%E5%B7%A7.md

> related to @find_all_anagrams_in_a_string.go and @minimum_window_substring.go
*/
func lengthOfLongestSubstring(input string) int {
	dupChecker := make(map[uint8]bool)
	maxLength := 0
	left, right := 0, 0

	for right < len(input) {

		for right < len(input) && !dupChecker[input[right]] {
			dupChecker[input[right]] = true
			right++
			if (right - left) > maxLength {
				maxLength = right - left
			}
		}
		for left < right && right < len(input) {
			dupChecker[input[left]] = false
			left++
			//this means we just moved over the char that's duplicate of "right" pointer has seen;
			//break out of look
			if input[left-1] == input[right] {
				break
			}
		}

	}
	return maxLength
}
