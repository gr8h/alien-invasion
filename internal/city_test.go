package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCityTrapped_ShouldBeNo_WhenMapIsInitiated(t *testing.T) {

	const N = 0

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(world.Cities))

	for _, c := range world.Cities {

		assert.Equal(t, true, c.IsAlive())
		assert.Equal(t, 0, len(c.Aliens))
		hasTrappedAliens, err := c.HasTrappedAliens()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, false, hasTrappedAliens)
	}

	//assert.EqualError(t, err, "")
}

func TestCityAlien_ShouldHaveTwo_WhenMapIsInitiatedWithTwoAliens(t *testing.T) {

	const N = 2

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(world.Cities))
	assert.Equal(t, 2, len(world.Aliens))

	var aliensInCityCount int = 0
	for _, c := range world.Cities {

		aliensInCityCount += len(c.Aliens)
	}

	assert.Equal(t, 2, aliensInCityCount)
}

func TestCityDestroy_ShouldBeDestroyed_WhenTwoAliensInCity(t *testing.T) {

	const N = 3

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(world.Cities))
	assert.Equal(t, 3, len(world.Aliens))

	hasAliveCity := world.Evaluate()
	assert.Equal(t, false, hasAliveCity)

	for _, c := range world.Cities {

		if len(c.Aliens) >= 2 {
			assert.Equal(t, false, c.IsAlive())
		}
	}
}
