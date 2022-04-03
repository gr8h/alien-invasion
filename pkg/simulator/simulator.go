package simulator

import (
	"alien-invasion/internal"
	"alien-invasion/pkg/helper"
	"fmt"
)

const steps = 10000

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Simulate(N int64, filePath string) {

	// Read File
	simpleWorldMap, err := helper.ReadWorldMapFile(filePath)
	Check(err)

	// Initate world
	world := internal.NewWorld()

	world.Construct(simpleWorldMap)

	world.InhabitAlien(int(N))

	for i := 0; i < steps; i++ {

		world.Evaluate()

		var allTrapped = world.MoveAlien()

		if allTrapped {
			fmt.Println("All aliens are trapped, simulation is done...")
			break
		}
	}

	world.PrintWorld()
}
