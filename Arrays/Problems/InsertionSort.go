package problems

func InsertionSort(arr *[]int) {
	k := 0
	j := 0
	for i, v := range *arr {
		k = v
		j = i - 1
		for j >= 0 && (*arr)[j] > k {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = k
	}
}
