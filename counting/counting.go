package counting

// MultiplyBy
// a > 0, n > 0
//    1.a = a
// (n+1)a = na + a
func MultiplyBy(n, a int) int {
	if n == 1 {
		return a
	}
	return MultiplyBy(n-1, a) + a
}
