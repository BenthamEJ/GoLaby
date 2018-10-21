package nuclear

import "strings"

type quark struct {
	fundamentalParticle
	Colour  string
	Flavour string
	Anti    bool
}

func (q quark) SetColour(c string) {
	c = strings.ToLower(c)

	// Consistency and corrections
	if c == "antired" || c == "anti-red" {
		c = "cyan"
	}
	if c == "antigreen" || c == "anti-green" {
		c = "magenta"
	}
	if c == "antiblue" || c == "anti-blue" {
		c = "yellow"
	}

	// Validation
	if c == "red" || c == "green" || c == "blue" ||
		c == "cyan" || c == "magenta" || c == "yellow" {
		q.Colour = c
	} else {
		q.Colour = ""
	}
}

func (q quark) SetFlavour(f string) {
	f = strings.ToLower(f)

	// Consistency and corrections
	if f == "truth" || f == "t" {
		f = "top"
	}
	if f == "beauty" || f == "b" {
		f = "bottom"
	}
	if f == "c" {
		f = "charm"
	}
	if f == "s" {
		f = "strange"
	}
	if f == "d" {
		f = "down"
	}
	if f == "u" {
		f = "up"
	}

	// Validation
	if f == "top" || f == "bottom" ||
		f == "charm" || f == "strange" ||
		f == "down" || f == "up" {
		q.Flavour = f
	} else {
		q.Flavour = ""
	}
}

func (q quark) SetAnti() {
	if q.Colour == "cyan" || q.Colour == "magenta" || q.Colour == "yellow" {
		q.Anti = true
	}
}
