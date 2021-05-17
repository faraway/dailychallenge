package main

import "fmt"

/**
https://leetcode.com/problems/minimum-window-substring/ [HARD]

Given two strings s and t of lengths m and n respectively,
return the minimum window in s which will contain all the characters in t.
If there is no such window in s that covers all characters in t, return the empty string "".

Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.



Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"

Example 2:
Input: s = "a", t = "a"
Output: "a"

Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of English letters.
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", minWindow("ADOBECODEBANC", "ABC")) //BANC
	fmt.Println("Answer is:", minWindow("a", "a")) //a
	fmt.Println("Answer is:", minWindow("a", "aa")) //""
}

/**
Using sliding window solution
See: https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%8A%80%E5%B7%A7.md

> related to @longest_substring_without_repeating_characters.go and @find_all_anagrams_in_a_string.go
 */
func minWindow(s string, t string) string {
	result := ""

	targetMap := make(map[uint8]int) //key is type uint8 because that's the ASCII of the character
	for i := range t {
		targetMap[t[i]]++
	}

	windowMap := make(map[uint8]int)
	left, right := 0, 0
	for right < len(s) {
		//move right pointer until we find it covers target
		for !containsAllTargetChars(&windowMap, &targetMap) && right < len(s){
			windowMap[s[right]]++
			right ++
		}
		//move left pointer to shrink the window until it no longer covers "t"
		for containsAllTargetChars(&windowMap, &targetMap) && left < right {
			newResult := s[left:right]
			if result == "" || len(newResult) < len(result) {
				result = newResult
			}
			windowMap[s[left]]--
			left ++
		}
	}
	return result
}

//This might be too expensive to do in each step, possible to be blended in the main loop and check individually
func containsAllTargetChars(windowMap *map[uint8]int, targetMap *map[uint8]int) bool {
	for key, _ := range *targetMap {
		if (*windowMap)[key] < (*targetMap)[key] {
			return false
		}
	}
	return true
}