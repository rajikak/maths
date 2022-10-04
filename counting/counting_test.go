package counting

import "testing"


// go test -test.v
func TestSieve(t *testing.T) {

	n := 50
	p := Sieve(n)

	for i, isPrime := range p {
		t.Logf("%d , %v", i, isPrime)
	}
}

func BenchmarkSieve(b *testing.B) {

	n := 1000
	for i:= 0; i < b.N; i++ {
		Sieve(n)
	}
}


func TestMultiply(t *testing.T) {
	type args struct {
		n int
		a int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{2, 3}, 6},
		{args{2, 5}, 10},
		{args{3, 5}, 15},
	}
	for _, test := range tests {
		if got := Multiply(test.args.n, test.args.a); got != test.want {
			t.Errorf("Multiply(%v, %v) = %v, want = %v", test.args.n, test.args.a, got, test.want)
		}
	}
}

func TestMultiply1And2(t *testing.T) {
	type args struct {
		n int
		a int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{2, 3}, 6},
		{args{2, 5}, 10},
		{args{3, 5}, 15},
	}
	r := 0
	for _, test := range tests {
		if got := Multiply1(r, test.args.n, test.args.a); got != test.want {
			t.Errorf("Multiply1(%v, %v, %v) = %v, want = %v", r, test.args.n, test.args.a, got, test.want)
		}
		if got := Multiply2(r, test.args.n, test.args.a); got != test.want {
			t.Errorf("Multiply2(%v, %v, %v) = %v, want = %v", r, test.args.n, test.args.a, got, test.want)
		}
	}
}

func TestOblongNumber(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 2},
		{2, 6},
		{3, 12},
		{4, 20},
		{5, 30},
		{6, 42},
		{7, 56},
	}

	for _, test := range tests {
		if got := OblongNumber(test.n); got != test.want {
			t.Errorf("OblongNumber(%v) = %v, want = %v", test.n, got, test.want)
		}

	}
}

