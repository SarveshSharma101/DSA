package problems

func AppendLinkedLists(head1 *LL, head2 *LL) LL {
	temp := head1
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = head2
	return *head1
}
