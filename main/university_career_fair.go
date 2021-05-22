package main

import (
	"fmt"
	"sort"
)

/**
Hacker Rank problem - TESTED

A team organizing a university career fair has a list of companies along with their respective arrival times and
their duration of stay. Only one company can present at any time.
Given each company's arrival time and the duration they will stay,
determine the maximum number of presentations that can be hosted during the career fair.


Example
n = 5
arrival = [1, 3, 3, 5, 7]
duration = [2, 2, 1, 2, 1]

Explanation
The first company arrives at time 1 and stays for 2 hours.
At time 3, two companies arrive, but only 1 can stay for either 1 or 2 hours.
The next companies arrive at times 5 and 7 and do not conflict with any others.
In total, there can be a maximum of 4 promotional events.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", maxEvents([]int32{1,3,3,5,7},[]int32{2,2,1,2,1})) //4
}

func maxEvents(arrival []int32, duration []int32) int32 {
	var segments []Segment = make([]Segment, len(arrival))
	for i := range arrival {
		segments[i] = Segment{start: arrival[i], end: arrival[i]+duration[i]}
	}

	sort.Slice(segments, func(i int, j int) bool {
	    return segments[i].end < segments[i].end
	})

	var result int32 = 0
	var currentTime int32 = 0
	for _, event := range segments {
		if event.start >= currentTime {
			result ++
			currentTime = event.end
		}
	}
	return result
}


type Segment struct {
	start int32
	end int32
}



//sort lib not available, a simple bubble sort
func sortSegments(segments []Segment){
	n := len(segments)
	for i:=0; i<n-1; i++ {
		for j:=0; j<n-i-1; j++ {
			if segments[j].end > segments[j+1].end {
				segments[j], segments[j+1] = segments[j+1], segments[j]
			}
		}
	}
}