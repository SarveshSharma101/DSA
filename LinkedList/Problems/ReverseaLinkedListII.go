package problems

func ReverseaLinkedListII(head *LL, x, y int) LL {
	var C []*LL
	var i int = 1
	temp := head

	for {
		i++
		if i == x {
			break
		}
		temp = temp.Next
	}

	j := i

	prev := temp.Next
	for prev != nil {

		C = append(C, prev)
		prev = prev.Next
		j++
		if j > y {
			break
		}
	}

	C[0].Next = C[len(C)-1].Next
	for i := len(C) - 1; i >= 0; i-- {

		node := C[i]
		temp.Next = node
		temp = temp.Next
	}
	return *head
}
