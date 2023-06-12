package problems

func DeleteNodeXthEndofLinkedList(head *LL, x int) LL {
	temp := head
	var i int
	for temp != nil {
		temp = temp.Next
		i++
	}
	var c int = i - (x)
	if c == 0 {
		head = head.Next
		return *head
	}
	temp = head
	i = 0
	for i != c-1 {
		temp = temp.Next
		i++
	}

	temp.Next = temp.Next.Next
	return *head
}
