package problems

func SquareRoot(x int) int {

	if x < 4 {
		return x
	}

	l := 1
	h := x / 2
	m := 0
	r := 0
	for l <= h {
		m = (l + h) / 2
		sqr := m * m
		if sqr == x {
			return m
		} else if sqr < x {
			l = m + 1
			r = m
		} else {
			h = m - 1
		}
	}

	return r
}
