package problems

func FindxthNodefromEndofLinkedList(head *LL, x int) int {
	var X int
	var C []int
	temp := head

	for temp != nil {
		C = append(C, temp.Data)
		temp = temp.Next
	}

	X = C[len(C)-x]

	return X

}
