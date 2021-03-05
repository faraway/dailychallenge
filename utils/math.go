package utils

func MIN(num ...int) int {
	min := num[0]
	for _, x := range num {
		if x < min {
			min = x
		}
	}
	return min
}

func MAX(num ...int) int {
	max := num[0]
	for _, x := range num {
		if x > max {
			max = x
		}
	}
	return max
}