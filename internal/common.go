package internal

import (
	"alien-invasion/pkg/helper"
	"path/filepath"
)

// Defaults
var oppositeDirection = map[string]string{
	"north": "south",
	"east":  "west",
	"south": "north",
	"west":  "east",
}

const validWorldSmallPath = "valid_map_small.txt"
const trappedWorldSmallPath = "trapped_map_small.txt"
const extraDirection = "invalid_connection_extra_direction.txt"
const missingDirection = "invalid_connection_missing_direction.txt"
const wrongDirection = "invalid_connection_wrong_direction.txt"

const showExtraMessages = false

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func setTestWorld(N int, worldFilePath string) (World, error) {

	path, err := filepath.Abs("../test/" + worldFilePath)

	simpleWorldMap, err := helper.ReadWorldMapFile(path)

	var world World = NewWorld()

	err = world.ValidateMap(simpleWorldMap)
	if err != nil {
		return world, err
	}

	err = world.Construct(simpleWorldMap)
	if err != nil {
		return world, err
	}

	err = world.InhabitAlien(int(N))
	if err != nil {
		return world, err
	}

	return world, err
}
