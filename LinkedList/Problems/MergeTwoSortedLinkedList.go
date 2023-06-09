package problems

func MergeTwoSortedLinkedList(head1 *LL, head2 *LL) LL {

	var C LL
	if head1.Data < head2.Data {
		C.Data = head1.Data
		head1 = head1.Next
	} else {
		C.Data = head2.Data
		head2 = head2.Next
	}
	var temp = &C
	for head1 != nil && head2 != nil {
		var n LL

		if head1.Data < head2.Data {
			n.Data = head1.Data
			head1 = head1.Next
		} else {
			n.Data = head2.Data
			head2 = head2.Next
		}

		temp.Next = &n
		temp = temp.Next
		// break
	}

	for head1 != nil {
		var n LL

		n.Data = head1.Data
		head1 = head1.Next

		temp.Next = &n
		temp = temp.Next
	}

	for head2 != nil {
		var n LL

		n.Data = head2.Data
		head2 = head2.Next

		temp.Next = &n
		temp = temp.Next
	}
	return C
}
