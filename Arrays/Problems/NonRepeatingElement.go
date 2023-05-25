package problems

// Non-Repeating Element
// Medium
// 50
// Given a sorted list of numbers in which all elements appear twice except one element that appears only once, find the number that appears only once.

// Example:
// 1, 1, 2, 3, 3, 4, 4
// Here, ‘2’ appears once and all other elements appear twice.

// findNonRepeatingElement([1, 1, 2, 3, 3, 4, 4]) => 2
// Expected Time Complexity: O(log n)

func NonRepeatingElement(a *[]int) int {
	i := 0

	for i < len(*a) {
		if i == len(*a)-1 {
			return (*a)[i]
		}
		if (*a)[i] == (*a)[i+1] {
			i = i + 2
			continue
		}
		return (*a)[i]
	}
	return -1
}
