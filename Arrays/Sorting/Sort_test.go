package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 41, 223, 12, 433, 34, 78, 78798, 654, 12}
	fmt.Println("Unsorted: ", arr)
	BubbleSort(&arr)
	fmt.Println("Sorted using Bubble sort: ", arr)

}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 41, 223, 12, 433, 34, 78, 78798, 654, 12}
	fmt.Println("Unsorted: ", arr)
	QuickSort(&arr, 0, len(arr)-1)
	fmt.Println("Sorted using Quick sort: ", arr)

}

func TestMergeSort(t *testing.T) {
	arr := []int{3, 41, 223, 12, 433, 34, 78, 78798, 654, 12}
	fmt.Println("Unsorted: ", arr)
	MergeSort(&arr, 0, len(arr)-1)
	fmt.Println("Sorted using Merge sort: ", arr)

}

func TestInsertionSort(t *testing.T) {
	arr := []int{3, 41, 223, 12, 433, 34, 78, 78798, 654, 12}
	fmt.Println("Unsorted: ", arr)
	InsertionSort(&arr)
	fmt.Println("Sorted using Insertion sort: ", arr)

}

func TestCountSort(t *testing.T) {
	arr := []int{12, 12, 13, 12, 4, 5, 0, 0, 1}
	fmt.Println("Unsorted: ", arr)
	CountSort(&arr, len(arr), 13)
	fmt.Println("Sorted using Count sort: ", arr)

}
