package main

import (
	"alien-invasion/pkg/simulator"
	"fmt"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 2 {

		N, err := strconv.ParseInt(argsWithoutProg[0], 6, 12)
		Check(err)

		filePath := argsWithoutProg[1]

		simulator.Simulate(N, filePath)

	} else {
		fmt.Println("Please enter valid arguments. \n Number of alliens follwed by the file name. \n ex: 50 ../../map1.txt")
	}
}
