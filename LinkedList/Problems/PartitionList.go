package problems

func PartitionList(head *LL, k int) LL {
	temp1 := head

	var C []*LL

	for temp1.Data != k {
		C = append(C, temp1)
		temp1 = temp1.Next
	}
	C = append(C, temp1)

	temp1 = temp1.Next

	var B []*LL

	for temp1 != nil {
		B = append(B, temp1)
		temp1 = temp1.Next
	}

	quickSort(&C, 0, len(C)-1)
	quickSort(&B, 0, len(B)-1)

	i := 0
	j := 0
	var headNode LL

	if C[i].Data < B[j].Data {
		headNode.Data = C[i].Data
		i++
	} else {
		headNode.Data = B[j].Data
		j++
	}
	temp2 := &headNode

	for i < len(C) && j < len(B) {
		if C[i].Data <= B[j].Data {
			node := LL{Data: C[i].Data}
			temp2.Next = &node
			temp2 = temp2.Next
			i++
		} else {
			node := LL{Data: B[j].Data}
			temp2.Next = &node
			temp2 = temp2.Next
			j++
		}
	}

	for i < len(C) {
		node := LL{Data: C[i].Data}
		temp2.Next = &node
		temp2 = temp2.Next
		i++
	}

	for j < len(B) {
		node := LL{Data: B[j].Data}
		temp2.Next = &node
		temp2 = temp2.Next
		j++
	}

	return headNode
}

func quickSort(A *[]*LL, l, h int) {
	if l < h {

		p := findPivot(A, l, h)
		quickSort(A, l, p-1)
		quickSort(A, p+1, h)
	}
}

func findPivot(A *[]*LL, l, h int) int {
	p := (*A)[h]
	i := l
	j := h - 1

	for {
		for i < len(*A) && (*A)[i].Data < p.Data {
			i++
		}
		for j >= 0 && (*A)[j].Data > p.Data {
			j--
		}

		if i >= j {
			break
		}
		temp := (*A)[i]
		(*A)[i] = (*A)[j]
		(*A)[j] = temp
	}

	temp := (*A)[i]
	(*A)[i] = (*A)[h]
	(*A)[h] = temp

	return i
}
