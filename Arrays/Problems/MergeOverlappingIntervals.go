package problems

func MergeOverlappingIntervals(m [][]int) [][]int {

	for i := 0; i < len(m)-1; i++ {
		if m[i][0] == 0 && m[i][1] == 0 {
			continue
		}
		for j := i + 1; j < len(m); j++ {
			if m[i][0] == m[j][0] || m[i][0] == m[j][1] || m[i][1] == m[j][0] || m[i][1] == m[j][1] {
				m[i] = []int{m[i][0], m[j][1]}
				m[j] = []int{0, 0}
			}
		}
	}
	return m
}
