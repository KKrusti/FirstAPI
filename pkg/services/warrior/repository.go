package warrior

import "apirest/model"

type Repository interface {
	Insert(warrior *model.Warrior) (*model.Warrior, error)
	GetAll() ([]*model.Warrior, error)
	FindByRace(race model.Race) []*model.Warrior
	FindById(id string) *model.Warrior
}
