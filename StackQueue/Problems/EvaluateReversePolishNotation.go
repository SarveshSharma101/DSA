package problems

import (
	"strconv"
)

func EvaluateReversePolishNotation(arr []string) int {
	sum := 0
	s := InitStack(len(arr))
	for _, v := range arr {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			b := s.Pop()
			a := s.Pop()
			sum = eval(a, b, v)
			s.Push(sum)
			sum = 0
			continue
		}
		A, _ := strconv.Atoi(v)
		s.Push(A)
	}

	return s.Pop()
}

func eval(A, B int, o string) int {
	var e int

	switch o {
	case "+":
		e = A + B
	case "-":
		e = A - B
	case "*":
		e = A * B
	case "/":
		e = A / B
	}

	return e
}
