package polyplus

type polyCo struct {
	coeffs []float64
	symbol string
}

func Factorial(input int64) int64 {
	var output int64 = 1
	if input == 0 {
		return output
	}
	if (input > 0) && (input <= 20) {
		for i := input; i >= 1; i-- {
			output *= i
		}
		return output
	} else {
		// Use big ints or print "invalid" message
		return 0
	}
}

func Binomial(n int64, k int64) float64 {
	numerator := Factorial(n)
	denominator := Factorial(k) * Factorial(n-k)
	return float64(numerator) / float64(denominator)
}
