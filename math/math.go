package math

func Gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return Gcd(b, a%b)
}

func P(m, n int) int {
	out := n
	for i := n - 1; i > n-m; i-- {
		out *= i
	}
	return out
}

func C(m, n int) int {
	if n-m < m {
		m = n - m
	}
	out := n
	for i := n - 1; i > n-m; i-- {
		out *= i
	}
	d := m
	for i := m - 1; i >= 2; i-- {
		d *= i
	}
	return out / d
}
