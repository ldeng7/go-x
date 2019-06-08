package math

func Gcd(a, b int) int {
	if a != 0 && b != 0 {
		ea, eb := a&1 == 0, b&1 == 0
		if ea && eb {
			return Gcd(a>>1, b>>1) << 1
		} else if ea {
			return Gcd(a>>1, b)
		} else if eb {
			return Gcd(a, b>>1)
		} else if a <= b {
			return Gcd(b-a, a)
		} else {
			return Gcd(a-b, b)
		}
	} else if a == 0 {
		return b
	} else {
		return a
	}
}

func P(m, n int) int {
	out := n
	for i := n - 1; i > n-m; i-- {
		out *= i
	}
	return out
}

func C(m, n int) int {
	if n-m > m {
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
