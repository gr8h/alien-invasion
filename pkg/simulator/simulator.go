package simulator

import (
	"alien-invasion/internal"
	"alien-invasion/pkg/helper"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
- Run the simulation
-- Read the world map from a file
-- Validate the map
-- Construct the world
-- Inhabit the aliens
-- Iterate until steps are completed
--- or until all aliens are trapped
--- or until all cities are destroyed
- Print the remaining cities
*/
func Simulate(N int64, filePath string, steps int) {

	// Read File
	simpleWorldMap, err := helper.ReadWorldMapFile(filePath)
	check(err)

	// Initate world
	var world internal.World = internal.NewWorld()

	err = world.ValidateMap(simpleWorldMap)
	check(err)

	err = world.Construct(simpleWorldMap)
	check(err)

	err = world.InhabitAlien(int(N))
	check(err)

	for i := 0; i < steps; i++ {

		//fmt.Printf("Iteration %d... \n", i)

		world.Evaluate()

		zeroMpves, err := world.MoveAliens()
		check(err)

		if zeroMpves {
			fmt.Println("All aliens are trapped/dead, simulation is done...")
			break
		}
	}

	fmt.Println(fmt.Sprintf("%d Steps has been taken!", steps))

	world.PrintWorld()
}
