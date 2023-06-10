package problems

func DeleteDuplicateNodes(head *LL) {
	temp := head
	prev := head

	for temp != nil && temp.Next != nil {
		if temp.Data == temp.Next.Data {
			for temp.Next != nil && temp.Data == temp.Next.Data {
				temp = temp.Next
			}
			prev.Next = temp.Next
		}

		temp = temp.Next
		prev = temp
	}
}
