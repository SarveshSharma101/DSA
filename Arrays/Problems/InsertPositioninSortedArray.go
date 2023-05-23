package problems

func InsertPositioninSortedArray(a *[]int, x int) int {
	l, m, h := 0, 0, len(*a)-1
	if (*a)[l] > x {
		return l
	}
	if x > (*a)[h] {
		return h + 1
	}
	for l <= h {
		m = (l + h) / 2
		if (*a)[m] > x {
			for m >= 0 && (*a)[m] > x {
				m--
			}
			if (*a)[m] == x {
				return m
			}
			return m + 1
		}
		l = m + 1
	}

	return 0
}
