package problems

import (
	"fmt"
	"math"
)

func LongestSubarraywithZeroSum(a *[]int) {
	m := make(map[int]int)
	mlen := 0
	sum := 0

	for i, v := range *a {
		sum += v

		if _, ok := m[sum]; ok {
			mlen = int(math.Max(float64(mlen), float64(i-m[sum])))
		} else {
			m[sum] = i
		}
	}
	fmt.Println("--->", mlen)
}
