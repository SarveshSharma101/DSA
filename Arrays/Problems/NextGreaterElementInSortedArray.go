package problems

// Next Greater Element In Sorted Array
// Easy
// 30
// Given a sorted array and a number key, find the smallest array element which is greater than the key.

// If the key is greater than or equal to the largest element then return the key itself.

func NextGreaterElementInSortedArray(a *[]int, x int) int {
	l, h, m := 0, len(*a)-1, 0

	if x >= (*a)[h] {
		return x
	} else if (*a)[l] > x {
		return (*a)[l]
	} else {
		for l < h {

			m = (l + h) / 2
			if (*a)[m] > x {
				for m >= 0 && (*a)[m] > x {
					m--
				}
				return (*a)[m+1]
			}
			if (*a)[m] <= x {
				l = m + 1
			}
		}
	}
	return 0
}
