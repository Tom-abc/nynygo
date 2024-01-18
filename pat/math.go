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
