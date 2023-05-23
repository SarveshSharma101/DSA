package problems

func SearchRange(a *[]int, x int) []int {
	l := 0
	h := len(*a) - 1

	for (*a)[l] < x && l < len(*a)-1 {
		l++
	}

	for l < len(*a) && (*a)[h] > x && h > -1 {
		h--
	}

	if l > len(*a)-1 || h < 0 || ((*a)[l] != x && (*a)[h] != x) {
		return []int{-1, -1}
	}
	return []int{l, h}
}
