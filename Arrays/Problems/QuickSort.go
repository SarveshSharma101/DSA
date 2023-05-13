package problems

func QuickSort(a *[]int, l, h int) {
	if l < h {
		p := Partition(a, l, h)
		QuickSort(a, l, p-1)
		QuickSort(a, p+1, h)
	}
}

func Partition(a *[]int, l, h int) int {
	i := l
	j := h - 1
	pivot := (*a)[h]

	for {

		for (*a)[i] < pivot {
			i++
		}
		for j >= 0 && (*a)[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(a, i, j)
	}
	swap(a, i, h)
	return i
}

func swap(a *[]int, i, j int) {
	temp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = temp
}
