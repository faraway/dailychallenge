package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
You are given two linked-lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */

func main() {
	fmt.Println("initializing test data...")
	//test case 1
	l1 := &ListNode{Value: 2}
	l1.Next = &ListNode{Value: 4}
	l1.Next.Next = &ListNode{Value: 3}

	l2 := &ListNode{Value: 5}
	l2.Next = &ListNode{Value: 6}
	l2.Next.Next = &ListNode{Value: 4}

	result1 := addTwoNumbers(l1, l2)
	printResult(result1) //7->0->8
	fmt.Println("\n--------------------")
	//test case 2
	l3 := &ListNode{Value: 1}

	l4 := &ListNode{Value: 9}
	l4.Next = &ListNode{Value: 9}
	l4.Next.Next = &ListNode{Value: 9}

	result2 := addTwoNumbers(l3, l4)
	printResult(result2) //0->0->0->1
	fmt.Println("\n--------------------")
	//test case 3
	l5 := &ListNode{Value: 5}
	l5.Next = &ListNode{Value: 6}


	l6 := &ListNode{Value: 2}
	l6.Next = &ListNode{Value: 7}
	l6.Next.Next = &ListNode{Value: 1}

	result3 := addTwoNumbers(l5, l6)
	printResult(result3) //7->3->2
}

func printResult(result *ListNode) {
	for result != nil {
		fmt.Print(result.Value)
		result = result.Next
		if result != nil {
			fmt.Print("->")
		}
	}
}

//It's much easier the nodes goes from 1s as first node and 10s as second node ...
//If it's highest digit on the first node of the linked list, then we probably should reverse the linked list first.
func addTwoNumbers(first *ListNode, second *ListNode) *ListNode {
	var head, current *ListNode
	carry := 0

	for first != nil || second != nil || carry != 0 {
		sum := carry
		if first != nil {
			sum += first.Value
			first = first.Next
		}
		if second != nil {
			sum += second.Value
			second = second.Next
		}
		carry = sum/10
		if current == nil {
			head = &ListNode{Value: sum%10}
			current = head
		} else {
			current.Next = &ListNode{Value: sum%10}
			current = current.Next
		}
	}
	return head
}