package problems

func IdenticalTwins(a *[]int) int {
	var C map[int]int = map[int]int{}

	for _, v := range *a {

		val := C[v]
		if v == 0 {
			C[v] = 1
		} else {
			C[v] = val + 1
		}
	}

	it := 0

	for _, v := range C {
		it += (v * (v - 1)) / 2
	}

	return it
}
