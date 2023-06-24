package problems

import "fmt"

type SUQ struct {
	Q1   Queue
	Q2   Queue
	Size int
	Top  int
}

func InitSuq(size int) *SUQ {
	return &SUQ{
		Q1:   *InitQueue(size),
		Q2:   *InitQueue(size),
		Size: size,
		Top:  -1,
	}
}

func (s *SUQ) SUQPush(x int) {
	if s.Q1.QIsFull() {
		return
	}

	s.Q2.QPut(x)
	for !s.Q1.QIsEmpty() {
		s.Q2.QPut(s.Q1.QDelete())
	}
	for !s.Q2.QIsEmpty() {
		s.Q1.QPut(s.Q2.QDelete())
	}
	s.Top++
}

func (s *SUQ) SUQPop() int {
	if s.Q1.QIsEmpty() {
		return -100
	}
	s.Top--
	return s.Q1.QDelete()
}

func (s *SUQ) SUQTop() int {
	return s.Top
}

func (s *SUQ) SUQSize() int {
	return s.Q1.QSize()
}

func (s *SUQ) SUQIsEmpty() bool {
	return s.Q1.QIsEmpty()
}

func (s *SUQ) Display() {
	if s.SUQIsEmpty() {
		return
	}
	for i := s.Q1.Front; i <= s.Q1.Rear; i++ {
		fmt.Println("-->", s.Q1.Q[i])
	}
}
