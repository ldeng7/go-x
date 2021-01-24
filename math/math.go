package math

func Gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return Gcd(b, a%b)
}

func P(m, n int) int {
	o := 1
	for i := n; i > n-m; i-- {
		o *= i
	}
	return o
}

func C(m, n int) int {
	if n-m < m {
		m = n - m
	}
	o := 1
	for i := n; i > n-m; i-- {
		o *= i
	}
	d := 1
	for i := m; i > 1; i-- {
		d *= i
	}
	return o / d
}
