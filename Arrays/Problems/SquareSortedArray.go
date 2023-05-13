package problems

func SquraeSortedArray(a *[]int) []int {
	square := 0
	B := make([]int, len(*a))
	for i, v := range *a {
		square = v * v
		(B)[i] = square
	}
	MergeSort(&B, 0, len(B)-1)
	return B
}
