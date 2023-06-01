package problems

func SortedArraysIntersection(a, b *[]int) []int {
	var C [][]int
	var K []int
	var L []int

	j := 0
	// j := 0
	if len(*a) >= len(*b) {
		K = *a
		L = *b
	} else {
		K = *b
		L = *a
	}

	for i, v := range K {

		j = binarysearch_(&L, 0, len(L)-1, v)
		if j != -1 {
			C = append(C, getIntersection(&K, &L, i, j))
		}
	}
	var m int = 0
	var max []int
	for _, v := range C {
		if len(v) >= m {
			max = v
			m = len(v)
		}
	}

	return max
}

func getIntersection(a, b *[]int, i, j int) []int {
	var C []int

	for (i < len(*a) && j < len(*b)) && (*a)[i] == (*b)[j] {
		C = append(C, (*a)[i])
		i++
		j++
	}

	return C
}

func binarysearch_(a *[]int, l, h, x int) int {

	m := 0
	for l <= h {
		// if (*a)[l] == x {
		// 	return l
		// } else if (*a)[h] == x {
		// 	return h
		// } else {
		m = (l + h) / 2
		if (*a)[m] == x {
			return m
		} else if (*a)[m] > x {
			h = m - 1
		} else {
			l = m + 1
		}
		// }
	}

	return -1
}
