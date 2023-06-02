package problems

import "math"

// Trapping Rain Water
// Hard
// 80
// Given an array A where each element denotes a the height of blocks, calculate the total volume of water that can be trapped when it rains.

// Note: one cubic block has a volume of 1 unit.

// Example:
// A: [ 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1 ]

// The total volume of water is 1 + 3 + 1 + 1 = 6 units.

// Follow the steps mentioned below to implement the idea:

// Traverse the array from start to end:
// For every element:
// Traverse the array from start to that index and find the maximum height (a) and
// Traverse the array from the current index to the end, and find the maximum height (b).
// The amount of water that will be stored in this column is min(a,b) – array[i], add this value to the total amount of water stored
// Print the total amount of water stored.
// Below is the implementation of the above approach.

func TrappingRainWater(a *[]int) int {
	v := 0
	r := 0
	l := 0
	for i, v2 := range *a {
		l = v2
		for j := 0; j < i; j++ {
			if (*a)[j] > v2 {
				l = (*a)[j]
			}
		}

		r = v2
		for j := i + 1; j < len(*a); j++ {
			if (*a)[j] > v2 {
				r = (*a)[j]
			}
		}

		v = v + (int(math.Min(float64(l), float64(r))) - v2)
	}

	return v
}
