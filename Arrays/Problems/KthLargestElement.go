package problems

import "fmt"

// Kth Largest Element
// Medium
// 50
// Given a list of numbers below:
// 4, 3, 6, 4, 1

// What is the largest element in the list? → 6

// What is the 3rd largest element in the list? → 4

// Given a list of numbers, find the kth largest element in the list.

func KthLargestElement(a *[]int, k int) int {
	MergeSort(a, 0, len(*a)-1)
	fmt.Println("-------->", a)
	return (*a)[k-1]
}
