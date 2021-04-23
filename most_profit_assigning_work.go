package main

import (
	"fmt"
	"sort"
)

/**

https://leetcode.com/problems/most-profit-assigning-work/

We have jobs: difficulty[i] is the difficulty of the ith job, and profit[i] is the profit of the ith job.

Now we have some workers. worker[i] is the ability of the ith worker,
which means that this worker can only complete a job with difficulty at most worker[i].

Every worker can be assigned at most one job, but one job can be completed multiple times.

For example, if 3 people attempt the same job that pays $1, then the total profit will be $3.
If a worker cannot complete any job, his profit is $0.

What is the most profit we can make?

Example 1:
Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
Output: 100
Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get profit of [20,20,30,30] seperately.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", maxProfitAssignment([]int{2,4,6,8,10}, []int{10,20,30,40,50}, []int{4,5,6,7})) //100
	fmt.Println("Answer is:", maxProfitAssignment([]int{85,47,57}, []int{24,66,99}, []int{40,25,25})) //0

	// NOTE: IT'S NOT the higher the difficulty, the higher the profit !!! Easier job *may* have high profit!!
	fmt.Println("Answer is:", maxProfitAssignment([]int{68,35,52,47,86}, []int{67,17,1,81,3}, []int{92,10,85,84,82})) //324

}


func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var works = make([]Work, len(profit))
	for i:=0; i<len(profit); i++ {
		works[i] = Work{difficulty[i], profit[i]}
	}

	sort.Slice(works, func(i, j int) bool {
		return works[i].Difficulty < works[j].Difficulty
	})

	//sort workers as well, see why it's useful below
	sort.Ints(worker)

	maxProfit := 0
	workIdx, best := 0, 0
	for _, capability := range worker {
		//finds the best profit within the current capability; NOTE workInx only increases and not get reset every iteration
		for workIdx < len(works) && capability >= works[workIdx].Difficulty {
			if works[workIdx].Profit > best {
				best = works[workIdx].Profit
			}
			workIdx++
		}
		maxProfit += best
	}
	return maxProfit
}

type Work struct {
	Difficulty int
	Profit int
}

/**
 Binary search for the highest capable job you can find.
 */
func searchForMaxCapableWork(works []Work, capability int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := (start+end)/2
	if works[mid].Difficulty == capability {
		return mid //assuming same difficulty = same profit, we stop searching here
	} else if works[mid].Difficulty > capability {
		return searchForMaxCapableWork(works, capability, mid+1, end)
	} else {
		result := searchForMaxCapableWork(works, capability, 0, mid-1)
		//take this job if there is no better job
		if result == -1 {
			result = mid
		}
		return result
	}
}