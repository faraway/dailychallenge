package main

import (
	"container/heap"
	"fmt"
	"sort"
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
	fmt.Println("Answer is:", leastInterval([]byte{'A', 'A', 'A', 'A', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'D'}, 1)) //12
	fmt.Println("Answer is:", leastInterval([]byte{'A', 'A', 'A', 'A', 'A','B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'D'}, 2)) //14

	fmt.Println("Answer is:", leastInterval_optimal([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)) //8
	fmt.Println("Answer is:", leastInterval_optimal([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0)) //6
	fmt.Println("Answer is:", leastInterval_optimal([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)) //16
	fmt.Println("Answer is:", leastInterval_optimal([]byte{'A', 'A', 'A', 'A', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'D'}, 1)) //12
	fmt.Println("Answer is:", leastInterval_optimal([]byte{'A', 'A', 'A', 'A', 'A','B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'D'}, 2)) //14
}

/**
My unique way to solve this problem. Essentially a simulation of task scheduler
But not as efficient as "right answers"
 */
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
			fmt.Print(string(h[0].id), ",") //for debugging
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

/**
Need a little bit math and deduction :)
https://leetcode.com/problems/task-scheduler/solution/

My notes：
1。找到频率最高的task，最多可能的空余cycle就是 （f（频率）- 1）X n
比如 A - - A - - A - - A - - A
2。然后找其他频率次高的task来填满这些空档 （比如task B 和 C 都频率为3）
    A - - A - - A - - A - - A
	  B - - B - - B - -
		C - - C - - C - -

3。只要剩余的task有足够的频率，总能填满这些空档。并且不会产生新的空档。(需要很多情况推理，不易证明）
	比如有另一个task D，频率为3。
	那一个D可以接在B后面， 一个D可以接在C后面，最后一个D可以插在任意之前ABC后面 （因为只要不产生空档就行）
	接在C后面那个新的D会和B后面那个D冲突，但可以容易的与前面的C调换来解决这个问题。 因为C的频率大于等于D，所以有足够的C来调换
	A B C A B C A B C A D D D A
=>  A B D A B C D A B C A D C A
      (swap) (insert)

4. as long as the idles between the most frequent tasks are filled,
it's guaranteed that the rest of tasks will not incur new idle time (bcs we can just insert them in between 'A's)

 */
func leastInterval_optimal(tasks []byte, n int) int {
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
	var tList []*Task
	for _, task := range taskMap {
		tList = append(tList, task)
	}
	sort.Slice(tList, func(i, j int) bool {
		return tList[i].count > tList[j].count
	})

	maxFrequency := tList[0].count
	idleTimes := (maxFrequency - 1) * n
	for i := 1; i < len(tList) && idleTimes > 0; i++ {
		f := 0
		if tList[i].count > maxFrequency - 1 {
			f = maxFrequency - 1
		} else {
			f = tList[i].count
		}
		idleTimes -= f
	}
	if idleTimes < 0 {
		idleTimes = 0
	}
	return len(tasks) + idleTimes
}

type Task struct {
	id byte
	count int
	coolDown int
}

type TaskHeap []*Task

func (h TaskHeap) Len() int {
	return len(h)
}

func (h TaskHeap) Less(i int, j int) bool {
	if h[i].coolDown < h[j].coolDown {
		return true
	}
	if h[i].coolDown == h[j].coolDown {
		return h[i].count > h[j].count
	}
	return false
}

func (h TaskHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TaskHeap) Push(x interface{}) {
	(*h) = append(*h, x.(*Task))
}

func (h *TaskHeap) Pop() interface{} {
	last := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return last
}