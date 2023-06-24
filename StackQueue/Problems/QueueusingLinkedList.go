package problems

import "fmt"

type LLQueue struct {
	QueueHead *LL
	Front     int
	Rear      int
	Size      int
}

func InitLLQueue(size int) *LLQueue {
	return &LLQueue{
		QueueHead: nil,
		Front:     0,
		Rear:      -1,
		Size:      size,
	}
}

func (q *LLQueue) IsLLQueueEmpty() bool {
	return q.QueueHead == nil
}

func (q *LLQueue) LLQSize() int {
	return (q.Rear + 1) - (q.Front)
}

func (q *LLQueue) LLQFront() int {
	return q.Front
}

func (q *LLQueue) LLQRear() int {
	return q.Rear
}

func (q *LLQueue) QPut(x int) {
	if q.IsLLQueueEmpty() {
		node := LL{
			Data: x,
		}
		q.QueueHead = &node
		q.Rear++
		return
	}
	node := LL{
		Data: x,
	}
	temp := q.QueueHead

	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &node
	q.Rear++
}

func (q *LLQueue) QDelete() int {
	if q.IsLLQueueEmpty() {
		return -100
	}

	x := q.QueueHead.Data
	q.QueueHead = q.QueueHead.Next
	q.Front++
	return x
}

func (s *LLQueue) QDisplay() {
	if s.IsLLQueueEmpty() {
		return
	}
	temp := s.QueueHead

	for temp != nil {
		fmt.Println("-->", temp.Data)
		temp = temp.Next
	}
}
