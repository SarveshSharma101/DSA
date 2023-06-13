package problems

func AddTwoNumbersasLists(head1 *LL, head2 *LL) LL {
	var Res LL = LL{
		Data: head1.Data + head2.Data,
	}
	temp1 := head1.Next
	temp2 := head2.Next
	temp := &Res

	for temp1 != nil {
		node := LL{
			Data: temp1.Data + temp2.Data,
		}
		temp.Next = &node
		temp1 = temp1.Next
		temp2 = temp2.Next
		temp = temp.Next
	}

	return Res
}
