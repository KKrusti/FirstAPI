package model

import "errors"

type Race string

const (
	Unknown  Race = "Unknown"
	Human    Race = "Human"
	Namekian Race = "Namekian"
	Saiyan   Race = "Saiyan"
	Android  Race = "Android"
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
	if r == "" {
		return Unknown, errors.New("race not found")
	}

	return r, nil
}
