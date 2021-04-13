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
