package problems

func Kdiffpairs(a *[]int, k int) int {
	c := 0
	for i, v := range *a {
		for j := i + 1; j < len(*a); j++ {
			if (*a)[j] == (*a)[j-1] {
				continue
			}
			if (*a)[j]-v == k {
				c++
			}
		}
	}
	return c
}
