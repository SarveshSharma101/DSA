package problems

import (
	"fmt"
	"math"
)

func MaximumKSubarraySum(a *[]int, k int) int {
	sum := 0
	res := 0
	j := 0
	for i := 0; i <= len(*a)-k; i++ {
		j = i
		sum = 0
		for j < i+k {
			sum = sum + (*a)[j]
			j++
		}
		fmt.Println("sum---->", sum)
		res = int(math.Max(float64(res), float64(sum)))
	}

	return res
}
