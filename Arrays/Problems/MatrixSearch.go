package problems

// Matrix Search
// Medium
// 50
// Given a matrix, check if the matrix contains a given key.

// The matrix has the following properties:

// Integer in each row is arranged in non-decreasing order from left to right.
// The first integer in every row is greater than the last integer of the previous row.

// Examples
// Matrix:
// 1 2 3
// 4 5 6
// 7 8 9
// Key: 6
// Answer: true
// Matrix:
// 1 2 3
// 4 5 6
// 7 8 9
// 10 11 12
// Key: 15
// Answer: false

func MatrixSearch(a *[][]int, x int) bool {

	for _, v := range *a {
		if v[0] == x {
			return true
		} else if v[len(v)-1] == x {
			return true
		} else if v[0] < x && v[len(v)-1] > x {
			binarysearch(&v, 0, len(v)-1, x)
		}
	}

	return false
}
