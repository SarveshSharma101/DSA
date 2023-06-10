package problems

func FindIntersectionOfLL(head1 *LL, head2 *LL) int {
	var X int = -1
	var A []int
	var B []int
	temp1 := head1
	temp2 := head2

	for temp1 != nil {
		A = append(A, temp1.Data)
		temp1 = temp1.Next
	}
	for temp2 != nil {
		B = append(B, temp2.Data)
		temp2 = temp2.Next
	}

	var C []int
	var D []int
	if len(A) > len(B) {
		C = B
		D = A

	} else {
		C = A
		D = B
	}

	for _, v := range C {
		X = Seacrh(v, D)
		if X >= 0 {
			break
		}
	}

	return X
}

func Seacrh(v int, D []int) int {
	for _, v2 := range D {
		if v2 == v {
			return 0
		}
	}
	return -1
}
