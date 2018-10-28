package trig

import "math"

// Extended trigonometrical functions
func Csc(x float64) float64 {
	return 1 / math.Sin(x)
}

func Sec(x float64) float64 {
	return 1 / math.Cos(x)
}

func Cot(x float64) float64 {
	return 1 / math.Tan(x)
}

func Acsc(x float64) float64 {
	return 1 / math.Asin(x)
}

func Asec(x float64) float64 {
	return 1 / math.Acos(x)
}

func Acot(x float64) float64 {
	return 1 / math.Atan(x)
}

func Csch(x float64) float64 {
	return 1 / math.Sinh(x)
}

func Sech(x float64) float64 {
	return 1 / math.Cosh(x)
}

func Coth(x float64) float64 {
	return 1 / math.Tanh(x)
}

// func(x + y)
func sinh2(x, y float64) float64 {
	return math.Sinh(x)*math.Cosh(y) + math.Cosh(x)*math.Sinh(y)
}

func cosh2(x, y float64) float64 {
	return math.Cosh(x)*math.Cosh(y) + math.Sinh(x)*math.Sinh(y)
}

func tanh2(x, y float64) float64 {
	numerator := math.Tanh(x) + math.Tanh(y)
	denominator := 1 + math.Tanh(x)*math.Tanh(y)
	return numerator / denominator
}

// Alternate names
func Cosec(x float64) float64 {
	return Csc(x)
}

func Arcsin(x float64) float64 {
	return math.Asin(x)
}

func Arccos(x float64) float64 {
	return math.Acos(x)
}

func Arctan(x float64) float64 {
	return math.Atan(x)
}

func Arccsc(x float64) float64 {
	return Acsc(x)
}

func Arcsec(x float64) float64 {
	return Asec(x)
}

func Arccot(x float64) float64 {
	return Acot(x)
}

func Arcsinh(x float64) float64 {
	return math.Asinh(x)
}

func Arccosh(x float64) float64 {
	return math.Acosh(x)
}

func Arctanh(x float64) float64 {
	return math.Atanh(x)
}

func Cosech(x float64) float64 {
	return Csch(x)
}
