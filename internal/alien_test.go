package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlienDestroy_ShouldBeDead_WhenDestroyIsCalled(t *testing.T) {

	var alien Alien = NewAlien(1)

	alien.Destroy()

	assert.Equal(t, false, alien.IsAlive())
}

func TestAlienDestroy_ShouldThrougError_WhenAlienIsAlreadyDead(t *testing.T) {

	var alien Alien = NewAlien(1)
	alien.Destroy()
	err := alien.Destroy()
	assert.Equal(t, false, alien.IsAlive())
	assert.EqualError(t, err, "Aline:Destroy - Already dead 1")
}

func TestAlienCanMove_ShouldThrougError_WhenAlienIHasNoCity(t *testing.T) {

	var alien Alien = NewAlien(1)
	canMove, err := alien.CanMove()
	assert.Equal(t, false, canMove)
	assert.EqualError(t, err, "Aline:CanMove - No city found 1")
	err = alien.Move()
	assert.EqualError(t, err, "Aline:Move - No city found 1")
}
func TestAlienCanMove_ShouldReturnTrue_WhenAvailableConnection(t *testing.T) {

	const N = 1

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(world.Aliens))

	var aline Alien = world.Aliens[0]

	assert.NotEqual(t, nil, aline.city)

	canMove, err := aline.CanMove()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, canMove)
}

func TestAlienMove_ShouldReturnNewCity_WhenAvailableConnection(t *testing.T) {

	const N = 1

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(world.Aliens))

	var aline Alien = world.Aliens[0]

	assert.NotEqual(t, nil, aline.city)

	var oldCity City = *aline.city

	err = aline.Move()
	if err != nil {
		t.Error(err)
	}

	var newCity City = *aline.city

	assert.NotEqual(t, oldCity, newCity)
}
