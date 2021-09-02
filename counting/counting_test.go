package counting

import "testing"

func TestMultiplyBy(t *testing.T) {
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
	}
	for _, test := range tests {
		if got := MultiplyBy(test.args.n, test.args.a); got != test.want {
			t.Errorf("MultiplyBy() = %v, want %v", got, test.want)
		}
	}
}
