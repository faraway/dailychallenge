package utils

import (
	"container/heap"
	"fmt"
)

//package init function is optional, no parameter, no return value
func init() {
	fmt.Println("utils package initialized")
}

//-------------------------------List Node -----------------------------------
type ListNode struct {
	Value int
	Next *ListNode
}

//-------------------------------Tree Node-----------------------------------
type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

//Print the tree with in "pre order"
func (node *TreeNode) PrintPreorder() {
	if node == nil { return }
	fmt.Print(node.Value," ")
	node.Left.PrintPreorder()
	node.Right.PrintPreorder()
}

//Not that I don't know generics..Just for backwards compatibility
//(so that I don't have to refactor all existing code)
type TreeNodeStr struct {
	Value string
	Left *TreeNodeStr
	Right *TreeNodeStr
}

//Print the tree with in "pre order"
func (node *TreeNodeStr) PrintPreorder() {
	if node == nil { return }
	fmt.Print(node.Value," ")
	node.Left.PrintPreorder()
	node.Right.PrintPreorder()
}

//-------------------------------Stack-----------------------------------
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

//-------------------------------Set-----------------------------------
type Set struct {
	list map[int]struct{} //empty structs occupy 0 memory
}

func (s *Set) Has(v int) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v int) {
	s.list[v] = struct{}{}
}

func (s *Set) Remove(v int) {
	delete(s.list, v)
}

func (s *Set) GetAll() []int {
	result := make([]int, len(s.list))
	for key,_ := range s.list {
		result = append(result,key)
	}
	return result
}

func (s *Set) Clear() {
	s.list = make(map[int]struct{})
}

func (s *Set) Size() int {
	return len(s.list)
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[int]struct{})
	return s
}

//-------------------------------Interval-----------------------------------
type Interval struct {
	Start int
	End int
}

//-------------------------------IntHeap-----------------------------------
//from https://golang.org/pkg/container/heap/
//This IntHeap is a *min* heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


//Note that Push and Pop in this interface are for package "container/heap" implementation to call.
//To add and remove things from the heap, use heap.Push and heap.Pop.
//See TestHeap() below
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


//-------------------------------Test Func-----------------------------------
func TestFunc() {
	fmt.Println("calling package function...")
}

func TestHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
