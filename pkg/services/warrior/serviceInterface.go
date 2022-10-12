package warrior

import "apirest/model"

type ServiceInterface interface {
	GetAll() ([]*model.Warrior, error)
	CreateOne(id, name string, race model.Race, gender model.Gender, power int, origin string) *model.Warrior
	AddSome(warriors []*model.Warrior) int
	Add(warrior model.Warrior) error
	FindByRace(race string) []*model.Warrior
	FindById(id string) *model.Warrior
}
