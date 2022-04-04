package simulator

import (
	"alien-invasion/internal"
	"alien-invasion/pkg/helper"
	"fmt"
)

const steps = 10

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

	err = world.ValidateMap(simpleWorldMap)
	Check(err)

	err = world.Construct(simpleWorldMap)
	Check(err)

	err = world.InhabitAlien(int(N))
	Check(err)

	for i := 0; i < steps; i++ {

		fmt.Printf("Iteration %d... \n", i)

		err = world.Evaluate()
		Check(err)

		zeroMpves, err := world.MoveAlien()
		Check(err)

		if zeroMpves {
			fmt.Println("All aliens are trapped/dead, simulation is done...")
			break
		}
	}

	world.PrintWorld()
}
