package main

import (
	. "dailychallenge/utils"
	"fmt"
	"strconv"
)

/**
You are given a binary tree representation of an arithmetic expression.
In this tree, each leaf is an integer value,,
and a non-leaf node is one of the four operations: '+', '-', '*', or '/'.

Write a function that takes this tree and evaluates the expression.

Example:

    *
   / \
  +    +
 / \  / \
3  2  4  5

This is a representation of the expression (3 + 2) * (4 + 5), and should return 45.

*/

func main() {
	fmt.Println("initializing test data...")

	root := &TreeNodeStr{Value: "*"}
	root.Left = &TreeNodeStr{Value: "+"}
	root.Right = &TreeNodeStr{Value: "+"}

	root.Left.Left = &TreeNodeStr{Value: "3"}
	root.Left.Right = &TreeNodeStr{Value: "2"}

	root.Right.Left = &TreeNodeStr{Value: "4"}
	root.Right.Right = &TreeNodeStr{Value: "5"}

	fmt.Println("Answer is:", evaluateArithmeticExpression(root)) //45

}

func evaluateArithmeticExpression(root *TreeNodeStr) int {
	if !isOperator(root.Value) {
		i, _ := strconv.Atoi(root.Value)
		return i
	}
	left := evaluateArithmeticExpression(root.Left)
	right := evaluateArithmeticExpression(root.Right)
	var result int
	switch root.Value {
	case "+": result = left + right
	case "-": result = left - right
	case "*": result = left * right
	case "/": result = left / right
	}
	return result
}

func isOperator(val string) bool {
	return val == "+" || val == "-" || val == "*" || val == "/"
}
