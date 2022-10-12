package warrior

import (
	"apirest/model"
	"errors"
)

type Service struct {
	repository Repository
}

func New(repository Repository) Service {
	return Service{
		repository: repository,
	}
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
	war, _ := srv.repository.Insert(warrior)
	return war
}

func (srv Service) Add(warrior model.Warrior) error {
	_, err := srv.repository.Insert(&warrior)
	if err != nil {
		return errors.New("storing warrior into db")
	}
	return nil
}

func (srv Service) GetAll() ([]*model.Warrior, error) {
	warriors, err := srv.repository.GetAll()
	if err != nil {
		return nil, err
	} else {
		return warriors, nil
	}
}

func (srv Service) AddSome(warriors []*model.Warrior) int {
	//TODO think on correct error handling
	for _, warrior := range warriors {
		_, err := srv.repository.Insert(warrior)
		if err != nil {
			continue
		}
	}

	return len(warriors)
}

func (srv Service) FindByRace(raceString string) []*model.Warrior {
	var foundWarriors []*model.Warrior
	race, err := model.ParseRace(raceString)
	if err != nil {
		return foundWarriors
	}

	warriors := srv.repository.FindByRace(race)

	for _, warrior := range warriors {
		if warrior.Race == race {
			foundWarriors = append(foundWarriors, warrior)
		}
	}
	return foundWarriors
}

func (srv Service) FindById(id string) *model.Warrior {
	return srv.repository.FindById(id)
}
