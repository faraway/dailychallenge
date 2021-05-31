package main

import (
	"container/heap"
	"fmt"
)

/**
https://leetcode.com/problems/task-scheduler/

Given a characters array tasks, representing the tasks a CPU needs to do, where each letter represents a different task.
Tasks could be done in any order. Each task is done in one unit of time.
For each unit of time, the CPU could complete either one task or just be idle.

However, there is a non-negative integer n that represents the cooldown period between two same
tasks (the same letter in the array), that is that there must be at least n units of time between any two same tasks.

Return the least number of units of times that the CPU will take to finish all the given tasks.

Example 1:
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation:
A -> B -> idle -> A -> B -> idle -> A -> B
There is at least 2 units of time between any two same tasks.


Example 2:
Input: tasks = ["A","A","A","B","B","B"], n = 0
Output: 6
Explanation: On this case any permutation of size 6 would work since n = 0.
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
And so on.


Example 3:
Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
Output: 16
Explanation:
One possible solution is
A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A


Constraints:
1 <= task.length <= 104
tasks[i] is upper-case English letter.
The integer n is in the range [0, 100].

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)) //8
	fmt.Println("Answer is:", leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0)) //6
	fmt.Println("Answer is:", leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)) //16
}

func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]*Task)
	for _,task := range tasks {
		var t *Task
		if taskMap[task] == nil {
			t = &Task{task, 1, 0}
			taskMap[task] = t
		} else {
			taskMap[task].count++
		}
	}
	var h TaskHeap
	for _, task := range taskMap {
		h = append(h, task)
	}
	heap.Init(&h)
	timer := 0
	//this is the simulation of the CPU takes the task and run them one by one on each time slice
	for h.Len() > 0 {
		if h[0].coolDown == 0 {
			fmt.Print(string(h[0].id), ",")
			if h[0].count == 1 {
				heap.Pop(&h)
			} else {
				h[0].count--
				h[0].coolDown = n+1 //count for global cool down
			}
			heap.Fix(&h, 0)
		}
		for i, task := range h {
			if task.coolDown > 0 {
				task.coolDown--
			}
			if task.coolDown == 0 {
				heap.Fix(&h, i) //need to fix and pop up the tasks with cool down = 0 and high counts
			}
		}
		timer++
	}
	return timer
}

type Task struct {
	id byte
	count int
	coolDown int
}

type TaskHeap []*Task

func (this TaskHeap) Len() int {
	return len(this)
}

func (this TaskHeap) Less(i int, j int) bool {
	if this[i].coolDown < this[j].coolDown {
		return true
	}
	if this[i].coolDown == this[j].coolDown {
		return this[i].count > this[j].count
	}
	return false
}

func (this TaskHeap) Swap(i int, j int) {
	this[i], this[j] = this[j], this[i]
}

func (h *TaskHeap) Push(x interface{}) {
	(*h) = append(*h, x.(*Task))
}

func (h *TaskHeap) Pop() interface{} {
	last := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return last
}