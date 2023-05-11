package sorting

func QuickSort(arr *[]int, l, h int) {
	if l < h {
		p := partition(arr, l, h)
		QuickSort(arr, l, p-1)
		QuickSort(arr, p+1, h)
	}
}

func partition(arr *[]int, l, h int) int {
	pivot := (*arr)[h]
	i := l
	j := h - 1

	for {

		for (*arr)[i] < pivot {
			i++
		}

		for (*arr)[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		swap(arr, i, j)
	}
	swap(arr, i, h)
	return i
}
