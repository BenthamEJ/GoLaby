package physics

import "github.com/BenthamEJ/goLaby/maths"

func SolarMass() float64 {
	base := 1.989
	return base * maths.TenToThe(30)
}
