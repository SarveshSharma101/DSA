package problems

func Kthelementoftwosortedlists(a *[]int, b *[]int, k int) int {
	// C := make([]int, len(*a)+len(*b)+1)
	x := 0
	k = k - 1
	i := 0
	j := 0

	for i < len(*a) && j < len(*b) {
		if (*a)[i] <= (*b)[j] {
			if x == k {
				return (*a)[i]
			}
			i++
			x++
		} else if (*a)[i] > (*b)[j] {
			if x == k {
				return (*b)[j]
			}
			j++
			x++
		}
	}

	for i < len(*a) {
		if x == k {
			return (*a)[i]
		}
		i++
		x++

	}

	for j < len(*b) {
		if x == k {
			return (*b)[j]
		}
		j++
		x++
	}

	return -1
}
