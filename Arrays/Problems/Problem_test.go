package problems

import (
	"fmt"
	"testing"
)

func TestCumalitiveSum(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 2, 3, 4}, {1, 1, 1, 1, 1}, {1, 3, 5, 7, 9}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Cumalitive Sum: ", CumulativeSum(&v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestPositiveCumalitiveSum(t *testing.T) {
	var TestArr [][]int = [][]int{{1, -2, 3, 4, -6}, {1, -1, -1, -1, 1}, {1, 3, 5, 7}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Positive Cumalitive Sum: ", PositiveCumulativeSum(&v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestIdenticalTwins(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 2, 3, 2, 1}, {1, 2, 2, 3, 2, 1}, {1, 1, 1, 1}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Number of identical twins: ", IdenticalTwins(&v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestEvenNumberOfDigit(t *testing.T) {
	var TestArr [][]int = [][]int{{11, 42, 564, 5775, 34, 123, 454, 1, 5, 45, 3556, 23442}, {3, 11, 4, 200}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Number with even num of digit: ", EvenNumberOfDigits(&v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestInsertionSort(t *testing.T) {
	var TestArr [][]int = [][]int{{5, 4, 2, 5, 3, 1}, {3, 11, 4, 200}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		InsertionSort(&v)
		fmt.Println("Sorted array: ", v)
		fmt.Println("------------------------------------------------------------")
	}
}

func TestMerge2SortedArray(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 3, 3, 4, 4}, {5, 6}}, {{1, 3, 3, 3, 3, 4}, {9, 11}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		C := Merge2SortedArray(&v[0], &v[1])
		fmt.Println("Sorted array: ", C)
		fmt.Println("------------------------------------------------------------")
	}
}

func TestMergeSortedSubArray(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 3, 5, 7, 9, 11, 0, 4, 6, 8}, {10, 5}}, {{3, 3, 9, 11, 1, 3, 3, 4}, {8, 3}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		MergeSortedSubArray(&v[0], v[1][1])
		fmt.Println("Sorted array: ", v[0])
		fmt.Println("------------------------------------------------------------")
	}
}

func TestMergeSort(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 2, 3, 4, 5}, {10, 5, 3131, 2, 1222, 1}, {3, 3, 9, 11, 1, 3, 3, 4}, {18, 2, 3, 12, 32, 12, 3}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		MergeSort(&v, 0, len(v)-1)
		fmt.Println("Sorted array: ", v)
		fmt.Println("------------------------------------------------------------")
	}
}

func TestQuickSort(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 2, 3, 4, 5}, {10, 5, 3131, 2, 1222, 1}, {5, 4, 2, 5, 3, 1}, {3, 11, 4, 200}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		QuickSort(&v, 0, len(v)-1)
		fmt.Println("Sorted array: ", v)
		fmt.Println("------------------------------------------------------------")
	}
}

func TestSqureSortArray(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 2, 3, 4, 5}, {10, 5, 3131, 2, 1222, 1}, {5, 4, 2, 5, 3, 1}, {3, 11, 4, 200}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Sorted array: ", SquraeSortedArray(&v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestMaxOnes(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 1, 3, 2, 3, 1, 1, 1}, {1, 1, 1, 2, 2}, {1, 2, 1, 1, 1, 2}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Max ones: ", MaxConsecutiveOnes(&v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestAP(t *testing.T) {
	var TestArr [][]int = [][]int{{9, 13, 5, 15, 7, 11}, {1, 1, 1}, {4, 1, 2}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Max ones: ", ArithmeticSequence(v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestLargestContigiousSum(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 2, 3, 4, 5}, {1, 1, 1}, {4, -6, 2, 5}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Max ones: ", LargestContiguousSum(&v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestPascalTriangle(t *testing.T) {
	var TestArr []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Max ones: ", PascalTriangle(v))
		fmt.Println("------------------------------------------------------------")
	}
}

func TestRowCloumnTest(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 1, 0}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 0}, {1, 1, 1}}, {{1, 0}, {1, 0}, {1, 0}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ")
		for _, c := range v {
			fmt.Println(c)
		}
		RowColumnZero(&v)
		fmt.Println("Result: ")
		for _, c := range v {
			fmt.Println(c)
		}
		fmt.Println("------------------------------------------------------------")
	}
}

func TestRotateMatrix(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{1, 2}, {3, 4}, {5, 6}}, {{1, 0}, {1, 0}, {1, 0}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ")
		for _, c := range v {
			fmt.Println(c)
		}
		B := RotateMatrix(&v)
		fmt.Println("Result: ")
		for _, c := range B {
			fmt.Println(c)
		}
		fmt.Println("------------------------------------------------------------")
	}
}

func TestPrimeUptoN(t *testing.T) {
	var TestArr []int = []int{5, 1, 12}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Result: ", PrimesuptoN(v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestMergeOverlappingIntervals(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2}, {2, 3}, {1, 4}, {5, 6}}, {{1, 1}, {2, 2}, {3, 3}}, {{1, 4}, {6, 7}, {4, 5}}, {{1, 2}, {2, 3}, {5, 5}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Result: ", MergeOverlappingIntervals(v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestKthLargestElement(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{4, 3, 2, 1}, {2}}, {{1, 2, 3, 4, 5}, {5}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)

		fmt.Println("Result: ", KthLargestElement(&v[0], v[1][0]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestNextGreaterPermutation(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 3, 2}, {3, 2, 1}, {2, 2, 9}, {2, 9, 9}, {4}, {1, 2, 3, 4}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		NextGreaterPermutation(&v)
		fmt.Println("Result: ", v)

		fmt.Println("------------------------------------------------------------")
	}
}

func TestInversionCount(t *testing.T) {
	var TestArr [][]int = [][]int{{8, 4, 1, 2}, {1, 2, 2, 3}, {3, 2, 1}, {10, 1, 2, 3, 4}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", InversionCount(&v))

		fmt.Println("------------------------------------------------------------")
	}
}
