package problems

func DistinctNumbersinWindow(a *[]int, k int) []int {
	var C []int

	for i := 0; i <= len(*a)-k; i++ {
		m := make(map[int]int)
		j := i
		c := k
		for c > 0 {
			m[(*a)[j]] = m[(*a)[j]] + 1
			c--
			j++
		}
		C = append(C, len(m))
	}

	return C
}
