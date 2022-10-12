package db

import (
	"apirest/model"
	"errors"
	"github.com/hashicorp/go-memdb"
)

type MemDbService struct {
	db *memdb.MemDB
}

func New() MemDbService {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"warrior": {
				Name: "warrior",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"race": {
						Name:    "race",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Race"},
					},
					"gender": {
						Name:    "gender",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Gender"},
					},
					"power": {
						Name:    "power",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Power"},
					},
					"origin": {
						Name:    "origin",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Origin"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return MemDbService{db: db}
}

func (srv MemDbService) Insert(warrior *model.Warrior) (*model.Warrior, error) {
	tx := srv.db.Txn(true)
	if err := tx.Insert("warrior", warrior); err != nil {
		return &model.Warrior{}, errors.New("occurred while inserting warrior")
	}
	tx.Commit()
	return warrior, nil
}

func (srv MemDbService) GetAll() ([]*model.Warrior, error) {
	var warriors []*model.Warrior
	tx := srv.db.Txn(false)
	it, err := tx.Get("warrior", "id")
	if err != nil {
		return []*model.Warrior{}, errors.New("no warriors found")
	}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		w := obj.(*model.Warrior)
		warriors = append(warriors, w)
	}
	return warriors, nil
}

func (srv MemDbService) FindByRace(race model.Race) []*model.Warrior {
	var warriors []*model.Warrior
	tx := srv.db.Txn(false)
	it, err := tx.Get("warrior", "race", string(race))
	if err != nil {
		return warriors
	}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		w := obj.(*model.Warrior)
		warriors = append(warriors, w)
	}
	return warriors
}

func (srv MemDbService) FindById(id string) *model.Warrior {
	tx := srv.db.Txn(false)
	war, err := tx.First("warrior", "id", id)
	if err != nil {
		return nil
	}
	warrior := war.(*model.Warrior)
	return warrior
}
