package problems

// Contains Element?
// Easy
// 30
// Given a sorted array and a number key, return whether the key is present in the array or not.

// Expected Time Complexity: O(log n)

// Examples
// Array: [1, 2, 3, 3, 3, 4, 4, 5]
// Number: 2
// Answer: true
// Array: [1, 2, 3, 3, 3, 4, 4, 5]
// Number: 6
// Answer: false

func ContainsElement(a *[]int, x int) bool {
	l, h := 0, len(*a)-1
	m := 0
	for l < h {
		if (*a)[l] == x {
			return true
		} else if (*a)[h] == x {
			return true
		} else {
			m = (l + h) / 2
			if (*a)[m] == x {
				return true
			} else if (*a)[m] > x {
				h = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return false
}
