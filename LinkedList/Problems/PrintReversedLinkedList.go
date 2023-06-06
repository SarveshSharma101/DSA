package problems

import "fmt"

func PrintReversedLinkedList(head *LL) {
	temp := head

	var C []int

	for temp != nil {
		C = append(C, temp.Data)
		temp = temp.Next
	}

	n := len(C)
	for i := n - 1; i >= 0; i-- {
		fmt.Println("-->", C[i])
	}
}
