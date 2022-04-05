package internal

import (
	"alien-invasion/pkg/helper"
	"path/filepath"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func setTestWorld(N int, worldFilePath string, t *testing.T) World {

	path, err := filepath.Abs("../test/" + worldFilePath)

	simpleWorldMap, err := helper.ReadWorldMapFile(path)
	if err != nil {
		t.Error(err)
	}

	var world World = NewWorld()

	err = world.ValidateMap(simpleWorldMap)
	if err != nil {
		t.Error(err)
	}

	err = world.Construct(simpleWorldMap)
	if err != nil {
		t.Error(err)
	}

	err = world.InhabitAlien(int(N))
	if err != nil {
		t.Error(err)
	}

	return world
}
