package math

import "testing"

func BenchmarkFastSegmentRemainder(b *testing.B) {
	n := 1000
	for i := 0; i < n; i++ {
		FastSegmentRemainder(20, 190)
	}
}

func BenchmarkFastSegmentRemainder2(b *testing.B) {
	n := 1000
	for i := 0; i < n; i++ {
		FastSegmentRemainder2(20, 190)
	}
}

func BenchmarkGreatestCommonMeasure(b *testing.B) {

	n := 1000
	for i := 0; i < n; i++ {
		GreatestCommonMeasure(196, 42)
	}
}

func BenchmarkEuclid(b *testing.B) {
	n := 1000
	for i := 0; i < n; i++ {
		Euclid(196, 42)
	}
}

func TestFastSegmentRemainder(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{196, 42}, 28},
		{args{196, 200}, 196},
	}

	for _, test := range tests {
		if got := FastSegmentRemainder(test.args.a, test.args.b); got != test.want {
			t.Errorf("FastSegmentRemainder gcd(%v, %v) = %v, want = %v", test.args.a, test.args.b, got, test.want)
		}

		if got := FastSegmentRemainder2(test.args.a, test.args.b); got != test.want {
			t.Errorf("FastSegmentRemainder2 gcd(%d, %d) = %d, want = %d", test.args.a, test.args.b, got, test.want)
		}
	}
}

func TestGCM(t *testing.T) {

	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{196, 42}, 14},
		{args{196, 200}, 4},
		{args{25, 200}, 25},
		{args{200, 1200}, 200},
	}

	for _, test := range tests {
		if got := GreatestCommonMeasure(test.args.a, test.args.b); got != test.want {
			t.Errorf("gcd(%v, %v) = %v, want = %v", test.args.a, test.args.b, got, test.want)
		}
		if got := Euclid(test.args.a, test.args.b); got != test.want {
			t.Errorf("Euclid's gcd(%v, %v) = %v, want = %v", test.args.a, test.args.b, got, test.want)
		}
	}
}

// go test -test.v
func TestSieve(t *testing.T) {

	n := 1000
	p := Sieve(n)

	for i, isPrime := range p {
		if isPrime {
			t.Logf("%d", i)
		}
	}
}

func BenchmarkSieve(b *testing.B) {

	n := 1000
	for i := 0; i < b.N; i++ {
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
