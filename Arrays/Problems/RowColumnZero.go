package problems

// Given a matrix, if any of the cells in the matrix is 0, set all the elements in its row and column to 0.

// Examples
// 1 1 1        1 0 1
// 1 0 1   =>   0 0 0
// 1 1 1        1 0 1

// 0 1 2        0 0 0
// 3 4 5   =>   0 4 5
// 1 2 3        0 2 3

// 0 1 0        0 0 0
// 3 4 5   =>   0 4 0
// 1 2 3        0 2 0

// 0 1 0        0 0 0
// 3 0 5   =>   0 0 0
// 1 2 3        0 0 0

func RowColumnZero(m *[][]int) {
	B := getZeroIndex(m)
	i, j := 0, 0
	for _, v := range B {
		i, j = v[0], v[1]
		for j < len((*m)[0]) {
			(*m)[v[0]][j] = 0
			j++
		}
		j = v[1]
		for j >= 0 {
			(*m)[v[0]][j] = 0
			j--
		}

		for i < len(*m) {
			(*m)[i][v[1]] = 0
			i++
		}
		i = v[0]
		for i >= 0 {
			(*m)[i][v[1]] = 0
			i--
		}
	}
}

func getZeroIndex(m *[][]int) [][]int {
	var B [][]int
	for i, a := range *m {
		for j, v := range a {
			if v == 0 {
				B = append(B, []int{i, j})
			}
		}
	}
	return B
}
