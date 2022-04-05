package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const validWorldSmallPath = "valid_map_small.txt"
const trappedWorldSmallPath = "trapped_map_small.txt"

func TestDestroy_ShouldBeDead_WhenDestroyIsCalled(t *testing.T) {

	var alien Alien = NewAlien(1)

	alien.Destroy()

	if alien.IsAlive() {
		t.Errorf("Alien %d, should be dead!", alien.Id)
	}
}

func TestCanMove_ShouldReturnTrue_WhenAvailableConnection(t *testing.T) {

	const N = 1

	var world World = setTestWorld(N, validWorldSmallPath, t)

	assert.Equal(t, len(world.Aliens), 1)

	var aline Alien = world.Aliens[0]

	assert.NotEqual(t, aline.city, nil)

	canMove, err := aline.CanMove()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, canMove, true)
}

func TestMove_ShouldReturnNewCity_WhenAvailableConnection(t *testing.T) {

	const N = 1

	var world World = setTestWorld(N, validWorldSmallPath, t)

	assert.Equal(t, len(world.Aliens), 1)

	var aline Alien = world.Aliens[0]

	assert.NotEqual(t, aline.city, nil)

	var oldCity City = *aline.city

	err := aline.Move()
	if err != nil {
		t.Error(err)
	}

	var newCity City = *aline.city

	assert.NotEqual(t, oldCity, newCity)
}

func TestMove_ShouldReturnSameCity_WhenNoAvailableConnection(t *testing.T) {

	const N = 1

	var world World = setTestWorld(N, trappedWorldSmallPath, t)

	assert.Equal(t, len(world.Aliens), 1)

	var aline Alien = world.Aliens[0]

	assert.NotEqual(t, aline.city, nil)

	aline.city.Destroy()
	var oldCity City = *aline.city

	err := aline.Move()
	if err != nil {
		t.Error(err)
	}

	var newCity City = *aline.city

	assert.Equal(t, oldCity, newCity)
}
