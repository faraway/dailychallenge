package main

import "fmt"

/**

Given a singly-linked list, reverse the list. This can be done iteratively or recursively. Can you get both solutions?

Example:
Input: 4 -> 3 -> 2 -> 1 -> 0 -> NULL
Output: 0 -> 1 -> 2 -> 3 -> 4 -> NULL

*/
type LinkedListNode struct {
	val int
	next *LinkedListNode //pointer to the next node. Simply 'listNode' won't work
}

func main() {
	fmt.Println("initializing test data...")
	head := &LinkedListNode{val: 4}
	head.next = &LinkedListNode{val: 3}
	head.next.next = &LinkedListNode{val: 2}
	head.next.next.next = &LinkedListNode{val: 1}
	head.next.next.next.next = &LinkedListNode{val: 0}
	newHead := reverseIteratively(head)
	printList(newHead)
	fmt.Println("\n-----------------------")
	newHead2 := reverseRecursively(newHead) //reverse again, should be back to original order
	printList(newHead2)
}

func printList(head *LinkedListNode) {
	for head != nil {
		fmt.Print(head.val, "->")
		head = head.next
	}
	fmt.Print("nil")
}

/**
Reverse the linked list iteratively
 */
func reverseIteratively(head *LinkedListNode) *LinkedListNode {
	if head == nil { return head } // nil list
	previous := head
	current := previous.next
	if current == nil { return head } //single node linkedList

	previous.next = nil //this is now the tail of the to be reversed linked list
	for current != nil {
		temp := current.next
		current.next = previous
		previous = current
		current = temp
	}
	//'current' will be nil, 'previous' will be the last valid node, which will be the head
	return previous
}

/**
Reverse the linked list recursively
 */
func reverseRecursively(head *LinkedListNode) *LinkedListNode {
	if head == nil { return nil }
	if head.next == nil { return head }

	newHead := reverseRecursively(head.next)
	head.next.next = head
	head.next = nil //break the current link
	return newHead
}
