package problems

// A subarray is a part of an array including one or more contiguous/adjacent elements.

// Example
// Array: [1, 2, 3, 4, 5]
// Subarrays:
// [1]
// [2]
// [3]
// [4]
// [5]
// [1, 2]
// [2, 3]
// [3, 4]
// [4, 5]
// [1, 2, 3]
// [2, 3, 4]
// [3, 4, 5]
// [1, 2, 3, 4]
// [2, 3, 4, 5]
// [1, 2, 3, 4, 5]
// If we find the sum of the elements of any subarray then that sum will be known as a contiguous sum.

func LargestContiguousSum(a *[]int) int {
	max := 0
	sum := 0
	for i := range *a {
		if (*a)[i] > max {
			max = (*a)[i]
		}
		sum = (*a)[i]
		for j := i + 1; j < len(*a); j++ {
			sum = sum + (*a)[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
