package main

import (
	"container/heap"
	. "dailychallenge/utils"
	"fmt"
	"sort"
)

/**
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
find the minimum number of conference rooms required.
For example, Given [[0, 30],[5, 10],[15, 20]], return 2.

ref: https://leetcode.com/problems/meeting-rooms-ii/
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", minimumNumOfRooms([]Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	})) //2

	fmt.Println("Answer is:", minimumNumOfRooms([]Interval{
		{Start: 0, End: 10},
		{Start: 15, End: 20},
		{Start: 20, End: 30},
	})) //1

	fmt.Println("Answer is:", minimumNumOfRooms([]Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 5, End: 25},
	})) //3

	fmt.Println("Answer is:", minimumNumOfRooms([]Interval{
		{Start: 1, End: 5},
		{Start: 8, End: 9},
		{Start: 8, End: 9},
	})) //2

	fmt.Println("Answer is:", minimumNumOfRooms([]Interval{
		{Start: 1293, End: 2986},
		{Start: 848, End: 3846},
		{Start: 4284, End: 5907},
		{Start: 4466, End: 4781},
		{Start: 518, End: 2918},
		{Start: 300, End: 5870},
	})) //4


	// optimal solution using heap
	fmt.Println("Answer is:", minimumNumOfRooms_Optimal([]Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	})) //2

	fmt.Println("Answer is:", minimumNumOfRooms_Optimal([]Interval{
		{Start: 0, End: 10},
		{Start: 15, End: 20},
		{Start: 20, End: 30},
	})) //1

	fmt.Println("Answer is:", minimumNumOfRooms_Optimal([]Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 5, End: 25},
	})) //3

	fmt.Println("Answer is:", minimumNumOfRooms_Optimal([]Interval{
		{Start: 1, End: 5},
		{Start: 8, End: 9},
		{Start: 8, End: 9},
	})) //2

	fmt.Println("Answer is:", minimumNumOfRooms_Optimal([]Interval{
		{Start: 1293, End: 2986},
		{Start: 848, End: 3846},
		{Start: 4284, End: 5907},
		{Start: 4466, End: 4781},
		{Start: 518, End: 2918},
		{Start: 300, End: 5870},
	})) //4
}

type Room struct {
	MeetingList []Interval
	EarliestAvailable int
}

/**
 Less optimal solution compared to the one using heap (see below)
 But it records the meeting list and easy to understand and to print out rooms with all scheduled meetings.
 Also passes all the leet code tests for problem "meeting room II"
 */
func minimumNumOfRooms(input []Interval) int {
	//sort the input intervals by start time
	sort.Slice(input, func(i, j int) bool {
		return input[i].Start < input[j].Start
	})
	fmt.Println("sorted intervals: ", input)

	var rooms []*Room
	for _, meeting := range input {
		foundRoom := false
		for _, r := range rooms {
			if r.EarliestAvailable <= meeting.Start {
				foundRoom = true
				r.MeetingList = append(r.MeetingList, meeting)
				r.EarliestAvailable = meeting.End
				break
			}
		}
		if !foundRoom {
			rooms = append(rooms, &Room{MeetingList: []Interval{meeting}, EarliestAvailable: meeting.End})
		}
	}
	return len(rooms)
}

/**
 Optimal solution, use a heap to record the (potential) latest available room (instead of scanning entire array of rooms)
 */
func minimumNumOfRooms_Optimal(input []Interval) int {
	//sort the input intervals by start time
	sort.Slice(input, func(i, j int) bool {
		return input[i].Start < input[j].Start
	})

	minHeap := &IntHeap{}
	heap.Init(minHeap)
	for _, meeting := range input {
		//see if we can use an existing room
		if minHeap.Len() > 0 && (*minHeap)[0] <= meeting.Start {
			heap.Pop(minHeap)
		}
		/**
			either way, we add this latest meeting's end time to the heap.
			if we can use an existing room (as checked above),
			then we maintain the same room number, but with End time updated
		*/
		heap.Push(minHeap, meeting.End)
	}
	//number of room required is the size of the heap
	return minHeap.Len()
}