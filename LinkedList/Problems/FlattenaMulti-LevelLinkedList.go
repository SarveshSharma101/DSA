package problems

import "fmt"

func FlattenaMultiLevelLinkedList(head *LLD) LL {
	var C []int
	temp := head

	for temp != nil {
		C = append(C, temp.Data)
		tempD := temp.Down
		for tempD != nil {
			C = append(C, tempD.Data)
			tempD = tempD.Down
		}
		temp = temp.Next
	}
	fmt.Println("Before--->", C)

	QuickSortLL(&C, 0, len(C)-1)
	fmt.Println("After--->", C)
	headNode := LL{Data: C[0]}
	temp1 := &headNode
	for i := 1; i < len(C); i++ {
		if i == len(C)-1 {
			temp1.Next = nil
			continue
		}
		temp1.Next = &LL{Data: C[i]}
		temp1 = temp1.Next
	}
	return headNode
}

func QuickSortLL(C *[]int, l, h int) {
	if l < h {
		p := partition(C, l, h)
		QuickSortLL(C, l, p-1)
		QuickSortLL(C, p+1, h)
	}
}

func partition(C *[]int, l, h int) int {
	i := l
	j := h - 1
	p := (*C)[h]

	for {

		for i < h && (*C)[i] <= p {
			i++
		}

		for j > -1 && (*C)[j] > p {
			j--
		}
		if i >= j {
			break
		}
		temp := (*C)[i]
		(*C)[i] = (*C)[j]
		(*C)[j] = temp
	}
	temp := (*C)[i]
	(*C)[i] = p
	(*C)[h] = temp
	return i
}
