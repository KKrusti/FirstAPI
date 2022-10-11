package model

type Warrior struct {
	ID     string `json:"ID"`
	Name   string `json:"name,omitempty"`
	Race   Race   `json:"race,omitempty"`
	Gender Gender `json:"gender,omitempty"`
	Power  int    `json:"power,omitempty"`
	Origin string `json:"origin,omitempty"`
}
