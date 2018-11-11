package polyplus

import "math"

// Single variable polynomial
type Poly1 struct {
	Coeffs []float64
	Order  int
}

func (p Poly1) Initialise() {
	p.Order = len(p.Coeffs) - 1
}

func (p Poly1) SolveAsFunction(input float64) float64 {
	var answer float64
	for i := 0; i <= p.Order; i++ {
		answer += p.Coeffs[i] * math.Pow(input, float64(i))
	}
	return answer
}

// Factorial, Binomial
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
