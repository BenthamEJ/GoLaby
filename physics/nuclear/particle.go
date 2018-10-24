package nuclear

type particle struct {
	mass   float64
	charge float64
}

type fundamentalParticle struct {
	particle
	spin float64
}

type nucleon struct {
	particle
	nucleonType string
	quarks      [3]quark
}

type atom struct {
	particle
	atomType string
	protons  int
	neutrons int
	halfLife float64
}
