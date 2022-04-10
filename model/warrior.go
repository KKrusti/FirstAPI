package model

type Warrior struct {
	ID     string `json:"ID"`
	Name   string `json:"name,omitempty"`
	Race   string `json:"race,omitempty"`
	Gender string `json:"gender,omitempty"`
	Power  int    `json:"power,omitempty"`
	Origin string `json:"origin,omitempty"`
}
