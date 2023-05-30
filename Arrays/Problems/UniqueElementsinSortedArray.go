package problems

func UniqueElementsinSortedArray(a *[]int) int {
	c := len(*a)
	j := 0

	i := 0
	for i < len(*a) {
		j = i + 1
		for j < len(*a) && (*a)[j] == (*a)[i] {
			c--
			j++
		}
		i = j
	}

	return c
}
