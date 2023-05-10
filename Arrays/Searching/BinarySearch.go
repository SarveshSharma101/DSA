package searching

func BinarySearch(arr *[]int, x int) int {
	l := 0
	h := len(*arr) - 1
	mid := 0
	for l < h {
		if (*arr)[l] == x {
			return l
		} else if (*arr)[h] == x {
			return h
		}
		mid = (l + h) / 2
		if (*arr)[mid] == x {
			return mid
		} else if (*arr)[mid] < x {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return -1
}
