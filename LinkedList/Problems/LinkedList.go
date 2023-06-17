package problems

type LL struct {
	Data int
	Next *LL
}

type LLD struct {
	Data int
	Next *LLD
	Down *LLD
}

func Add(head *LL, data int) LL {
	if head == nil {
		head = &LL{
			Data: data,
			Next: nil,
		}
		return *head
	}

	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}

	node := LL{
		Data: data,
		Next: nil,
	}
	temp.Next = &node

	return node
}

func AddNext(head *LLD, data int) *LLD {
	if head == nil {
		head = &LLD{
			Data: data,
			Next: nil,
		}
		return head
	}

	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}

	node := LLD{
		Data: data,
		Next: nil,
	}
	temp.Next = &node

	return &node
}

func AddDown(head *LLD, data int) LLD {
	if head == nil {
		head = &LLD{
			Data: data,
			Next: nil,
		}
		return *head
	}

	temp := head
	for temp.Down != nil {
		temp = temp.Down
	}

	node := LLD{
		Data: data,
		Next: nil,
	}
	temp.Down = &node

	return node
}
