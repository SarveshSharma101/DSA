package problems

import "fmt"

func PrintLinkedList(head *LL) {
	temp := head
	for temp != nil {
		fmt.Println("-->", temp.Data)
		temp = temp.Next
	}
}
