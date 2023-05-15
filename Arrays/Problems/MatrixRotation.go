package problems

// Given a matrix, rotate the matrix 90 degrees clockwise
// Examples
// Matrix:
// 1 2 3
// 4 5 6
// 7 8 9
// After rotation:
// 7 4 1
// 8 5 2
// 9 6 3
// Matrix:
// 1 2
// 3 4
// 5 6
// After rotation:
// 5 3 1
// 6 4 2

func RotateMatrix(m *[][]int) [][]int {
	var B [][]int = make([][]int, len((*m)[0]))
	for i := range B {
		B[i] = make([]int, len(*m))
	}

	j := 0
	for i := range *m {
		j = len((*m)) - 1 - i
		for k := 0; k < len((*m)[0]); k++ {
			B[k][j] = (*m)[i][k]
		}
	}
	return B
}
