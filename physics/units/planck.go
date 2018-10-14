package units

import "github.com/BenthamEJ/goLaby/maths"

func PlanckTimeInSeconds() float64 {
	base := 5.39
	return base * maths.TenToThe(-44)
}
