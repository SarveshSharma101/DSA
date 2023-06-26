package problems

import "fmt"

func IsBalancedParentheses(str string) bool {
	stack := InitBPStack(100)
	for _, char := range str {
		if string(char) == "(" || string(char) == "{" || string(char) == "[" {
			stack.BPStackPush(string(char))
		} else {
			if stack.IsBPStackEmpty() {
				return false
			}
			if string(char) == ")" {
				if stack.BPStackPeek() != "(" {
					return false
				}
				stack.BPStackPop()
			} else if string(char) == "}" {
				if stack.BPStackPeek() != "{" {
					return false
				}
				stack.BPStackPop()
			} else if string(char) == "]" {
				if stack.BPStackPeek() != "[" {
					return false
				}
				stack.BPStackPop()
			}

		}
	}

	return stack.IsBPStackEmpty()
}

type BPStack struct {
	S    []string
	Top  int
	Size int
}

func InitBPStack(size int) *BPStack {
	return &BPStack{
		S:    make([]string, size),
		Top:  -1,
		Size: size,
	}
}

func (s *BPStack) IsBPStackEmpty() bool {
	return s.Top == -1
}

func (s *BPStack) BPStackSize() int {
	return s.Top + 1
}

func (s *BPStack) BPStackTop() int {
	return s.Top
}

func (s *BPStack) BPStackPeek() string {
	return s.S[s.Top]
}

func (s *BPStack) BPIsFull() bool {
	return s.Top == s.Size-1
}

func (s *BPStack) BPStackPush(x string) {
	if s.BPIsFull() {
		return
	}
	s.Top = s.Top + 1
	s.S[s.Top] = x
}

func (s *BPStack) BPStackPop() string {
	if s.IsBPStackEmpty() {
		return ""
	}
	x := s.S[s.Top]
	s.Top = s.Top - 1
	return x
}

func (s *BPStack) Display() {
	if s.IsBPStackEmpty() {
		return
	}
	for i := 0; i <= s.Top; i++ {
		fmt.Println("--->", s.S[i])
	}
}
