package counting

// Multiply https://en.wikipedia.org/wiki/Ancient_Egyptian_multiplication
func Multiply(n, a int) int {
	if n == 1 {
		return a
	}
	result := Multiply(half(n), a + a)
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

// Multiply1 r + na, where r is a running result that accumulates the partial products na
func Multiply1(r, n, a int) int {
	if n == 1 {
		return r + a
	}
	if odd(n) {
		return Multiply1(r + a, half(n), a + a)
	} 
		return Multiply1(r, half(n), a + a)
}
