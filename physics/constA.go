package physics

import (
	"github.com/BenthamEJ/goLaby/maths"
	"math"
)

func LittleC() float64 {
	base := 2.99792458
	return base * maths.TenToThe(8)
}

func BigG() float64 {
	base := 6.67408
	return base * maths.TenToThe(-11)
}

func PlanckConst() float64 {
	base := 6.62607004
	return base * maths.TenToThe(-34)
}

func PlanckBar() float64 {
	base := PlanckConst()
	return base / (2 * math.Pi)
}

func BoltzmannConst() float64 {
	base := 1.380649
	return base * maths.TenToThe(-23)
}
