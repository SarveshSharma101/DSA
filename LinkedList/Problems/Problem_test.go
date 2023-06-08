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

func TestAddKthElementInLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)
	head = AddElementatKthPositioninLinkedList(&head, 0, 0)

	AddElementatKthPositioninLinkedList(&head, 5, 0)

	AddElementatKthPositioninLinkedList(&head, 4, 23)
	AddElementatKthPositioninLinkedList(&head, 3, 3)
	PrintLinkedList(&head)
}

func TestRemoveKthElementInLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)
	head = RemoveElementatKthPositioninLinkedList(&head, 4)

	PrintLinkedList(&head)
}

func TestAppendLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)
	var head2 LL = LL{
		Data: 6,
	}

	Add(&head2, 7)
	Add(&head2, 8)
	Add(&head2, 9)
	Add(&head2, 0)
	head = AppendLinkedLists(&head, &head2)

	PrintLinkedList(&head)
}

func TestReverseLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)
	head = ReverseaLinkedList(&head)

	PrintLinkedList(&head)
}

func TestRemoveOccurenceFromLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 2)
	Add(&head, 2)
	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 2)
	Add(&head, 5)
	Add(&head, 2)
	Add(&head, 6)

	head = RemoveoccurrencesinLinkedList(&head, 2)

	PrintLinkedList(&head)
}

func TestPrintMiddleElement(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)

	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 2)
	Add(&head, 5)
	Add(&head, 2)
	Add(&head, 6)

	MiddleElementofLinkedList(&head)

}
