package warriorService

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarriorService_FindById(t *testing.T) {
	srv := NewWarriorService()
	got := srv.FindById("2")
	warriors := *srv.warriors
	expected := warriors[1]
	assert.Equal(t, expected, got, "The two words should be the same.")
}
