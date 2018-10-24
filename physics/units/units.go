package units

type unit struct {
	magnitude  float64
	dimensions [7]float64
}

type scalar struct {
	unit
}

type vector struct {
	unit
	direction [3]float64
}

// Time
//// here there would be property setting functions rather than new types

// Distance
