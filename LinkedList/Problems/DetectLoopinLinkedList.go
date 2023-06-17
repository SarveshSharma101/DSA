package problems

func DetectinLinkedList(head *LL) int {
	var x int = -1
	var C []*LL

	temp := head
	for {
		if temp == nil {
			break
		}
		C = append(C, temp)
		if search(C, temp.Next) {
			return temp.Next.Data
		}
		temp = temp.Next
	}

	return x
}

func search(C []*LL, node *LL) bool {
	for _, n := range C {
		if n == node {
			return true
		}
	}
	return false
}
