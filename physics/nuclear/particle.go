package nuclear

type particle struct {
	Mass   float64
	Charge float64
}

type fundamentalParticle struct {
	particle
	Spin float64
}

type nucleon struct {
	particle
	NucleonType string
	Quarks      [3]quark
}

type atom struct {
	particle
	AtomType string
	Protons  int
	Neutrons int
	HalfLife float64
}
