package problems

// Given an array of integers, find the elements which have an even number of digits.

// Example
// Array: [42, 564, 5775, 34, 123, 454, 1, 5, 45, 3556, 23442]
// Answer: 42, 5775, 34, 45, 3556
// The order of the returned elements should be the same as the order of the initial array

func EvenNumberOfDigits(arr *[]int) []int {
	var B []int
	for _, v := range *arr {
		if isEvenDigit(v) {
			B = append(B, v)
		}
	}
	return B
}

func isEvenDigit(num int) bool {
	n := 0
	x := num
	for num > 0 {
		num = num % 10
		x = x / 10
		if x == 0 {
			n++
			break
		}
		n++
	}

	if n != 0 && n%2 == 0 {
		return true
	}

	return false
}
