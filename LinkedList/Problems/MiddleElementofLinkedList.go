package problems

import "fmt"

func MiddleElementofLinkedList(head *LL) {
	var C []int

	temp := head

	for temp != nil {
		C = append(C, temp.Data)
		temp = temp.Next
	}
	fmt.Println("--->", C[(len(C)-1)/2])
}
