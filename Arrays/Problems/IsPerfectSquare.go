package problems

// Given a positive integer num, write a function that returns true if num is a perfect square else false.

// Note: Do not use the in-built methods to calculate square root or power.

// Examples
// isPerfectSquare(25) => true
// isPerfectSquare(30) => false

func IsPerfectSquare(num int) bool {
	l, m, h := 0, 0, num/2

	for l <= h {
		if l*l == num {
			return true
		} else if h*h == num {
			return true
		} else {
			m = (l + h) / 2
			if m*m == num {
				return true
			} else if m*m > num {
				h = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return false
}
