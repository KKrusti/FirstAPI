package warrior

import (
	"apirest/model"
	"apirest/pkg/services/db"
)

type Service struct {
	memDB    db.MemDbService
	warriors *[]model.Warrior
}

func New() Service {
	return Service{warriors: &[]model.Warrior{},
		memDB: db.New()}
}

func (srv Service) CreateOne(id, name string, race model.Race, gender model.Gender, power int, origin string) *model.Warrior {
	warrior := &model.Warrior{
		ID:     id,
		Name:   name,
		Race:   race,
		Gender: gender,
		Power:  power,
		Origin: origin,
	}
	war, _ := srv.memDB.InsertWarrior(warrior)
	return war
}

func (srv Service) GetAll() ([]*model.Warrior, error) {
	warriors, err := srv.memDB.GetAllWarriors()
	if err != nil {
		return nil, err
	} else {
		return warriors, nil
	}
}

func (srv Service) AddSome(warriors []model.Warrior) int {
	for _, warrior := range warriors {
		*srv.warriors = append(*srv.warriors, warrior)
	}

	return len(warriors)
}

func (srv Service) Add(warrior model.Warrior) {
	*srv.warriors = append(*srv.warriors, warrior)
}

func (srv Service) FindByRace(raceString string) []model.Warrior {
	var foundWarriors []model.Warrior
	race, err := model.ParseRace(raceString)
	if err != nil {
		return foundWarriors
	}

	for _, warrior := range *srv.warriors {
		if warrior.Race == race {
			foundWarriors = append(foundWarriors, warrior)
		}
	}
	return foundWarriors
}

func (srv Service) FindById(id string) model.Warrior {
	for _, warrior := range *srv.warriors {
		if warrior.ID == id {
			return warrior
		}
	}
	return model.Warrior{}
}
