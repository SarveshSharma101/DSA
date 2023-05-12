package sorting

func MergeSort(arr *[]int, l, h int) {
	if l < h {
		m := (l + h) / 2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, h)
		Merge(arr, l, m, h)

	}
}

func Merge(arr *[]int, l, mid, h int) {
	i := l
	j := mid + 1
	k := l
	var B []int = make([]int, h+100)

	for i <= mid && j <= h {
		if (*arr)[i] < (*arr)[j] {
			B[k] = (*arr)[i]
			k++
			i++
		} else {
			B[k] = (*arr)[j]
			k++
			j++
		}
	}

	for i <= mid {
		B[k] = (*arr)[i]
		k++
		i++
	}

	for j <= h {
		B[k] = (*arr)[j]
		k++
		j++
	}

	for i := l; i <= h; i++ {
		(*arr)[i] = B[i]
	}
}
