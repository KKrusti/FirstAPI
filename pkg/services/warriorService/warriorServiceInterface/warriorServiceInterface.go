package warriorServiceInterface

import "apirest/model"

type WarriorServiceInterface interface {
	FindById(id string) model.Warrior
	FindAll() []model.Warrior
	FindByRace(race string) []model.Warrior
	SetWarriors([]model.Warrior)
}
