package main

import (
	"github.com/BenthamEJ/goLaby/physics"
	"github.com/BenthamEJ/goLaby/physics/cosmic"
)

func main() {
	print("Console works without being called main.\n\n")

	println("Question: What is the Hawking Radiation Temperature of a 1 solar mass black hole?")
	answer := cosmic.TempH(physics.SolarMass())
	println("Answer:", answer)
}
