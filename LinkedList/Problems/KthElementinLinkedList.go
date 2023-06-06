package problems

import "fmt"

func KthElementinLinkedList(head *LL, k int) {
	i := 0
	temp := head

	for temp != nil {
		if i == k {
			fmt.Println("-->", temp.Data)
			return
		}
		temp = temp.Next
		i++
	}
}
