package cosmic

import (
	"github.com/BenthamEJ/goLaby/physics"
	"math"
)

// Hawking Radiation Temperature
func TempH(mass float64) float64 {
	var numerator = physics.PlanckBar() * physics.LittleC() * physics.LittleC() * physics.LittleC()
	var denominator = 8 * math.Pi * physics.BigG() * mass * physics.BoltzmannConst()
	return numerator / denominator
}
