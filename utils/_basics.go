package utils

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func testFun() {
	//copy array
	var destination [5]int
	var source = []int{1,2,3,4,5}
	copy(destination, source)

	//sort
	slice := []string{"john", "yuan", "yao"}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	//strings
	strArray := strings.Split("I:love:you", ":")
	words := strings.Fields("In my hometown, there is a dog")

	//type conversion
	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.Itoa(123))

	//other type conversion
	var number int32 = 42
	fmt.Println(float64(number))

	//max numbers
	fmt.Println(math.MaxInt32, math.MinInt32)

	//initialize a 4x3 matrix
	m, n := 4, 3
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	//slices
	array := make([]int, 0, 10) //0 is length, 10 is underlying array initial capacity
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array[0:2])//slicing

	//delete map elements
	s := map[int]bool{5: true, 2: true}
	_, ok := s[6] // check for existence
	delete(s, 2)

}
