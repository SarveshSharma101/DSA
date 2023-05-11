package sorting

func BubbleSort(arr *[]int) {
	for i := range *arr {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func swap(arr *[]int, j, k int) {
	temp := (*arr)[j]
	(*arr)[j] = (*arr)[k]
	(*arr)[k] = temp
}
