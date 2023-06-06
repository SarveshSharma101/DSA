package problems

type LL struct {
	Data int
	Next *LL
}

func Add(head *LL, data int) {
	if head == nil {
		head = &LL{
			Data: data,
			Next: nil,
		}
		return
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
}
