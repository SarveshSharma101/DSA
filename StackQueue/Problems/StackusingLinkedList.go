package problems

import "fmt"

type LLStack struct {
	StackHead *LL
	Top       int
	Size      int
}

func InitLLStack(size int) *LLStack {
	return &LLStack{
		StackHead: nil,
		Top:       -1,
		Size:      size,
	}
}

func (s *LLStack) IsLLStackEmpty() bool {
	return s.Top == -1
}

func (s *LLStack) LLStackSize() int {
	return s.Top + 1
}

func (s *LLStack) Push(x int) {
	if s.IsLLStackEmpty() {
		node := LL{
			Data: x,
		}
		s.StackHead = &node
		s.Top = s.Top + 1
		return
	}
	node := LL{
		Data: x,
	}
	temp := s.StackHead

	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &node
	s.Top++
}

func (s *LLStack) Pop() int {
	if s.IsLLStackEmpty() {
		return -100
	}

	temp := s.StackHead
	if temp.Next == nil {
		s.StackHead = nil
		s.Top--
		return -100
	}

	for temp.Next.Next != nil {
		temp = temp.Next
	}
	x := temp.Next.Data
	temp.Next = nil
	s.Top--
	return x
}

func (s *LLStack) Display() {
	if s.IsLLStackEmpty() {
		return
	}
	temp := s.StackHead

	for temp != nil {
		fmt.Println("-->", temp.Data)
		temp = temp.Next
	}
}
