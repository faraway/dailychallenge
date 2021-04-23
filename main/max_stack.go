package main

import "fmt"

/**
Implement a class for a stack that supports all the regular functions (push, pop)
and an additional function of max() which returns the maximum element in the stack (return None if the stack is empty).
Each method should run in constant time - O(1) time

*/

func main() {
	fmt.Println("initializing test data...")

	var maxStack = MaxStack{mainStack: []int{}, maxStack: []int{}}
	maxStack.push(2)
	maxStack.push(5)
	maxStack.push(4)
	maxStack.push(7)
	maxStack.push(9)
	maxStack.push(3)
	fmt.Println("Current max is:", maxStack.max()) //9
	fmt.Println("Pop item:", maxStack.pop()) //3
	fmt.Println("Current max is:", maxStack.max()) //9
	fmt.Println("Pop item:", maxStack.pop()) //9
	fmt.Println("Current max is:", maxStack.max()) //7
	fmt.Println("Pop item:", maxStack.pop()) //7
	fmt.Println("Pop item:", maxStack.pop()) //4
	fmt.Println("Pop item:", maxStack.pop()) //5
	fmt.Println("Current max is:", maxStack.max()) //2
	fmt.Println("Pop item:", maxStack.pop()) //2
	//fmt.Println("Pop item:", maxStack.pop()) //panic
}


/**
 Implementation of max stack.
 It uses two stacks:
    main stack - stores elements as a normal stack
	max stack - stores elements of current max, in the same order elements are stored in main stack
e.g.  mainStack   maxStack
         3           9
         9           9
         7           7
         4           5
         5           5
         2           2
 */
type MaxStack struct {
	mainStack []int
	maxStack []int
}

func (s *MaxStack) push(item int) {
	s.mainStack = append(s.mainStack, item)
	length := len(s.maxStack)
	if length==0 || s.maxStack[length-1] < item {
		s.maxStack = append(s.maxStack, item)
	} else {
		s.maxStack = append(s.maxStack, s.maxStack[length-1])
	}
}

func (s *MaxStack) pop() int {
	if s.isEmpty() {
		panic("empty stack!")
	} else {
		item := s.mainStack[s.size()-1]
		s.mainStack = s.mainStack[:s.size()-1]
		s.maxStack = s.maxStack[:s.size()-1]
		return item
	}
}

func (s *MaxStack) max() int {
	if s.isEmpty() {
		panic("empty stack!")
	} else {
		return s.maxStack[s.size()-1]
	}
}

func (s *MaxStack) isEmpty() bool {
	return s.size() == 0
}

func (s *MaxStack) size() int {
	return len(s.maxStack)
}