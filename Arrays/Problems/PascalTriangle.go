package problems

// The triangle below is known as Pascal’s triangle.
// In this triangle, the value at a position is equal to the sum of values of the 2 elements just above it.

// Examples
// The 2nd element of 5th row is 1+3 => 4
// The 3rd element of 5th row is 3+3 => 6
// The 4th element of 5th row is 3+1 => 4
// For the leftmost and the rightmost position in each row, the value is 1. The element in the first row is also 1.

// Given a number n, find the nth row of pascal’s triangle.

func PascalTriangle(line int) []int {
	var B []int
	var val int
	line = line - 1
	B = append(B, 1)
	for i := 1; i <= line; i++ {

		val = (factorial(line)) / (factorial(line-i) * factorial(i))
		B = append(B, val)
	}
	return B
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
