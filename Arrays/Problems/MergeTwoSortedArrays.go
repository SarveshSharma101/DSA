package problems

// Given two sorted arrays A and B, find the merged sorted array C by merging A and B.

// Example:
// A: [1, 2, 3, 4, 4]
// B: [2, 4, 5, 5]
// C: [1, 2, 2, 3, 4, 4, 4, 5, 5]

func TwoSortedArrays(a *[]int, b *[]int) []int {
	var C []int

	i, j := 0, 0

	for i < len(*a) && j < len(*b) {
		if (*a)[i] <= (*b)[j] {
			C = append(C, (*a)[i])
			i++
		} else {
			C = append(C, (*b)[j])
			j++
		}
	}
	for i < len(*a) {
		C = append(C, (*a)[i])
		i++
	}
	for j < len(*b) {
		C = append(C, (*b)[j])
		j++
	}
	return C
}
