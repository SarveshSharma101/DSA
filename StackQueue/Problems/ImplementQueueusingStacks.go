package problems

// Use 2 stack s1, S2
// Enque :
// 	s1 -> s2
// 	x -> s1
// 	s2 -> s1

type QUS struct {
	S1    *Stack
	S2    *Stack
	Size  int
	Front int
	Rear  int
}

func InitQUS(size int) *QUS {
	return &QUS{
		S1:    InitStack(size),
		S2:    InitStack(size),
		Size:  size,
		Front: -1,
		Rear:  -1,
	}
}

func (q *QUS) Enque(x int) {
	if q.S1.IsFull() {
		return
	}
	q.Rear++
	if q.S1.IsEmpty() {
		q.S1.Push(x)
		return
	}

	for !q.S1.IsEmpty() {
		q.S2.Push(q.S1.Pop())
	}
	q.S1.Push(x)
	for !q.S2.IsEmpty() {
		q.S1.Push(q.S2.Pop())
	}
}

func (q *QUS) Dequeue() int {
	if !q.S1.IsEmpty() {
		q.Front++
		return q.S1.Pop()
	}
	return -100
}

func (q *QUS) QusSize() int {
	return q.S1.StackSize()
}

func (q *QUS) QusIsEmpty() bool {
	return q.S1.IsEmpty()
}

func (q *QUS) QusFrontElement() int {
	return q.Front + 1
}

func (q *QUS) QusRearElement() int {
	return q.Rear + 1
}
