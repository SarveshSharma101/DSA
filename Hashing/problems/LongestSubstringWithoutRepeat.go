package problems

func LongestSubstringWithoutRepeat(s string) int {
	Ls := ""

	m := make(map[string]int)

	for _, char := range s {
		if m[string(char)]+1 > 1 {
			ts := ""

			for k := range m {
				ts += k
			}
			if len(ts) > len(Ls) {
				Ls = ts
			}
			m = make(map[string]int)
			continue
		}
		m[string(char)] = m[string(char)] + 1
	}
	return len(Ls)
}
