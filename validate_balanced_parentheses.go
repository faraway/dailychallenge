package main

import "fmt"


/**
Imagine you are building a compiler. Before running any code, the compiler must check that the parentheses in the program are balanced. Every opening bracket must have a corresponding closing bracket. We can approximate this using strings.

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
- Open brackets are closed by the same type of brackets.
- Open brackets are closed in the correct order.
- Note that an empty string is also considered valid.
*/

func main() {
	fmt.Println("initializing test data...")
	//test case 1
	fmt.Println("Answer is:", validateBalancedParentheses("((()))")) //True
	fmt.Println("Answer is:", validateBalancedParentheses("[()]{}")) //True
	fmt.Println("Answer is:", validateBalancedParentheses("({[)]")) //False
	fmt.Println("Answer is:", validateBalancedParentheses("")) //True
	fmt.Println("Answer is:", validateBalancedParentheses("{}[](())[{}]")) //True
	fmt.Println("Answer is:", validateBalancedParentheses("({[]})[[")) //False
	fmt.Println("Answer is:", validateBalancedParentheses("({[]})}")) //False
}


func validateBalancedParentheses(input string) bool {
	//an empty string is also considered valid
	if len(input) == 0 { return true }
	var stack Stack

	for _, char := range input {
		item := string(char)
		if item == "(" || item == "[" || item == "{" {
			stack.Push(item)
		} else {
			pair, success := stack.Pop()
			if !success || !isPair(item, pair.(string)) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func isPair(closingItem string, item string) bool {
	if closingItem == ")" {
		return item == "("
	} else if closingItem == "]" {
		return item == "["
	} else if closingItem == "}" {
		return item == "{"
	} else {
		return false
	}
}

/* Definition of stack */
type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		lastIndex := len(*s) - 1
		item := (*s)[lastIndex]
		*s = (*s)[:lastIndex]
		return item, true
	}
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		return (*s)[len(*s) - 1], true
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}
