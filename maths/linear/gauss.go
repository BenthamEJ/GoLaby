package linear

import "math"

// Based on interpreting the C version, assuming fabs = math.Abs
// A "void" in C is represented by nothing in go.
func Normal(a [100][100]float64, b [100]float64, m int32) {
	var i, j int32
	var big float64

	for i = 0; i < m; i++ {
		big = 0.0

		for j = 0; j < m; j++ {
			if big < math.Abs(a[i][j]) {
				big = math.Abs(a[i][j])
			}
		}

		for j = 0; j < m; j++ {
			a[i][j] = a[i][j] / big
		}

		b[i] = b[i] / big
	}
}
