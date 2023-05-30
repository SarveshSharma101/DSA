package problems

// Two Sum Sorted
// Easy
// 30
// Given a sorted array, check if there exist two numbers whose sum is zero.

// Example
// A: [-3, 1, 3, 4]
// Answer: true
// A: [-2, 1, 3, 4]
// Answer: false

func TwoSumSorted(a *[]int) bool {
	for i, v := range *a {
		for j := i + 1; j < len(*a); j++ {
			if v+(*a)[j] == 0 {
				return true
			}
		}
	}
	return false
}
