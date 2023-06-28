package problems

import (
	"fmt"
	"unicode"
)

func SimplifyDirectoryPath(path string) string {
	var sp string
	stack := InitSPStack(1000)

	path_byte := []byte(path)
	i := 0
	for i < len(path_byte) {
		// v := path_byte[i]
		s := ""
		for i < len(path_byte) && unicode.IsLetter(rune(path_byte[i])) {
			s = s + string(path_byte[i])
			i++
		}
		//push
		if len(s) > 0 {
			stack.SPStackPush(s)

		}

		for i < len(path_byte) && string(path_byte[i]) == "/" {
			i++
		}

		if i < len(path_byte) && string(path_byte[i]) == "." {
			i++
			if i < len(path_byte) && string(path_byte[i]) == "." {
				i++
				//pop
				if !stack.IsSPStackEmpty() {
					stack.SPStackPop()
				}
			}
			continue
		}

	}
	stack1 := InitSPStack(1000)
	for !stack.IsSPStackEmpty() {
		stack1.SPStackPush(stack.SPStackPop())
	}
	for !stack1.IsSPStackEmpty() {
		sp = sp + "/" + stack1.SPStackPop()
	}
	if len(sp) == 0 {
		sp = "/"
	}

	return sp
}

type SPStack struct {
	S    []string
	Top  int
	Size int
}

func InitSPStack(size int) *SPStack {
	return &SPStack{
		S:    make([]string, size),
		Top:  -1,
		Size: size,
	}
}

func (s *SPStack) IsSPStackEmpty() bool {
	return s.Top == -1
}

func (s *SPStack) SPStackSize() int {
	return s.Top + 1
}

func (s *SPStack) SPStackTop() int {
	return s.Top
}

func (s *SPStack) SPStackPeek() string {
	return s.S[s.Top]
}

func (s *SPStack) SPIsFull() bool {
	return s.Top == s.Size-1
}

func (s *SPStack) SPStackPush(x string) {
	if s.SPIsFull() {
		return
	}
	s.Top = s.Top + 1
	s.S[s.Top] = x
}

func (s *SPStack) SPStackPop() string {
	if s.IsSPStackEmpty() {
		return ""
	}
	x := s.S[s.Top]
	s.Top = s.Top - 1
	return x
}

func (s *SPStack) Display() {
	if s.IsSPStackEmpty() {
		return
	}
	for i := 0; i <= s.Top; i++ {
		fmt.Println("--->", s.S[i])
	}
}
