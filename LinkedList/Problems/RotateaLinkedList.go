package problems

func RotateALinkedList(head *LL, k int) LL {
	var C []*LL

	temp := head
	for temp != nil {
		C = append(C, temp)
		temp = temp.Next
	}
	var n int = len(C) - 1
	var j int = 1
	var B []int
	for j <= k {
		B = append(B, C[n].Data)
		n--
		j++
	}

	for i := 0; i < (len(C))-k; i++ {
		B = append(B, C[i].Data)
	}

	head2 := LL{
		Data: B[0],
	}

	for i := 1; i < len(B); i++ {
		add(&head2, B[i])
	}
	return head2
}
