package main

import "fmt"

/**
You are given a string of parenthesis.
Return the minimum number of parenthesis that would need to be removed in order to make the string valid.
"Valid" means that each open parenthesis has a matching closed parenthesis.

Example:

"()())()"

The above input should return 1.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("()())()")) //1
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("[()]{}")) //True
	fmt.Println("Answer is:", minimumRemovalsForValidParenthesis("({[)]")) //False
}

func minimumRemovalsForValidParenthesis(input string) int {

	return 0
}

