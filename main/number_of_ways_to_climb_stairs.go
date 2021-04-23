package main

import "fmt"

/**
You are given a positive integer N which represents the number of steps in a staircase.
You can either climb 1 or 2 steps at a time.
Write a function that returns the number of unique ways to climb the stairs.
 */
func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", stairs(1)) //1
	fmt.Println("Answer is:", stairs(2)) //2
	fmt.Println("Answer is:", stairs(5)) //8
	fmt.Println("Answer is:", stairs(10)) //89
}


func stairs(n int) int {
	//actually you may not need an array, just two iterating number count[i-2] and count [i-1]
	array := make([]int, n)
	for i,_ := range array {
		if i==0 {
			array[i] = 1
		} else if i==1 {
			array[i] = 2
		} else {
			array[i] = array[i-2] + array[i-1]
		}
	}
	return array[n-1]
}