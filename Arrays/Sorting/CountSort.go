package sorting

import "fmt"

func CountSort(arr *[]int, size int, range_ int) {
	var B []int = make([]int, range_+1)

	for _, v := range *arr {
		B[v] = B[v] + 1
	}
	fmt.Println(B)
	k := 0
	l := 0
	for i, v := range B {
		l = k
		for x := k; x < l+v; x++ {
			(*arr)[x] = i
			k++
		}
	}

}
