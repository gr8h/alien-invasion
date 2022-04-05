package internal

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type World struct {
	Cities map[string]*City
	Aliens []Alien

	CityNames []string
}

func init() {
	if showExtraMessages {
		fmt.Println("Initialize World...")
	}
}

func NewWorld() World {
	var e World = World{make(map[string]*City), nil, nil}
	return e
}

func (w *World) ValidateMap(simpleWorld map[string]map[string]string) error {

	var tempMap map[string][]string
	tempMap = make(map[string][]string)

	for from, elements := range simpleWorld {
		for direction, to := range elements {
			var key string
			if from > to {
				key = fmt.Sprintf("%s#%s", from, to)
			} else {
				key = fmt.Sprintf("%s#%s", to, from)
			}
			tempMap[key] = append(tempMap[key], direction)
		}
	}

	for _, directions := range tempMap {

		if len(directions) < 2 {
			//fmt.Errorf(fmt.Sprintf("ValidateMap: Missing Connection between %s", cityPair))
			return fmt.Errorf("ValidateMap: Missing Connection")
		}

		if len(directions) > 2 {
			//fmt.Errorf(fmt.Sprintf("ValidateMap: Extra Connection between %s", cityPair))
			return fmt.Errorf("ValidateMap: Extra Connection")
		}

		if strings.Compare(directions[0], oppositeDirection[directions[1]]) != 0 {
			//fmt.Errorf(fmt.Sprintf("ValidateMap: Wrong Direction between %s", cityPair))
			return fmt.Errorf("ValidateMap: Wrong Direction")
		}
	}

	return nil
}

func (w *World) Construct(simpleWorld map[string]map[string]string) error {

	// Create Connections
	for from, elements := range simpleWorld {

		_, found := w.Cities[from]
		if !found {
			temp := NewCity(from)
			w.Cities[from] = &temp
			w.CityNames = append(w.CityNames, temp.Name)
		}

		for dir, to := range elements {

			_, found := w.Cities[to]
			if !found {
				temp := NewCity(to)
				w.Cities[to] = &temp
				w.CityNames = append(w.CityNames, temp.Name)
			}

			var newConn = NewConnection(w.Cities[from], w.Cities[to], dir)

			w.Cities[from].AddConnection(&newConn)
			w.Cities[to].AddConnection(&newConn)
		}
	}

	return nil
}

func (w *World) InhabitAlien(n int) error {

	// Create Aline
	for i := 0; i < n; i++ {

		var newAline = NewAlien(i)

		// Get Rand City
		rand.Seed(time.Now().UnixNano())
		var r = rand.Intn(len(w.CityNames))

		// Inhabit
		currentCity := w.Cities[w.CityNames[r]]

		currentCity.AddAlien(&newAline)
		newAline.SetCity(currentCity)

		w.Aliens = append(w.Aliens, newAline)
	}

	return nil
}

func (w *World) MoveAliens() (bool, error) {

	var moveCount int = 0

	for i := range w.Aliens {
		var alien *Alien = &w.Aliens[i]

		if !alien.IsAlive() {
			continue
		}

		canMove, err := alien.CanMove()
		check(err)

		// If no moves are taken, this means that all aliens are trapped
		if canMove {
			err := alien.Move()
			check(err)

			moveCount += 1
		}

	}

	return moveCount == 0, nil
}

func (w *World) Evaluate() error {

	// Evaluate City

	for _, city := range w.Cities {

		// If city is alive and alens in the city is greater than two, then destry the city
		alienCount, isAlive := city.Evaluate()

		if isAlive && alienCount >= 2 {
			err := city.Destroy()
			check(err)
		}
	}

	return nil
}

func (w *World) PrintWorld() {

	if showExtraMessages {
		fmt.Println("Printing World...")
	}

	for _, city := range w.Cities {

		var sb strings.Builder

		if !city.IsAlive() {
			continue
		}

		sb.WriteString(city.Name)

		for _, conn := range city.Connections {
			if conn.IsAlive() && conn.From.Name == city.Name {
				sb.WriteString(fmt.Sprintf(" %s=%s", conn.Direction, conn.To.Name))
			}
		}

		fmt.Println(sb.String())
	}
}
