package problems

func RemoveLoopFromLinkedList(head *LL) {
	var C []*LL

	temp := head
	for {
		if temp == nil {
			break
		}
		C = append(C, temp)
		if search(C, temp.Next) {
			temp.Next = nil
		}
		temp = temp.Next
	}
}
