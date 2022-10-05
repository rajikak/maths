package counting

// Euclid, gcd
func Euclid(a, b int) int {
	for a != b {
		if b < a {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}  

// GreatestCommonMeasure https://en.wikipedia.org/wiki/Greatest_common_divisor
func GreatestCommonMeasure(a, b int) int {

	if a == b {
		return a
	} else if b < a {
		return GreatestCommonMeasure(a - b, b)
	} else {
		return GreatestCommonMeasure(a, b - a)
	}
}

// Multiply https://en.wikipedia.org/wiki/Ancient_Egyptian_multiplication
func Multiply(n, a int) int {
	if n == 1 {
		return a
	}
	result := Multiply(half(n), a+a)
	if odd(n) {
		result = result + a
	}
	return result
}

func odd(n int) bool {
	return n&1 == 1
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
		return Multiply1(r+a, half(n), a+a)
	}
	return Multiply1(r, half(n), a+a)
}

// Multiply2, tail-recursive
func Multiply2(r, n, a int) int {
	if odd(n) {
		r = r + a
		if n == 1 {
			return r
		}
	}
	return Multiply2(r, half(n), a+a)
}

// OblongNumber, n th oblong number
func OblongNumber(n int) int {
	return n * (n + 1)
}

// Sieve of Eratosthenes
// https://cp-algorithms.com/algebra/sieve-of-eratosthenes.html
func Sieve(n int) []bool{
	var primes = make([]bool, n)
	
	for i := 2; i < n; i++ {
		primes[i] = true 
	}

	for i := 2; i < n; i++ {
		if primes[i] && (i * i <= n) {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	return primes
}
