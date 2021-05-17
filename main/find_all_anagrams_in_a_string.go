package main

import "fmt"

/**
https://leetcode.com/problems/find-all-anagrams-in-a-string/

Given two strings s and p, return an array of all the start indices of p's anagrams in s.
You may return the answer in any order.

Example 1:
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".


Constraints:
1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.

NOTE: An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

func main() {
	fmt.Println("initializing test data...")
	fmt.Println("Answer is:", findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println("Answer is:", findAnagrams("abab", "ab")) //[0,1,2]
}

/**
 Using sliding window approach here, but it's much simpler because the window size is fixed in this case.
 See: https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%8A%80%E5%B7%A7.md

 > related to @longest_substring_without_repeating_characters.go and @minimum_window_substring.go
 */
func findAnagrams(s string, p string) []int {
	var results []int

	if len(s) < len(p) {
		return results
	}

	targetCounter := make(map[uint8]int)
	windowCounter := make(map[uint8]int)

	//left is inclusive, right is not. [left, right)
	left, right := 0, len(p)
	for i := range p {
		targetCounter[p[i]]++
	}
	for i := 0; i < right; i++ {
		windowCounter[s[i]]++
	}
	if findAnagramsCounterEquals(windowCounter, targetCounter) {
		results = append(results, left)
	}
	//move the window
	for right < len(s) {
		windowCounter[s[left]]--
		//remove key if 0, makes it easier for comparing the counters
		if windowCounter[s[left]] == 0 {
			delete(windowCounter, s[left])
		}
		left++

		windowCounter[s[right]]++ //right is not inclusive, so update counter first then ++
		right++
		if findAnagramsCounterEquals(windowCounter, targetCounter) {
			results = append(results, left)
		}
	}
	return results
}

func findAnagramsCounterEquals(a map[uint8]int, b map[uint8]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key := range a {
		if a[key] != b[key] {
			return false
		}
	}
	return true
}