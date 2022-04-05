package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection_ShouldBeEven_WhenCitiesAreConnected(t *testing.T) {
	const N = 0

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(world.Cities))

	var connectionCount int = 0
	for _, city := range world.Cities {
		connectionCount += len(city.Connections)
	}

	assert.Equal(t, 8, connectionCount)
}

func TestConnection_ShouldBeValid_WhenCitiesAreConnected(t *testing.T) {
	const N = 0

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	for _, city := range world.Cities {
		for _, con := range city.Connections {
			assert.Equal(t, true, con.IsAlive())
		}
	}
}

func TestConnection_ShouldBeNotValid_WhenAllCitiesAreDestroyed(t *testing.T) {
	const N = 10

	world, err := setTestWorld(N, validWorldSmallPath)
	if err != nil {
		t.Error(err)
	}

	world.Evaluate()
	world.MoveAliens()

	for _, city := range world.Cities {
		for _, con := range city.Connections {
			assert.Equal(t, false, con.IsAlive())
		}
	}
}
