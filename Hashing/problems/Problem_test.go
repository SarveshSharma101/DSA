package problems

import (
	"fmt"
	"testing"
)

func TestImplementCountingSort(t *testing.T) {
	fmt.Println("1. 4 2 5 3 1:", ImplementCountingSort(&[]int{4, 2, 5, 3, 1}))
	fmt.Println("2. 11 4 200:", ImplementCountingSort(&[]int{11, 4, 200}))
}

func TestIdenticalTwins(t *testing.T) {
	fmt.Println("1. 1, 2, 3, 2, 1:", IdenticalTwins(&[]int{1, 2, 3, 2, 1}))
	fmt.Println("2. 1, 2, 2, 3, 2, 1:", IdenticalTwins(&[]int{1, 2, 2, 3, 2, 1}))
	fmt.Println("2. 1, 1, 1, 1:", IdenticalTwins(&[]int{1, 1, 1, 1}))
}

func TestLongestConsecutiveSequence(t *testing.T) {
	fmt.Println("1. 24, 2, 34, 1, 3, 4:", LongestConsecutiveSequence(&[]int{24, 2, 34, 1, 3, 4}))

}

func TestLongestSubarrayWithZero(t *testing.T) {
	fmt.Println("1. 3, 0, -1, -2, 3, 0, -2:")
	LongestSubarraywithZeroSum(&[]int{3, 0, -1, -2, 3, 0, -2})

}
