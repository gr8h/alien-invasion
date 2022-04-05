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

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func main() {

	argsWithoutProg := os.Args[1:]

	var N int64 = 5
	filePath, err := filepath.Abs("../test/valid_map.txt")
	check(err)

	if len(argsWithoutProg) == 1 {
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
		fmt.Println("Please enter valid arguments. \n Number of alliens follwed by the file name. \n ex: 50 ../../map1.txt")
	}

	simulator.Simulate(N, filePath)
}
