package sorting

func InsertionSort(arr *[]int) {
	var k int
	for i, v := range *arr {
		k = v
		j := i - 1
		for j >= 0 && (*arr)[j] > k {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = k
	}
}
