package math

import "testing"

func TestSolve(t *testing.T) {

	type arg struct {
		a int
		b int
		c int
	}

	tests := []struct {
		arg  arg
		want int
	}{
		{arg{a: 10, b: 10, c: 10}, 0},
		{arg{a: 10, b: 20, c: 10}, -1},
	}

	for _, test := range tests {

		if got := Solve(test.arg.a, test.arg.b, test.arg.c); got != test.want {
			t.Errorf("Slove(%d, %d, %d) = %d, want = %d", test.arg.a, test.arg.b, test.arg.c, got, test.want)
		}

	}
}
