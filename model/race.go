package model

import "errors"

type Race int

const (
	Unknown Race = iota
	Human
	Namekian
	Saiyan
	Android
)

var (
	racesMap = map[string]Race{
		"Human":    Human,
		"Namekian": Namekian,
		"Saiyan":   Saiyan,
		"Android":  Android,
	}
)

func ParseRace(s string) (race Race, err error) {
	r := racesMap[s]
	if r == 0 {
		return Unknown, errors.New("race not found")
	}

	return r, nil
}
