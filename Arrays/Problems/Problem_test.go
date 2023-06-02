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

func TestContainsElement(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3, 3, 3, 4, 4, 5}, {2}}, {{1, 2, 3, 3, 3, 4, 4, 5}, {6}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", ContainsElement(&v[0], v[1][0]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestSearchRange(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3, 3, 3, 4, 4, 5}, {3}}, {{1, 2, 3, 3, 3, 4, 4, 5}, {6}}, {{1, 2, 3, 3, 3, 4, 4, 5}, {5}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", SearchRange(&v[0], v[1][0]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestCountNegative(t *testing.T) {
	var TestArr [][]int = [][]int{{-5, -3, -2, 3, 4, 6, 7, 8}, {0, 1, 2, 3, 4, 6, 7, 8}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", NegativeNumbersInSortedArray(&v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestNextGreaterNumber(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3, 3, 4, 4, 8, 10}, {4}}, {{1, 2, 3, 3, 3, 4, 4, 5}, {5}}, {{1, 2, 3, 3, 3, 4, 4, 5}, {-1}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", NextGreaterElementInSortedArray(&v[0], v[1][0]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestInsertinSortedArray(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3, 4, 5}, {3}}, {{1, 2, 3, 5}, {4}}, {{1, 2, 3, 4, 5}, {-1}}, {{1, 2, 3, 4, 5}, {6}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", InsertPositioninSortedArray(&v[0], v[1][0]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestIsPerfectSquare(t *testing.T) {
	var TestArr []int = []int{25, 20, 36, 144, 155, 225}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", IsPerfectSquare(v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestSearchInRotatedArray(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{4, 5, 6, 7, 1, 2, 3}, {6}}, {{4, 5, 6, 7, 0, 1, 2}, {7}}, {{3, 4, 1, 2}, {4}}, {{5, 1, 2, 3, 4}, {5}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", SearchRotatedSortedArray(&v[0], v[1][0]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestNonRepeatingElement(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 1, 2, 3, 3, 4, 4}, {1, 2, 2}, {3, 3, 4, 4, 5}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", NonRepeatingElement(&v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestSquareRoot(t *testing.T) {
	var TestArr []int = []int{25, 20, 36, 144, 155, 225}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", SquareRoot(v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestMatrixSearch(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}}
	keys := []int{6, 15}
	for i, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", MatrixSearch(&v, keys[i]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestMatricMedian(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 6, 7, 7, 8}, {2, 2, 3, 3, 4}, {1, 2, 2, 2, 2}}, {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{1, 2, 3}, {3, 3, 4}, {1, 1, 2}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", MedianofRowwiseSortedMatrix(&v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestRemoveOccurence(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 4, 2, 6, 2, 6, 9, 4}, {4}}, {{1, 1, 1, 2, 2}, {1}}, {{1, 3, 3, 3, 4, 4}, {3}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", RemoveOccurences(&v[0], v[1][0]))
		fmt.Println("--->", v[0])

		fmt.Println("------------------------------------------------------------")
	}
}

func TestTwoSumSorted(t *testing.T) {
	var TestArr [][]int = [][]int{{-3, 1, 3, 4}, {-2, 1, 3, 4}, {3, 3, 4, 4, 5}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", TwoSumSorted(&v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestMergeTwoSrotedArray(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3, 4, 4}, {2, 4, 5, 5}}, {{1, 1, 1, 2, 2}, {1}}, {{1, 3, 3, 3, 4, 4}, {3}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", TwoSortedArrays(&v[0], &v[1]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestKthSum(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{3, 5, 6, 2, 4, 7, 8}, {3}}, {{1, 3, 3, 3, 4, 4}, {1}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", KSubarraySum(&v[0], v[1][0]))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestUniqueElementInSortedArray(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 2, 3, 3, 3, 4, 5, 5}, {1, 1, 1, 2, 2}, {3, 3, 4, 4, 5}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", UniqueElementsinSortedArray(&v))

		fmt.Println("------------------------------------------------------------")
	}
}

func Test3Sum(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 1, 0, -1, -2}, {-1, 0, 1, 2}, {1, -1, 9, -8, 0}, {9, 8, 7, 6, 5, 5, 4, 3, 2, 1}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", ThreeSum(&v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestKthDiff(t *testing.T) {
	var TestArr []int = []int{1, 3, 5, 7, 10}
	var kArr []int = []int{2, 3, 1}
	for _, v := range kArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", Kdiffpairs(&TestArr, v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestKElement(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 2, 3, 4}, {2, 3, 4, 5}}, {{1, 1, 2, 3}, {3, 3, 4, 5, 6}}}
	var k []int = []int{3, 5}
	for i, v := range TestArr {
		fmt.Println("Test case: ", v)
		C := Kthelementoftwosortedlists(&v[0], &v[1], k[i])
		fmt.Println("Sorted array: ", C)
		fmt.Println("------------------------------------------------------------")
	}
}

func TestSortedArrayIntersection(t *testing.T) {
	var TestArr [][][]int = [][][]int{{{1, 3, 4, 5, 5, 6, 6, 7}, {2, 5, 6, 6, 7, 8}}, {{1, 1, 2, 3}, {3, 3, 4, 5, 6}}, {{1, 1, 3, 3}, {3, 3, 4, 5, 6}}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		C := SortedArraysIntersection(&v[0], &v[1])
		fmt.Println("Sorted array: ", C)
		fmt.Println("------------------------------------------------------------")
	}
}

func TestDutchFlag(t *testing.T) {
	var TestArr [][]int = [][]int{{2, 2, 0, 1}, {1, 0, 1, 2, 2}, {1, 0, 0, 0}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", DutchNationalFlag(&v))

		fmt.Println("------------------------------------------------------------")
	}
}

func TestTrappingWater(t *testing.T) {
	var TestArr [][]int = [][]int{{1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, {1, 0, 0, 1, 0}, {1, 0, 2, 0, 0, 1}, {3, 0, 2, 0, 4}}
	for _, v := range TestArr {
		fmt.Println("Test case: ", v)
		fmt.Println("Result: ", TrappingRainWater(&v))

		fmt.Println("------------------------------------------------------------")
	}
}
