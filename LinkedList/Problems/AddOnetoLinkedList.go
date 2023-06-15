package problems

func AddOnetoLinkedList(head *LL) LL {
	var C []*LL
	temp := head

	for temp != nil {
		C = append(C, temp)
		temp = temp.Next
	}

	var i int
	for i = len(C) - 1; i >= 0; i-- {
		node := C[i]
		node.Data = node.Data + 1
		if node.Data < 10 {
			break
		}
		node.Data = 0
	}
	if i < 0 {
		node := LL{
			Data: 1,
		}
		node.Next = head
		return node
	}
	return *head
}
