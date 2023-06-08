package problems

func RemoveoccurrencesinLinkedList(head *LL, k int) LL {

	for head.Data == k {
		head = head.Next
	}
	temp := head
	prev := head
	for temp != nil {
		temp = temp.Next
		for temp != nil && temp.Data == k {
			prev.Next = temp.Next
			temp = temp.Next
		}
		prev = temp

	}
	return *head
}
