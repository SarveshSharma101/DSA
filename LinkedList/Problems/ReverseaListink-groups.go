package problems

import "fmt"

func ReverseaLinkedListinkgroups(head *LL, x int) LL {

	var B []*LL

	var count int = 1

	temp := head

	var n int

	for temp != nil {
		n++
		temp = temp.Next
	}
	temp = head

	for i := 0; i < n/x; i++ {
		var C []*LL

		for count <= x {
			C = append(C, temp)
			temp = temp.Next
			count++
		}
		headNode := &LL{
			Data: C[len(C)-1].Data,
		}
		B = append(B, headNode)

		for i := len(C) - 2; i >= 0; i-- {
			add(headNode, C[i].Data)
		}
		// temp = temp.Next
		count = 1
	}

	for temp != nil {
		B = append(B, &LL{Data: temp.Data})
		temp = temp.Next
	}

	for i := 0; i <= len(B)-2; i++ {
		t := B[i]
		for t.Next != nil {
			t = t.Next
		}
		t.Next = B[i+1]
	}
	fmt.Println(B)
	return *B[0]
}

func add(head *LL, data int) {
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	node := LL{
		Data: data,
	}
	temp.Next = &node
}
