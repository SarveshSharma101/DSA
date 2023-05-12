package searching

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{3, 41, 223, 12, 433, 34, 78, 78798, 654, 12}
	e := []int{41, 223, 654, 78, 0, 12, 99}
	for _, v := range e {
		i := LinearSearch(&arr, v)
		if i < 0 {
			fmt.Printf("%v is not present in the array", v)
		}
		fmt.Println(v, " found at position: ", i)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{3, 41, 223, 12, 433, 34, 78, 78798, 654, 12}
	e := []int{41, 223, 654, 78, 0, 12, 99}
	for _, v := range e {
		i := BinarySearch(&arr, v)
		if i < 0 {
			fmt.Printf("%v is not present in the array", v)
		}
		fmt.Println(v, " found at position: ", i)
	}
}
