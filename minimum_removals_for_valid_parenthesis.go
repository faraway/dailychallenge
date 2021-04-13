package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
You are given a string of parenthesis.
Return the minimum number of parenthesis that would need to be removed in order to make the string valid.
"Valid" means that each open parenthesis has a matching closed parenthesis.

Example:

"()())()"

The above input should return 1.

** ALSO SEE BELOW WITH A LEET CODE VARIATION **

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("()())()")) //1
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis(")(")) //2
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("((()))")) //0
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("(())))")) //2
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("((())")) //1
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("(()))(()")) //2
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("(())()()))")) //2

	fmt.Println("Answer is:", leetcode_minimumRemovalsForValidParenthesis("lee(t(c)o)de)")) //"lee(t(c)o)de"
	fmt.Println("Answer is:", leetcode_minimumRemovalsForValidParenthesis("a)b(c)d")) //"ab(c)d"
	fmt.Println("Answer is:", leetcode_minimumRemovalsForValidParenthesis("))((")) //""
	fmt.Println("Answer is:", leetcode_minimumRemovalsForValidParenthesis("(a(b(c)d)")) //"a(b(c)d)"
}

func minimumRemovalsForValidParenthesis(input string) int {

	//an empty string is also considered valid
	if len(input) == 0 { return 0 }

	var stack Stack
	count := 0

	for _, char := range input {
		item := string(char)
		if item == "(" {
			stack.Push(item)
		} else if item == ")" {
			_, success := stack.Pop()
			if !success  {
				count ++
			}
		} else {
			fmt.Println(">>>> you should not be here")
		}
	}
	//un-matching ")" + whatever left in the stack of un-matching "("
	return count + stack.Size()
}

/**
 A similar problem on leetcode
  	https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

	Given a string s of '(' , ')' and lowercase English characters.

	Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.

	Formally, a parentheses string is valid if and only if:

	It is the empty string, contains only lowercase characters, or
	It can be written as AB (A concatenated with B), where A and B are valid strings, or
	It can be written as (A), where A is a valid string.


	Example 1:

	Input: s = "lee(t(c)o)de)"
	Output: "lee(t(c)o)de"
	Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
	Example 2:

	Input: s = "a)b(c)d"
	Output: "ab(c)d"
	Example 3:

	Input: s = "))(("
	Output: ""
	Explanation: An empty string is also valid.
	Example 4:

	Input: s = "(a(b(c)d)"
	Output: "a(b(c)d)"
 */
func leetcode_minimumRemovalsForValidParenthesis(input string) string {

	//an empty string is also considered valid
	if len(input) == 0 { return "" }

	var stack Stack
	indexList := make(map[int]bool)

	for index, char := range input {
		item := string(char)
		if item == "(" {
			stack.Push(index)
		} else if item == ")" {
			_, success := stack.Pop()
			if !success  {
				indexList[index] = true
			}
		} else {
			fmt.Println(">>>> normal english letters, skipping...")
		}
	}

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		indexList[item.(int)] = true
	}
	//Create the new string by ignoring the character in the string with identified index
	result := ""
	for index, char := range input {
		if !indexList[index] {
			result += string(char)
		}
	}
	return result
}