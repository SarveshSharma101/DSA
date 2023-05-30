package problems

// Remove occurences
// Easy
// 30
// Given an array and a number k, remove all occurrences of k from the array (in-place). Return the number of elements 'remainingSize' left after removing k. Note that removing the occurences is mandatory and will be checked in the main method. There can be anything beyond the first 'remainingSize' elements. It will be ignored.

// Example
// Array: [1, 4, 2, 6, 2, 6, 9, 4]
// Number: 4
// Answer: 6
// Explanation:[1, 2, 6, 2, 6, 9]

func RemoveOccurences(a *[]int, x int) int {

	s := len(*a)
	i := 0
	j := len(*a) - 1

	for i <= j {
		if (*a)[i] == x {
			s--
			(*a)[i] = -1
		}
		if (*a)[j] == x {
			s--
			(*a)[j] = -1
		}
		i++
		j--
	}
	return s
}
