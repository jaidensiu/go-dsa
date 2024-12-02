package array

import "math"

/*
AddTwoNumbers converts the two int arrays into integers, computes their sum, and
converts the sum back into an array.

time = O(n), space = O(n)
*/
func AddTwoNumbers(num1, num2 []int) []int {
	return intToArray(arrayToInt(num1) + arrayToInt(num2))
}

func arrayToInt(arr []int) int {
	result := 0

	for i, v := range arr {
		result += v * int(math.Pow10(len(arr)-i-1))
	}

	return result
}

func intToArray(num int) []int {
	if num == 0 {
		return []int{0}
	}

	var result []int

	for num > 0 {
		digit := num % 10
		result = append([]int{digit}, result...)
		num /= 10
	}

	return result
}
