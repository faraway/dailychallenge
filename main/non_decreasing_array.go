package main

import "fmt"
/**
ou are given an array of integers in an arbitrary order. Return whether or not it is possible to make the array non-decreasing by modifying at most 1 element to any value.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

Example:

[13, 4, 7] should return true, since we can modify 13 to any value 4 or less, to make it non-decreasing.

[13, 4, 1] however, should return false, since there is no way to modify just one element to make the array non-decreasing.

 */
func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", check([]int{13,4,7})) //true
	fmt.Println("Answer is:", check([]int{13,4,1})) //false
	fmt.Println("Answer is:", check([]int{5,10,3,7})) //false
}

func check(input []int) bool {
	idx := -1
	idxOriginalVal := 0
	// check to see if the array has decreasing value, if there is, find the first occurrence
	for index, _ := range input {
		if index > 0 && input[index]<input[index-1] {
			idx = index
			idxOriginalVal = input[index]
			break
		}
	}
	if idx == -1 { //non-decreasing, simply return true
		return true
	} else {
		//first decreasing happens, now let's try change one number to still make it a non-decreasing array
		//there are only two options to change one number to get us back on track

		//    index-2  bigger(index-1)  smaller(index)
		//                   =
		//      =            =
		//      =            =                 =              ...
		//      =            =                 =              ...
		//      =            =                 =              ...
		//-----------------------------------------------------------

		//option1: make the smaller number as big as the previous one
		input[idx] = input[idx-1]
		if isNonDecreasing(input) { return true }

		//option2: make the previous bigger value as small as current one
		input[idx-1] = idxOriginalVal
		input[idx] = idxOriginalVal
		if isNonDecreasing(input) { return true }
	}

	//no fix possible
	return false
}

func isNonDecreasing(array []int) bool {
	for index, _ := range array {
		if index > 0 && array[index]<array[index-1] {
			return false
		}
	}
	return true
}