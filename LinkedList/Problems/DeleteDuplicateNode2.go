package problems

func DeleteDuplicateNodes2(head *LL) LL {

	for head != nil && head.Next != nil && head.Data == head.Next.Data {
		for {
			head = head.Next
			if head.Data != head.Next.Data {
				head = head.Next
				break
			}
		}
	}
	if head == nil || head.Next == nil {
		return *head
	}

	temp := head
	var prev *LL

	for temp != nil && temp.Next != nil {
		prev = temp
		temp = temp.Next
		for temp.Next != nil && temp.Data == temp.Next.Data {
			for temp.Next != nil && temp.Data == temp.Next.Data {
				temp = temp.Next
				if temp.Data != temp.Next.Data {
					temp = temp.Next
					break

				}
			}
			prev.Next = temp
		}

	}
	return *head
}
