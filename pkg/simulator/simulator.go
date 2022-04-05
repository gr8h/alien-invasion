package simulator

import (
	"alien-invasion/internal"
	"alien-invasion/pkg/helper"
	"fmt"
)

const steps = 10

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Simulate(N int64, filePath string) {

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

		fmt.Printf("Iteration %d... \n", i)

		err = world.Evaluate()
		check(err)

		zeroMpves, err := world.MoveAliens()
		check(err)

		if zeroMpves {
			fmt.Println("All aliens are trapped/dead, simulation is done...")
			break
		}
	}

	world.PrintWorld()
}
