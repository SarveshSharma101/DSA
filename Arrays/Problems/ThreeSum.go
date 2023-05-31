package problems

import (
	"fmt"
)

func ThreeSum(a *[]int) [][]int {
	var C [][]int
	bubbleSort(a)
	fmt.Println("Sroted array: ", a)
	sum := 0
	index := 0
	for i := 0; i < len(*a)-2; i++ {
		sum = 0 - (*a)[i]
		for j := i + 1; j < len(*a); j++ {
			index = binarySearch(a, j+1, len(*a)-1, (*a)[j], sum)
			if index != -1 {
				C = append(C, []int{(*a)[i], (*a)[j], (*a)[index]})
			}
		}

	}
	return C
}

func bubbleSort(a *[]int) {
	for i := 0; i < len(*a); i++ {
		for j := i + 1; j < len(*a); j++ {
			if (*a)[i] > (*a)[j] {
				swap(a, i, j)
			}
		}
	}
}

func binarySearch(a *[]int, l, h, x, sum int) int {
	mid := 0

	for l <= h {
		if (*a)[l]+x == sum {
			return l
		} else if (*a)[h]+x == sum {
			return h
		} else {
			mid = (l + h) / 2
			if (*a)[mid]+x == sum {
				return mid
			}
			l++
		}
	}

	return -1
}
