package problems

func MergeSort(arr *[]int, l, h int) {
	if l < h {
		m := (l + h) / 2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, h)
		merge(arr, l, m, h)
	}
}

func merge(a *[]int, l, m, h int) {
	var B []int
	i := l
	j := m + 1

	for i <= m && j <= h {
		if (*a)[i] < (*a)[j] {
			B = append(B, (*a)[i])
			i++
		} else {
			B = append(B, (*a)[j])
			j++
		}
	}

	for i <= m {
		B = append(B, (*a)[i])
		i++
	}

	for j <= h {
		B = append(B, (*a)[j])
		j++
	}
	k := 0
	for i := l; i <= h; i++ {
		(*a)[i] = B[k]
		k++
	}

}
