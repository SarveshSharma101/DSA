package problems

func LinkedListPalindrome(head *LL) bool {

	var C []int
	temp := head

	for temp.Next != nil {
		C = append(C, temp.Data)
		temp = temp.Next
	}
	C = append(C, temp.Data)
	temp = head
	for i := len(C) - 1; i >= 0; i-- {
		if C[i] != temp.Data {
			return false
		}
		temp = temp.Next
	}
	return true
}
