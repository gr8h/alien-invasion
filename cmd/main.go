package main

import (
	"alien-invasion/pkg/simulator"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const defaultN = 2
const defaultSteps = 10000

func main() {

	argsWithoutProg := os.Args[1:]

	var N int64
	filePath, err := filepath.Abs("../test/valid_map.txt")
	check(err)

	if len(argsWithoutProg) == 0 {
		N = defaultN
	} else if len(argsWithoutProg) == 1 {
		N, err = strconv.ParseInt(argsWithoutProg[0], 6, 12)
		check(err)
	} else if len(argsWithoutProg) == 2 {
		N, err = strconv.ParseInt(argsWithoutProg[0], 6, 12)
		check(err)

		if N <= 0 || N > 10000 {
			panic("Please N should be between 1 and 10000!")
		}

		filePath = argsWithoutProg[1]
	} else {
		err := fmt.Errorf("Please enter valid arguments! ```go run cmd/main.go [NumberOfAliens] [FilePath]```")
		check(err)
	}

	simulator.Simulate(N, filePath, defaultSteps)
}
