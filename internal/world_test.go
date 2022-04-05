package internal

import (
	"alien-invasion/pkg/helper"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_ShouldBeOpposite_WhenTwoCitisAreConnected(t *testing.T) {
	const N = 0

	path, err := filepath.Abs("../test/" + wrongDirection)
	if err != nil {
		t.Error(err)
	}

	simpleWorldMap, err := helper.ReadWorldMapFile(path)
	if err != nil {
		t.Error(err)
	}

	var world World = NewWorld()

	err = world.ValidateMap(simpleWorldMap)
	assert.EqualError(t, err, "ValidateMap: Wrong Direction")
}

func TestMap_ShouldBeNotMoreThanTwo_WhenTwoCitisAreConnected(t *testing.T) {
	const N = 0

	path, err := filepath.Abs("../test/" + extraDirection)
	if err != nil {
		t.Error(err)
	}

	simpleWorldMap, err := helper.ReadWorldMapFile(path)
	if err != nil {
		t.Error(err)
	}

	var world World = NewWorld()

	err = world.ValidateMap(simpleWorldMap)
	assert.EqualError(t, err, "ValidateMap: Extra Connection")
}

func TestMap_ShouldBeEven_WhenTwoCitisAreConnected(t *testing.T) {
	const N = 0

	path, err := filepath.Abs("../test/" + missingDirection)
	if err != nil {
		t.Error(err)
	}

	simpleWorldMap, err := helper.ReadWorldMapFile(path)
	if err != nil {
		t.Error(err)
	}

	var world World = NewWorld()

	err = world.ValidateMap(simpleWorldMap)
	assert.EqualError(t, err, "ValidateMap: Missing Connection")
}
