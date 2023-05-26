package problems

func MedianofRowwiseSortedMatrix(a *[][]int) int {
	max := 0
	for _, v := range *a {
		if max < v[len(v)-1] {
			max = v[len(v)-1]
		}
	}
	var count []int = make([]int, max+1)

	for _, v := range *a {
		for _, v2 := range v {
			count[v2] = count[v2] + 1

		}
	}
	var k []int = make([]int, len(*a)*len((*a)[0]))
	i := 0
	for x, v := range count {
		for v > 0 {
			k[i] = x
			v--
			i++
		}
	}
	return k[(len(k)-1)/2]
}
