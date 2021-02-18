package main

import "fmt"

type listNode struct {
	val int
	next *listNode //pointer to the next node. Simply 'listNode' won't work
}


func main() {
	fmt.Println("initializing test data...")
	//test case 1
	l1 := &listNode{val: 2}
	l1.next = &listNode{val: 4}
	l1.next.next = &listNode{val: 3}

	l2 := &listNode{val: 5}
	l2.next = &listNode{val: 6}
	l2.next.next = &listNode{val: 4}

	result1 := addTwoNumbers(l1, l2)
	printResult(result1) //7->0->8
	fmt.Println("\n--------------------")
	//test case 2
	l3 := &listNode{val: 1}

	l4 := &listNode{val: 9}
	l4.next = &listNode{val: 9}
	l4.next.next = &listNode{val: 9}

	result2 := addTwoNumbers(l3, l4)
	printResult(result2) //0->0->0->1
	fmt.Println("\n--------------------")
	//test case 3
	l5 := &listNode{val: 5}
	l5.next = &listNode{val: 6}


	l6 := &listNode{val: 2}
	l6.next = &listNode{val: 7}
	l6.next.next = &listNode{val: 1}

	result3 := addTwoNumbers(l5, l6)
	printResult(result3) //7->3->2
}

func printResult(result *listNode) {
	for result != nil {
		fmt.Print(result.val)
		result = result.next
		if result != nil {
			fmt.Print("->")
		}
	}
}

//It's much easier the nodes goes from 1s as first node and 10s as second node ...
//If it's highest digit on the first node of the linked list, then we probably should reverse the linked list first.
func addTwoNumbers(first *listNode, second *listNode) *listNode {
	var head, current *listNode
	carry := 0

	for (first != nil || second != nil || carry != 0) {
		sum := carry
		if first != nil {
			sum += first.val
			first = first.next
		}
		if second != nil {
			sum += second.val
			second = second.next
		}
		carry = sum/10
		if current == nil {
			head = &listNode{val: sum%10}
			current = head
		} else {
			current.next = &listNode{val: sum%10}
			current = current.next
		}
	}
	return head
}