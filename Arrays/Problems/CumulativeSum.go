package problems

// The cumulative sum of an array at index i is defined as the sum of all elements of the array from index 0 to index i.
// Examples
// Initial Array: [1, 2, 3, 4]
// Cumulative Sum: [1, 3, 6, 10]
// Initial Array: [1, 1, 1, 1, 1]
// Cumulative Sum: [1, 2, 3, 4, 5]
// Initial Array: [1, 3, 5, 7, 9]
// Cumulative Sum: [1, 4, 9, 16, 25]

func CumulativeSum(arr *[]int) []int {

	sum := 0
	B := make([]int, len(*arr))
	for i, v := range *arr {
		sum = sum + v
		B[i] = sum
	}
	return B
}

// The positive cumulative sum of an array is a list of only those cumulative sums which are positive.
// Examples
// Initial Array: [1, -2, 3, 4, -6]
// Cumulative Sum: [1, -1, 2, 6, 0]
// Positive Cumulative Sum: [1, 2, 6]
// Initial Array: [1, -1, -1, -1, 1]
// Cumulative Sum: [1, 0, -1, -2, -1]
// Positive Cumulative Sum: [1]
// Initial Array: [1, 3, 5, 7]
// Cumulative Sum: [1, 4, 9, 16]
// Positive Cumulative Sum: [1, 4, 9, 16]

func PositiveCumulativeSum(arr *[]int) []int {

	sum := 0
	k := 0
	B := make([]int, len(*arr))
	for _, v := range *arr {
		sum = sum + v
		if sum > 0 {
			B[k] = sum
			k++
		}
	}
	return B
}
