package problems

import "fmt"

type Stack struct {
	S    []int
	Top  int
	Size int
}

func InitStack(size int) *Stack {
	return &Stack{
		S:    make([]int, size),
		Top:  -1,
		Size: size,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) StackSize() int {
	return s.Top + 1
}

func (s *Stack) StackTop() int {
	return s.Top
}

func (s *Stack) StackPeek() int {
	return s.S[s.Top]
}

func (s *Stack) IsFull() bool {
	return s.Top == s.Size-1
}

func (s *Stack) Push(x int) {
	if s.IsFull() {
		return
	}
	s.Top = s.Top + 1
	s.S[s.Top] = x
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -100
	}
	x := s.S[s.Top]
	s.S[s.Top] = 0
	s.Top = s.Top - 1
	return x
}

func (s *Stack) Display() {
	if s.IsEmpty() {
		return
	}
	for i := 0; i <= s.Top; i++ {
		fmt.Println("--->", s.S[i])
	}
}
