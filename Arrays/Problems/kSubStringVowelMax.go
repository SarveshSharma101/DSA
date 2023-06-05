package problems

func kSubstringVowelsMax(s string, k int) int {
	c := 0
	a := []byte(s)
	max := 0
	for i := 0; i <= len(a)-k; i++ {
		c = 0
		for j := i; j < i+k; j++ {
			switch string(a[j]) {
			case "a":
				c++
			case "e":
				c++
			case "i":
				c++
			case "o":
				c++
			case "u":
				c++
			}
		}
		if c > max {
			max = c
		}
	}

	return max
}
