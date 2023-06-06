package problems

func LinkedListtoArray(head *LL) []int {
	var C []int

	temp := head

	for temp != nil {

		C = append(C, temp.Data)
		temp = temp.Next

	}

	return C
}
