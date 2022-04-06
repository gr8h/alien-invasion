package internal

import (
	"fmt"
	"math/rand"
	"time"
)

type Alien struct {
	Id int

	alive bool
	city  *City // By Referance
}

// Operations

/*
- Generate new object with default values
Parameters: Alien Id
Returns: New object
*/
func NewAlien(Id int) Alien {
	var e Alien = Alien{Id, false, nil}
	return e
}

// Check if alien is alive
func (a *Alien) IsAlive() bool {
	return a.alive
}

// Get the current alien city
func (a *Alien) GetCity() *City {
	return a.city
}

// Change the alien city
func (a *Alien) SetCity(city *City) {
	a.city = city
}

// Destroy the alien
func (a *Alien) Destroy() error {

	if !a.IsAlive() {
		return fmt.Errorf("Aline:Destroy - Already dead %d", a.Id)
	}

	a.alive = false

	if showExtraMessages {
		fmt.Printf("Alien %d is dead. \n", a.Id)
	}
	return nil
}

// Check if the alien can move, or the alien is trapped or dead
func (a *Alien) CanMove() (bool, error) {

	if a.city == nil {
		return false, fmt.Errorf("Aline:CanMove - No city found %d", a.Id)
	}

	if !a.IsAlive() {
		return false, nil
	}

	// Out Connection
	var hasOutConnection bool = false
	for _, v := range a.city.Connections {
		if v.IsAlive() && v.From.Name == a.GetCity().Name {
			hasOutConnection = true
			break
		}
	}

	return hasOutConnection, nil
}

// Move the alien to a random city where connection exists and is valid
func (a *Alien) Move() error {

	//fmt.Println(a.GetCity().Name)
	if a.GetCity() == nil {
		return fmt.Errorf("Aline:Move - No city found %d", a.Id)
	}

	if !a.IsAlive() {
		return nil
	}

	// Curret City Connections
	var candidates []*Connection
	for _, v := range a.city.Connections {
		if v.IsAlive() && v.From.Name == a.GetCity().Name {
			candidates = append(candidates, v)
		}
	}

	// Get Rand City
	var count int = len(candidates)

	if count == 0 {
		return nil
	}

	rand.Seed(time.Now().UTC().UnixNano())
	var i = rand.Intn(count)

	// Remove alien
	a.city.RemoveAlien(a)

	// Set new city
	a.city = candidates[i].To
	candidates[i].To.AddAlien(a)

	if showExtraMessages {
		fmt.Printf("Alien %d moved to city %s. \n", a.Id, a.city.Name)
	}

	return nil
}
