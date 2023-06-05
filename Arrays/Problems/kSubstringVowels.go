package problems

func kSubstringVowels(s string, k int) []int {
	var C []int
	c := 0
	a := []byte(s)

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
		C = append(C, c)
	}

	return C
}
