package internal

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type World struct {
	Cities      map[string]*City
	Aliens      []Alien
	Connections []Connection

	CityNames []string
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	fmt.Println("Initialize World...")
}

func NewWorld() World {
	var e World = World{make(map[string]*City), nil, nil, nil}
	return e
}

func (w *World) Construct(simpleWorld map[string]map[string]string) {

	// Create City
	for name, _ := range simpleWorld {
		var newCity City = NewCity(name)
		w.Cities[name] = &newCity
		w.CityNames = append(w.CityNames, name)
	}

	// Create Connections
	for from, element := range simpleWorld {

		for dir, to := range element {

			var newConn = NewConnection(w.Cities[from], w.Cities[to], dir)
			w.Connections = append(w.Connections, newConn)

			w.Cities[from].AddConnection(&newConn)
			w.Cities[to].AddConnection(&newConn)
		}
	}
}

func (w *World) InhabitAlien(n int) {

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
}

func (w *World) MoveAlien() bool {

	var moveCount int = 0

	for i := range w.Aliens {
		var alien Alien = w.Aliens[i]

		canMove, err := alien.CanMove()
		Check(err)

		if canMove {
			moveCount += 1

			err := alien.Move()
			Check(err)
		}

	}

	return moveCount > 0
}

func (w *World) Evaluate() {
	// Evaluate City

	for _, city := range w.Cities {

		alienCount, isAlive := city.Evaluate()

		if isAlive == true && alienCount >= 2 {
			err := city.Destroy()
			Check(err)
		}
	}
}

func (w *World) PrintWorld() {

	for _, city := range w.Cities {

		if !city.IsAlive() {
			continue
		}

		var cityName string = city.Name

		var sb strings.Builder
		sb.WriteString(cityName)

		for _, conn := range city.Connections {

			if conn.From.Name == cityName && conn.IsAlive() {
				sb.WriteString(fmt.Sprintf(" %s=%s", conn.Direction, conn.To.Name))
			}
		}

		fmt.Println(sb.String())
	}
}
