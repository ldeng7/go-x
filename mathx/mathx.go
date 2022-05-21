package mathx

import "github.com/ldeng7/go-x/common"

func Gcd[T common.UintNumber](a, b T) T {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func P[T common.UintNumber](m, n T) T {
	var o T = 1
	for i := n; i > n-m; i-- {
		o *= i
	}
	return o
}

func C[T common.UintNumber](m, n T) T {
	if n-m < m {
		m = n - m
	}
	var o T = 1
	for i := n; i > n-m; i-- {
		o *= i
	}
	var d T = 1
	for i := m; i > 1; i-- {
		d *= i
	}
	return o / d
}
