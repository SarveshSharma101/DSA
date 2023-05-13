package problems

// Given an array A, find the maximum number of consecutive 1s in the array.

// Examples
// A: [1, 1, 3, 2, 3, 1, 1, 1]
// Max consecutive 1s: 3

func MaxConsecutiveOnes(a *[]int) int {
	maxOnes := 0
	x := 0
	for i := 0; i < len(*a); i = x {
		if (*a)[i] == 1 {
			n := 0
			for j := i; j < len(*a); j++ {
				if (*a)[j] != 1 {
					x = j
					break
				}
				n++
			}
			if n > maxOnes {
				maxOnes = n
			}
		}
		x = x + 1
	}
	return maxOnes
}
