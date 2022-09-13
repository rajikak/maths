package counting

// MultiplyBy https://en.wikipedia.org/wiki/Ancient_Egyptian_multiplication
func MultiplyBy(n, a int) int {
	if n == 1 {
		return a
	}
	result := MultiplyBy(half(n), a + a)
	if odd(n) {
		result = result + a
	}
	return result
}

func odd(n int) bool {
	return n & 1 == 1
}

func half(n int) int {
	return n >> 1
}
