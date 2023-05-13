package problems

import "fmt"

func ArithmeticSequence(a []int) bool {
	MergeSort(&a, 0, len(a)-1)
	fmt.Println(a)
	dif := a[0] - a[1]
	for i := 0; i <= len(a)-2; i++ {
		if a[i]-a[i+1] != dif {
			return false
		}
	}
	return true
}
