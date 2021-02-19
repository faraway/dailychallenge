package main

import "fmt"

/**
Given a string, find the length of the longest substring without repeating characters.
Example:
'abrkaabcdefghijjxxx' => 10
*/

func main() {
	fmt.Println("initializing test data...")
	//test case 1
	fmt.Println("Answer is:", lengthOfLongestSubstring("abrkaabcdefghijjxxx")) //10
	fmt.Println("Answer is:", lengthOfLongestSubstring("aaaaabbbbbbbcccccc")) //2
	fmt.Println("Answer is:", lengthOfLongestSubstring("abcdabcddddddabcefgda")) //7
}


func lengthOfLongestSubstring(input string) int {
	checkMap := make(map[int32]bool)
	maxLength := 0
	currentLength := 0
	for _, char := range input {
		//fmt.Println(idx, "=>", string(char))
		if !checkMap[char] {
			currentLength ++
			checkMap[char] = true
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 0
			//clear the map, or maybe just create a new map. i.e. checkMap := make(map[int32]bool)
			for k := range checkMap {
				delete(checkMap, k)
			}
		}
	}
	return maxLength
}
