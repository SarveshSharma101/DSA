package problems

func AddElementatKthPositioninLinkedList(head *LL, k, x int) LL {
	node := LL{
		Data: x,
	}
	if k == 0 {
		node.Next = head.Next
		return node
	}
	temp := head
	i := 0

	for temp.Next != nil {

		i++
		if i == k {
			break
		}
		temp = temp.Next
	}
	if temp.Next == nil {
		temp.Next = &node
		return *head
	}
	node.Next = temp.Next
	temp.Next = &node
	return *head
}
