package main

import "fmt"

/**
https://leetcode.com/problems/robot-bounded-in-circle/

On an infinite plane, a robot initially stands at (0, 0) and faces north. The robot can receive one of three instructions:

"G": go straight 1 unit;
"L": turn 90 degrees to the left;
"R": turn 90 degrees to the right.
The robot performs the instructions given in order, and repeats them forever.

Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.



Example 1:
Input: instructions = "GGLLGG"
Output: true
Explanation: The robot moves from (0,0) to (0,2), turns 180 degrees, and then returns to (0,0).
When repeating these instructions, the robot remains in the circle of radius 2 centered at the origin.

Example 2:
Input: instructions = "GG"
Output: false
Explanation: The robot moves north indefinitely.

Example 3:
Input: instructions = "GL"
Output: true
Explanation: The robot moves from (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ...
*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", isRobotBounded("GGLLGG")) //True
	fmt.Println("Answer is:", isRobotBounded("GG")) //False
	fmt.Println("Answer is:", isRobotBounded("GL")) //True
}

func isRobotBounded(instructions string) bool {
	var currentDirection uint8 = 'N'
	x, y := 0, 0
	round := 0
	//If the robot will ever return to the starting point, it will be at most 4 cycles (heuristics: we have 4 directions)
	//proof: https://leetcode.com/problems/robot-bounded-in-circle/solution/#appendix-mathematical-proof
	for round < 4 {
		for i:=0; i<len(instructions); i++ {
			command := instructions[i]
			if command == 'G' {
				switch currentDirection {
				case 'N': y++
				case 'S': y--
				case 'E': x++
				case 'W': x--
				}
			} else {
				currentDirection = turn(currentDirection, command)
			}
		}
		//we are at the original point. doesn't matter which direction we are facing...
		if x==0 && y==0 {
			return true
		}
		round ++
	}
	return false
}

//we can use numbers to simplify this
//(e.g. 1,2,3,4 to indicate directions. 'left' is (num-1)%4 while 'right is (num+1)%4 )
func turn(currentDirection uint8, command uint8) uint8 {
	switch currentDirection {
	case 'N':
		switch command {
		case 'L':
			return 'W'
		case 'R':
			return 'E'
		}
	case 'W':
		switch command {
		case 'L':
			return 'S'
		case 'R':
			return 'N'
		}
	case 'S':
		switch command {
		case 'L':
			return 'E'
		case 'R':
			return 'W'
		}
	case 'E':
		switch command {
		case 'L':
			return 'N'
		case 'R':
			return 'S'
		}
	}
	//we should not be here
	return 'N'
}
