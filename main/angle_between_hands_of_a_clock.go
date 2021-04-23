package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.com/problems/angle-between-hands-of-a-clock/
Given two numbers, hour and minutes. Return the smaller angle (in degrees) formed between
the hour and the minute hand.

Example 1:
Input: hour = 12, minutes = 30
Output: 165

Example 2:
Input: hour = 3, minutes = 30
Output: 75

Example 3:
Input: hour = 3, minutes = 15
Output: 7.5

Example 4:
Input: hour = 4, minutes = 50
Output: 155

Example 5:
Input: hour = 12, minutes = 0
Output: 0
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", angleClock(12, 30)) //165.0
	fmt.Println("Answer is:", angleClock(3, 30)) //75.0
	fmt.Println("Answer is:", angleClock(3, 15)) //7.5
	fmt.Println("Answer is:", angleClock(4, 50)) //155
	fmt.Println("Answer is:", angleClock(12, 0)) //0
}

func angleClock(hour int, minutes int) float64 {
	minutesPercentage := float64(minutes) / 60.0
	minutesAngle := 360.0 * minutesPercentage
	hoursAngle := (float64(hour) + minutesPercentage) * 30.0

	diff := math.Abs(minutesAngle-hoursAngle)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}