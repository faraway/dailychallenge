package main

import "fmt"

/**
Given a mathematical expression with just single digits, plus signs, negative signs, and brackets,
evaluate the expression.
Assume the expression is properly formed.

Example:
Input: - ( 3 + ( 2 - 1 ) )
Output: -4
 */

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", calculatorEval("- (3 + ( 2 - 1 ) )")) //-4
	fmt.Println("Answer is:", calculatorEval("(1+2)-4")) //-1
	fmt.Println("Answer is:", calculatorEval("4+1-3")) //2
}

//TODO: fill your solution
func calculatorEval(expression string) int {

	return 0
}
