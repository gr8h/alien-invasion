package helper

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadWorldMapFile(filePath string) (map[string]map[string]string, error) {

	f, err := os.Open(filePath)

	check(err)

	defer f.Close()

	var simpleWorldMap = make(map[string]map[string]string)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var line string = scanner.Text()
		lineSplit := strings.Split(line, " ")

		var cityName string = strings.TrimSpace(lineSplit[0])

		var simpleConnectionMap = make(map[string]string)

		for i := 1; i < len(lineSplit); i++ {
			cityConnection := strings.Split(lineSplit[i], "=")

			var toCity string = strings.TrimSpace(cityConnection[1])
			var dir string = strings.TrimSpace(cityConnection[0])

			simpleConnectionMap[dir] = toCity
		}

		simpleWorldMap[cityName] = simpleConnectionMap
	}

	return simpleWorldMap, nil
}
