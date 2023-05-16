package problems

func PrimesuptoN(n int) []int {
	var B []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			B = append(B, i)
		}
	}
	return B
}

func isPrime(x int) bool {

	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
