package problems

func RemoveElementatKthPositioninLinkedList(head *LL, k int) LL {

	if k == 0 {
		head = head.Next
		return *head
	}
	temp := head
	i := 0

	for temp.Next != nil {

		if i+1 == k {
			break
		}
		i++
		temp = temp.Next
	}

	temp.Next = temp.Next.Next

	return *head
}
