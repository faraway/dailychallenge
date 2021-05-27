package main

import "fmt"

/**
https://leetcode.com/problems/find-the-celebrity/


Suppose you are at a party with n people (labeled from 0 to n - 1), and among them, there may exist one celebrity.
The definition of a celebrity is that all the other n - 1 people know him/her, but he/she does not know any of them.

Now you want to find out who the celebrity is or verify that there is not one.
The only thing you are allowed to do is to ask questions like: "Hi, A. Do you know B?" to get information about whether A knows B.
You need to find out the celebrity (or verify there is not one) by asking as few questions as possible (in the asymptotic sense).

You are given a helper function bool knows(a, b) which tells you whether A knows B. Implement a function int findCelebrity(n).
There will be exactly one celebrity if he/she is in the party. Return the celebrity's label if there is a celebrity in the party.
If there is no celebrity, return -1.



Example 1:
Input: graph = [[1,1,0],[0,1,0],[1,1,1]]
Output: 1
Explanation: There are three persons labeled with 0, 1 and 2. graph[i][j] = 1 means person i knows person j,
otherwise graph[i][j] = 0 means person i does not know person j.
The celebrity is the person labeled as 1 because both 0 and 2 know him but 1 does not know anybody.

Example 2:
Input: graph = [[1,0,1],[1,1,0],[0,1,1]]
Output: -1
Explanation: There is no celebrity.


Constraints:
n == graph.length
n == graph[i].length
2 <= n <= 100
graph[i][j] is 0 or 1.
graph[i][i] == 1


Follow up: If the maximum number of allowed calls to the API knows is 3 * n,
could you find a solution without exceeding the maximum number of calls?

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", findTheCelebrity(func(a int, b int) bool{
		matrix := [][]int {
			{1,1,0},
			{0,1,0},
			{1,1,1},
		}
		return matrix[a][b] == 1
	})(3)) //1
}

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func findTheCelebrity(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		/**
			find the candidate, this is the key.
			if A knows B; then by definition, A can't be the celebrity
			if A doesn't B; then by definition, B can't be the celebrity
			with the winner from {A,B}, let's compare him/her with C. and so on so forth

			therefore, just by running through the list once, we should be able to narrow down to just one candidate
		 */
		candidate := 0
		for i := 1; i < n; i++ {
			if knows(candidate, i) {
				candidate = i
			}
		}
		//now we have the candidate, validate it
		//everyone should know the candidate
		for i := 0; i < n; i++ {
			if !knows(i, candidate) {
				return -1
			}
		}
		//candidate should know no one except him/herself
		for i := 0; i < n; i++ {
			if candidate != i && knows(candidate, i) {
				return -1
			}
		}
		return candidate
	}
}

