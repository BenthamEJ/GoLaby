package physics

import (
	"github.com/BenthamEJ/goLaby/maths"
	"math"
)

func LittleC() float64 {
	return 299792458
}

func BigG() float64 {
	var base = 6.67408
	return base * maths.TenToThe(-11)
}

func PlanckConst() float64 {
	var base = 6.62607004
	return base * maths.TenToThe(-34)
}

func PlanckBar() float64 {
	var base = PlanckConst()
	return base / (2 * math.Pi)
}

func BoltzmannConst() float64 {
	var base = 1.380649
	return base * maths.TenToThe(-23)
}
