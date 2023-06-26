package problems

import "fmt"

// Push(x): Insert x at the top of the stack

// If the stack is empty, insert x into the stack and make minEle equal to x.
// If the stack is not empty, compare x with minEle. Two cases arise:
// If x is greater than or equal to minEle, simply insert x.
// If x is less than minEle, insert (2*x – minEle) into the stack and make minEle equal to x.
// For example, let the previous minEle be 3. Now we want to insert 2. We update minEle as 2 and insert 2*2 – 3 = 1 into the stack

// POP
// Remove the element from the top. Let the removed element be y. Two cases arise:
// If y is greater than or equal to minEle, the minimum element in the stack is still minEle.
// If y is less than minEle, the minimum element now becomes (2*minEle – y), so update (minEle = 2*minEle – y). This is where we retrieve the previous minimum from the current minimum and its value in the stack.
// For example, let the element to be removed be 1 and minEle be 2. We remove 1 and update minEle as 2*2 – 1 = 3

type MinStack struct {
	S    []int
	Top  int
	Size int
	Min  int
}

func InitMinStack(size int) *MinStack {
	return &MinStack{
		S:    make([]int, size),
		Top:  -1,
		Size: size,
	}
}

func (s *MinStack) IsMinStackEmpty() bool {
	return s.Top == -1
}

func (s *MinStack) MinStackSize() int {
	return s.Top + 1
}

func (s *MinStack) MinStackTop() int {
	return s.Top
}

func (s *MinStack) MinStackPeek() int {
	return s.S[s.Top]
}

func (s *MinStack) IsMinStackFull() bool {
	return s.Top == s.Size-1
}

func (s *MinStack) MinPush(x int) {
	if s.IsMinStackFull() {
		return
	}
	if s.IsMinStackEmpty() {
		s.Top++
		s.S[s.Top] = x
		s.Min = x
	}
	if x >= s.Min {
		s.Top++
		s.S[s.Top] = x
	} else {
		s.Top++
		s.S[s.Top] = 2*x - s.Min
		s.Min = x
	}
}

func (s *MinStack) Pop() int {
	if s.IsMinStackEmpty() {
		return -100
	}
	var x int = s.S[s.Top]
	s.S[s.Top] = 0
	s.Top--
	if x >= s.Min {
		return x
	} else {
		s.Min = 2*s.Min - x
	}
	return x
}

func (s *MinStack) MinDisplay() {
	if s.IsMinStackEmpty() {
		return
	}
	for i := 0; i <= s.Top; i++ {
		fmt.Println("--->", s.S[i])
	}
}

func (s *MinStack) GetMin() int {
	return s.Min
}
