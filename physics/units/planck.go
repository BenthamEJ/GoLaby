package units

import (
	"github.com/BenthamEJ/goLaby/physics"
	"math"
)

func PlanckTimeInSeconds() float64 {
	calc := math.Pow(physics.PlanckBar()*physics.BigG()/math.Pow(physics.LittleC(), 5), 0.5)
	return calc
}

func Second(s float64) float64 {
	// float64 accuracy
	return math.Pow(PlanckTimeInSeconds(), -1)
}
