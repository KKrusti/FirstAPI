package warrior

import (
	"apirest/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_create_warrior(t *testing.T) {
	srv := New()
	assert.Equal(t, 0, len(*srv.warriors))
	_ = srv.CreateOne("1", "Son Goku", model.Saiyan, model.Male, 50000, "Earth")
	warriors, err := srv.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 1, len(warriors))
}

func Test_getAll_with_no_warriors(t *testing.T) {
	srv := New()
	warriors, err := srv.GetAll()
	require.Error(t, err, "no warriors found")
	require.Equal(t, 0, len(warriors))
}

func Test_getAll_after_filling(t *testing.T) {
	srv := New()
	addedWarriors := populateDb(srv)
	assert.Equal(t, 4, addedWarriors)
	warriors, err := srv.GetAll()
	require.NoError(t, err)
	assert.Equal(t, 4, len(warriors))
}

func populateDb(srv Service) int {
	warriors := []model.Warrior{
		{ID: "1", Name: "Goku", Race: model.Saiyan, Gender: model.Male, Power: 50000, Origin: "Earth"},
		{ID: "2", Name: "Vegeta", Race: model.Saiyan, Gender: model.Male, Power: 45000, Origin: "Planet Vegeta"},
		{ID: "3", Name: "Krilin", Race: model.Human, Gender: model.Male, Power: 1000, Origin: "Earth"},
		{ID: "4", Name: "A18", Race: model.Android, Gender: model.Female, Power: 35000, Origin: "Earth"}}
	return srv.AddSome(warriors)
}
