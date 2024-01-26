package pat

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	return a * b / Gcd(a, b)
}

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func Determinant(mk [][]float64) float64 {
	for i := 0; i < len(mk); i++ {
		if len(mk) != len(mk[i]) {
			panic("not a square matrix")
		}
	}
	m := make([][]float64, len(mk))
	for i := 0; i < len(mk); i++ {
		m[i] = make([]float64, len(mk[i]))
		copy(m[i], mk[i])
	}
	r := 1.0
	for j := 0; j < len(m)-1; j++ {
		if m[j][j] == 0 {
			b := true
			for i := j + 1; i < len(m); i++ {
				if m[i][j] != 0 {
					m[j], m[i] = m[i], m[j]
					r *= -1
					b = false
					break
				}
			}
			if b {
				return 0
			}
		}
		for i := j + 1; i < len(m); i++ {
			d := m[i][j] / m[j][j]
			for k := j; k < len(m); k++ {
				m[i][k] -= m[j][k] * d
			}
		}
		r *= m[j][j]
	}
	return r * m[len(m)-1][len(m)-1]
}
