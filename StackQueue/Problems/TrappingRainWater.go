package problems

import "math"

func TrappingRainWater(tpw *[]int) int {
	w := 0

	s := InitStack(100)
	i := 0
	for i < len(*tpw) {

		for !s.IsEmpty() && (*tpw)[i] > (*tpw)[s.StackPeek()] {
			t := s.Pop()
			if s.IsEmpty() {
				break
			}
			d := i - s.StackPeek() - 1
			h := int(math.Min(float64((*tpw)[i]), float64((*tpw)[s.StackPeek()]))) - (*tpw)[t]
			w += d * h
		}

		s.Push(i)
		i++
	}

	return w
}
