package problems

func SearchRotatedSortedArray(a *[]int, x int) int {
	l, h := 0, len(*a)-1
	if (*a)[l] == x {
		return l
	} else if (*a)[h] == x {
		return h
	}

	p := findPivot(a, l, h)
	if (*a)[p] >= x {
		return binarysearch(a, l, p, x)
	}

	return binarysearch(a, p+1, h, x)
}

func findPivot(a *[]int, i, j int) int {
	if (*a)[i] <= (*a)[j] {
		return j
	}
	return findPivot(a, i, j-1)
}

func binarysearch(a *[]int, l, h, x int) int {

	m := 0
	for l <= h {
		if (*a)[l] == x {
			return l
		} else if (*a)[h] == x {
			return h
		} else {
			m = (l + h) / 2
			if (*a)[m] == x {
				return m
			} else if (*a)[m] > x {
				h = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}
