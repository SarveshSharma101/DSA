package problems

// Consider an array that is divided into two parts and both of the parts are sorted individually.
//  Given the mentioned array and the end index of the first part, merge them to create a sorted array.
// Update the same array with the merged elements. You can use extra auxiliary space.

// Expected Time Complexity: O(n) where n denotes the size of the array.

// Example
// Array: [1, 3, 5, 7, 9, 11, 0, 4, 6, 8]
// End Index: 5
// Array after merging: [0, 1, 3, 4, 5, 6, 7, 8, 9, 11]

func MergeSortedSubArray(a *[]int, index int) {
	var B []int
	i := 0
	j := index + 1

	for i <= index && j < len(*a) {
		if (*a)[i] < (*a)[j] {
			B = append(B, (*a)[i])
			i++
		} else {
			B = append(B, (*a)[j])
			j++
		}
	}

	for i < len(*a) {
		B = append(B, (*a)[i])
		i++
	}

	for j < len(*a) {
		B = append(B, (*a)[j])
		j++
	}

	copy(*a, B)

}
