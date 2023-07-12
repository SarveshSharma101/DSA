package problems

func LongestConsecutiveSequence(a *[]int) []int {

	C := make([]int, 1000)

	for _, v := range *a {
		C[v] = 1
	}

	var B []int
	i := 0
	for i < len(C) {
		var c []int
		if C[i] == 1 {
			for C[i] == 1 {
				c = append(c, i)
				i++
			}
			if len(c) > len(B) {
				B = c
			}
			continue
		}

		i++
	}
	return B
}
