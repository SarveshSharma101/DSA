package problems

func DutchNationalFlag(a *[]int) []int {
	C := make([]int, 3)
	var D []int
	for _, v := range *a {
		C[v] = C[v] + 1
	}

	for i, v := range C {
		for v > 0 {
			D = append(D, i)
			v--
		}
	}

	return D
}
