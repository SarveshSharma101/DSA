package problems

func NegativeNumbersInSortedArray(a *[]int) int {
	c := 0
	if (*a)[0] >= 0 {
		return c
	}
	i := 0
	for (*a)[i] < 0 {
		c++
		i++
	}
	return c
}
