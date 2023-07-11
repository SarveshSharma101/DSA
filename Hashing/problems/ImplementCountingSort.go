package problems

import "fmt"

// Observe that we want the sum of a number of possible pairs of an element of the same type. This hints at the use of hashing.

// We can hash the frequency of all the numbers in the list using a hashmap and for each distinct number,
// we can form Nc2 = (N * (N - 1) / 2) pairs of that particular number, where N is the frequency of that number in the list.

func ImplementCountingSort(a *[]int) []int {
	C := []int{}
	D := make([]int, 10000)
	fmt.Println("--->", C)
	for _, v := range *a {
		D[v] = D[v] + 1
	}

	for i, v := range D {
		for v > 0 {
			C = append(C, i)
			v--
		}
	}

	return C
}
