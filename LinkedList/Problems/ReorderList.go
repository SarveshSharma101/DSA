package problems

func ReorderList(head *LL) LL {
	temp := head

	var C []*LL

	for temp != nil {
		C = append(C, temp)
		temp = temp.Next
	}
	n := len(C) / 2

	for i := 0; i < n; i++ {
		node := C[len(C)-(i+1)]
		node.Next = C[i].Next
		C[i].Next = node
	}
	C[n].Next = nil
	return *head
}
