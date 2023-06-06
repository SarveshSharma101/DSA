package problems

import (
	"fmt"
	"testing"
)

func TestPrintLinkedList(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)

	PrintLinkedList(&head)
}

func TestLinkedToList(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)

	fmt.Println("----->", LinkedListtoArray(&head))
}

func TestReversedLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)

	PrintReversedLinkedList(&head)
}

func TestKthElementInLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)

	KthElementinLinkedList(&head, 4)
}
