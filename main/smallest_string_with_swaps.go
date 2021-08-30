package main

import (
	. "dailychallenge/utils"
	"fmt"
	"sort"
)

/**
You are given a string s, and an array of pairs of indices in the string pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.
You can swap the characters at any pair of indices in the given pairs any number of times.

Return the lexicographically smallest string that s can be changed to after using the swaps.

Example 1:
Input: s = "dcab", pairs = [[0,3],[1,2]]
Output: "bacd"
Explaination:
Swap s[0] and s[3], s = "bcad"
Swap s[1] and s[2], s = "bacd"

Example 2:
Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
Output: "abcd"
Explaination:
Swap s[0] and s[3], s = "bcad"
Swap s[0] and s[2], s = "acbd"
Swap s[1] and s[2], s = "abcd"

Example 3:
Input: s = "cba", pairs = [[0,1],[1,2]]
Output: "abc"
Explaination:
Swap s[0] and s[1], s = "bca"
Swap s[1] and s[2], s = "bac"
Swap s[0] and s[1], s = "abc"


Constraints:
1 <= s.length <= 10^5
0 <= pairs.length <= 10^5
0 <= pairs[i][0], pairs[i][1] < s.length
s only contains lower case English letters.
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", smallestStringWithSwaps("dcab", [][]int{{0,3}, {1,2}})) //bacd
	fmt.Println("Answer is:", smallestStringWithSwaps("dcab", [][]int{{0,3}, {1,2}, {0,2}})) //abcd
	fmt.Println("Answer is:", smallestStringWithSwaps("cba", [][]int{{0,1}, {1,2}})) //abc
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	/**
		Since we can swap *any number of times*, this is still a Union-Find problem
		The basic idea is to identify the connected groups (indices of the string).
		Within each group, we sort the letters lexicographically. And then resemble the string from taking each group's smallest letter to largest
	 */
	ds := NewDisjointSet(len(s))
	for _, pair := range pairs {
		ds.Union(pair[0], pair[1])
	}

	//for each identified group (by root), find its member and their corresponding letter in the string
	//e.g. [ 0: {d,a}, 1: {c,b} ]
	groups := make(map[int][]int32)
	for i, char := range s {
		root := ds.Find(i)
		groups[root] = append(groups[root], char)
	}

	//sort each group's letters
	//e.g. [ 0: {a, d}, 1: {b,c} ]
	for root,_ := range groups {
		sort.Slice(groups[root], func(i, j int) bool {
			return groups[root][i] < groups[root][j]
		})
	}

	//reconstruct the string from taking each group's smallest letter to largest
	var result []int32
	for i:=0; i<len(s); i++ {
		root := ds.Find(i)
		result = append(result,groups[root][0])
		groups[root] = groups[root][1:] //take out the head of the list
	}
	return string(result)
}