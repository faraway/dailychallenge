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
	fmt.Println("Answer is:", minWindow("a", "aa")) //a
}

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
			windowMap[s[left]]--
			left ++
		}
		newResult := s[left-1:right] //back off one step on the left side to capture current result
		if result == "" || len(newResult) < len(result) {
			result = newResult
		}
	}
	return result
}

func containsAllTargetChars(windowMap *map[uint8]int, targetMap *map[uint8]int) bool {
	for key, _ := range *targetMap {
		if (*windowMap)[key] < (*targetMap)[key] {
			return false
		}
	}
	return true
}