package problems

func KSubarraySum(a *[]int, k int) []int {
	var C []int
	j := 0
	s := 0
	for i := 0; i <= len(*a)-k; i++ {
		j = i
		s = 0
		for j < (i + k) {
			s = s + (*a)[j]
			j++
		}
		C = append(C, s)
	}

	return C
}
