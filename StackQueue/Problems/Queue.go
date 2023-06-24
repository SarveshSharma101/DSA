package problems

type Queue struct {
	Q     []int
	Front int
	Rear  int
	Size  int
}

func InitQueue(size int) *Queue {
	return &Queue{
		Q:     make([]int, size),
		Size:  size,
		Rear:  -1,
		Front: -1,
	}
}

func (q *Queue) QIsFull() bool {
	return q.Rear == q.Size
}

func (q *Queue) QSize() int {
	if q.QIsEmpty() {
		return 0
	}
	return (q.Rear + 1) - (q.Front)
}

func (q *Queue) QIsEmpty() bool {
	return (q.Rear < q.Front) || (q.Rear == -1 && q.Front == -1)
}

func (q *Queue) QFront() int {
	return q.Front
}

func (q *Queue) QRear() int {
	return q.Rear
}

func (q *Queue) QPut(x int) {
	if q.QIsFull() {
		return
	}
	if q.QIsEmpty() {
		q.Rear = 0
		q.Front = 0
		q.Q[q.Rear] = x
		return
	}
	q.Rear = q.Rear + 1
	q.Q[q.Rear] = x

}

func (q *Queue) QDelete() int {
	if q.QIsEmpty() {
		return -100
	}
	x := q.Q[q.Front]
	q.Q[q.Front] = 0
	q.Front++
	return x
}
