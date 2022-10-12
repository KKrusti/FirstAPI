package warrior

import (
	"apirest/model"
	"apirest/pkg/services/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_create_warrior(t *testing.T) {
	srv := New(db.New())
	_ = srv.CreateOne("1", "Son Goku", model.Saiyan, model.Male, 50000, "Earth")
	warriors, err := srv.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 1, len(warriors))
}

func Test_getAll_with_no_warriors(t *testing.T) {
	srv := New(db.New())
	warriors, err := srv.GetAll()
	require.NoError(t, err)
	require.Equal(t, 0, len(warriors))
}

func Test_getAll_after_filling(t *testing.T) {
	srv := New(db.New())
	addedWarriors := populateDb(srv)
	assert.Equal(t, 4, addedWarriors)
	warriors, err := srv.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 4, len(warriors))
}

func Test_find_by_race(t *testing.T) {
	srv := New(db.New())
	addedWarriors := populateDb(srv)
	assert.Equal(t, 4, addedWarriors)
	warriors := srv.repository.FindByRace(model.Android)
	assert.Equal(t, 1, len(warriors))
}

func Test_find_by_id(t *testing.T) {
	srv := New(db.New())
	addedWarriors := populateDb(srv)
	assert.Equal(t, 4, addedWarriors)
	warrior := srv.repository.FindById("1")
	expected := model.Warrior{ID: "1", Name: "Goku", Race: model.Saiyan, Gender: model.Male, Power: 50000, Origin: "Earth"}
	assert.Equal(t, &expected, warrior)
}

func populateDb(srv Service) int {
	warriors := []*model.Warrior{
		{ID: "1", Name: "Goku", Race: model.Saiyan, Gender: model.Male, Power: 50000, Origin: "Earth"},
		{ID: "2", Name: "Vegeta", Race: model.Saiyan, Gender: model.Male, Power: 45000, Origin: "Planet Vegeta"},
		{ID: "3", Name: "Krilin", Race: model.Human, Gender: model.Male, Power: 1000, Origin: "Earth"},
		{ID: "4", Name: "A18", Race: model.Android, Gender: model.Female, Power: 35000, Origin: "Earth"}}
	return srv.AddSome(warriors)
}

//findById
