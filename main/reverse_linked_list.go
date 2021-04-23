package main

import  (
	. "dailychallenge/utils"
	"fmt"
)

/**

Given a singly-linked list, reverse the list. This can be done iteratively or recursively. Can you get both solutions?

Example:
Input: 4 -> 3 -> 2 -> 1 -> 0 -> NULL
Output: 0 -> 1 -> 2 -> 3 -> 4 -> NULL

*/

func main() {
	fmt.Println("initializing test data...")
	head := &ListNode{Value: 4}
	head.Next = &ListNode{Value: 3}
	head.Next.Next = &ListNode{Value: 2}
	head.Next.Next.Next = &ListNode{Value: 1}
	head.Next.Next.Next.Next = &ListNode{Value: 0}
	newHead := reverseIteratively(head)
	printList(newHead)
	fmt.Println("\n-----------------------")
	newHead2 := reverseRecursively(newHead) //reverse again, should be back to original order
	printList(newHead2)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Value, "->")
		head = head.Next
	}
	fmt.Print("nil")
}

/**
Reverse the linked list iteratively
 */
func reverseIteratively(head *ListNode) *ListNode {
	if head == nil { return head } // nil list
	previous := head
	current := previous.Next
	if current == nil { return head } //single node linkedList

	previous.Next = nil //this is now the tail of the to be reversed linked list
	for current != nil {
		temp := current.Next
		current.Next = previous
		previous = current
		current = temp
	}
	//'current' will be nil, 'previous' will be the last Valueid node, which will be the head
	return previous
}

/**
Reverse the linked list recursively
 */
func reverseRecursively(head *ListNode) *ListNode {
	if head == nil { return nil }
	if head.Next == nil { return head }

	newHead := reverseRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil //break the current link
	return newHead
}
