package problems

// Inversion Count
// Medium
// 50
// The inversion count of an array denotes how far is the array from being sorted.

// If the array is sorted, inversion count is 0. If the array is sorted in reverse order, the inversion count is maximum.

// More formally, the inversion count of an array A is the number of pairs (i, j) such A[i] < A[j] and i > j.

// Example
// Lets take the following array:
// 8, 4, 1, 2

// This array has an inversion count of 5.
// (8, 4), (8, 1), (8, 2), (4, 1), (4, 2)

// Given an array A, calculate the inversion count of the array.

func InversionCount(a *[]int) int {
	count := 0
	for i := 0; i < len(*a)-1; i++ {
		for j := i + 1; j < len(*a); j++ {
			if (*a)[i] > (*a)[j] {
				count++
			}
		}
	}
	return count
}
