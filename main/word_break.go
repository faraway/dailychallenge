package main

import "fmt"

/**
https://leetcode.com/problems/word-break/

Given a string s and a dictionary of strings wordDict,
return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false


Constraints:

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", wordBreak("leetcode", []string{"leet", "code"})) //True
	fmt.Println("Answer is:", wordBreak("applepenapple", []string{"apple", "pen"})) //True
	fmt.Println("Answer is:", wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) //False
}
func wordBreak(s string, wordDict []string) bool {
	//map the dictionary a little bit faster to look up
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	//exceeds time limit
	//wordBreakRecursive(s, dict, 0)

	//dp way of solving this problem, making the dp array len(s) + 1 may make things a little bit easier
	dp := make([]bool, len(s))
	for i := range dp {
		for j := i; j>=1; j-- {
			if dp[j-1] && dict[s[j:i+1]]{
				dp[i] = true
				break
			}
		}
		if !dp[i] && dict[s[0:i+1]] {
			dp[i] = true
		}
	}
	return dp[len(s)-1]
}

/**
Brute force approach. Potentially re-evaluate a lot of same (sub)strings over and over.
An easy improvement is to use a cache to memorize those intermediate results, otherwise this will exceed time limit
 */
func wordBreakRecursive(s string, dict map[string]bool, start int) bool {
	if start >= len(s) {
		return true
	}

	for i := start+1; i <= len(s); i++ {
		candidate := s[start:i]
		if dict[candidate] && wordBreakRecursive(s, dict, i) {
			return true
		}
	}
	return false
}