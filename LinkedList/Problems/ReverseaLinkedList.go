package problems

func ReverseaLinkedList(head *LL) LL {
	var C []int

	temp := head
	for temp != nil {
		C = append(C, temp.Data)
		temp = temp.Next
	}

	var head2 LL = LL{
		Data: C[len(C)-1],
	}

	for i := len(C) - 2; i >= 0; i-- {
		Add(&head2, C[i])

	}

	return head2
}
