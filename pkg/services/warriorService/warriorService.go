package warriorService

import (
	"apirest/model"
)

type WarriorService struct {
	warriors []model.Warrior
}

func NewWarriorService() WarriorService {
	warriors := []model.Warrior{
		{ID: "1", Name: "Goku", Race: "Saiyan", Gender: "Male", Power: 50000, Origin: "Earth"},
		{ID: "2", Name: "Vegeta", Race: "Saiyan", Gender: "Male", Power: 45000, Origin: "Planet Vegeta"},
		{ID: "3", Name: "Krilin", Race: "Human", Gender: "Male", Power: 1000, Origin: "Earth"},
		{ID: "4", Name: "A18", Race: "Android", Gender: "Female", Power: 35000, Origin: "Earth"}}

	return WarriorService{warriors: warriors}
}

func (srv WarriorService) FindById(id string) model.Warrior {
	for _, value := range srv.warriors {
		if value.ID == id {
			return value
		}
	}
	return model.Warrior{}
}

func (srv WarriorService) FindAll() []model.Warrior {
	return srv.warriors
}

func (srv WarriorService) FindByRace(race string) []model.Warrior {
	found := make([]model.Warrior, 0)
	for _, value := range srv.warriors {
		if value.Race == race {
			found = append(found, value)
		}
	}
	return found
}

func (srv WarriorService) SetWarriors(warriors []model.Warrior) {
	srv.warriors = warriors
}
