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

func TestMerge2SortedLL(test *testing.T) {
	var head LL = LL{
		Data: 2,
	}

	Add(&head, 3)

	Add(&head, 7)
	var head2 LL = LL{
		Data: 1,
	}
	Add(&head2, 4)
	Add(&head2, 7)

	h := MergeTwoSortedLinkedList(&head, &head2)

	PrintLinkedList(&h)

}

func TestLLPallindrome(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)

	Add(&head, 2)
	Add(&head, 1)
	Add(&head, 1)

	fmt.Println(LinkedListPalindrome(&head))

}

func TestFindINtersection(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)

	Add(&head, 2)
	Add(&head, 1)
	Add(&head, 1)

	var head2 LL = LL{
		Data: 0,
	}

	Add(&head2, 0)

	Add(&head2, 3)
	Add(&head2, 4)
	Add(&head2, 5)
	fmt.Println(FindIntersectionOfLL(&head, &head2))

}

func TestRemoveDuplicateNodes(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)

	Add(&head, 2)
	Add(&head, 1)
	Add(&head, 1)
	Add(&head, 3)
	Add(&head, 3)
	Add(&head, 1)

	DeleteDuplicateNodes(&head)
	PrintLinkedList(&head)
}

func TestRemoveDuplicateNodes2(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 1)

	Add(&head, 2)
	Add(&head, 1)
	Add(&head, 1)
	Add(&head, 3)
	Add(&head, 3)
	Add(&head, 1)

	head = DeleteDuplicateNodes2(&head)
	PrintLinkedList(&head)
}

func TestXthLastElementofLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)

	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)
	Add(&head, 6)
	Add(&head, 7)
	Add(&head, 8)

	x := FindxthNodefromEndofLinkedList(&head, 2)
	fmt.Println("Xth element: ", x)
}

func TestDeleteXthLastElementofLL(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)

	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)
	Add(&head, 6)
	Add(&head, 7)
	Add(&head, 8)

	head = DeleteNodeXthEndofLinkedList(&head, 4)
	PrintLinkedList(&head)
}

func TestAddTwiNumberLL(test *testing.T) {
	var head1 LL = LL{
		Data: 2,
	}

	Add(&head1, 3)

	Add(&head1, 1)
	var head2 LL = LL{
		Data: 1,
	}

	Add(&head2, 4)

	Add(&head2, 5)
	Res := AddTwoNumbersasLists(&head1, &head2)
	PrintLinkedList(&Res)
}

func TestReverseLLII(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 5)

	Add(&head, 7)
	Add(&head, 13)

	head = ReverseaLinkedListII(&head, 2, 3)
	PrintLinkedList(&head)
}

func TestReverseLLGrp(test *testing.T) {
	var head LL = LL{
		Data: 1,
	}

	Add(&head, 2)
	Add(&head, 3)
	Add(&head, 4)
	Add(&head, 5)
	Add(&head, 6)
	Add(&head, 7)
	Add(&head, 8)
	// Add(&head, 9)

	head = ReverseaLinkedListinkgroups(&head, 3)
	PrintLinkedList(&head)
}
