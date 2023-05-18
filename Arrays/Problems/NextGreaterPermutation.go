package problems

// Next Greater Permutation
// Hard
// 80
// Given an array, rearrange it to its next greater permutation. Do it in-place with extra constant memory only. Do not use any library function for the next permutation.

// Example
// Array: [1, 2, 3, 4]
// Next Greater Permutation: [1, 2, 4, 3]
// Next Greater Permutation: [1, 3, 2, 4]
// Next Greater Permutation: [1, 3, 4, 2]
// Next Greater Permutation: [1, 4, 2, 3]
// Next Greater Permutation: [1, 4, 3, 2]
// Next Greater Permutation: [2, 1, 3, 4]
// If the next greater permutation does not exist, return the lowest possible order (sorted in ascending order).

func NextGreaterPermutation(a *[]int) {
	i := 0
	x := 0
	k := 0
	for i = len(*a) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {

			if (*a)[j] < (*a)[i] {
				x = i
				k = (*a)[i]
				for x > j {
					(*a)[x] = (*a)[x-1]
					x--
				}
				(*a)[x] = k
				return
			}
		}
	}

	if i < 0 {
		MergeSort(a, 0, len(*a)-1)
	}
}
