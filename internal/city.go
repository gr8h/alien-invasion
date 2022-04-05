package internal

import "fmt"

type City struct {
	Name        string
	Connections []*Connection
	Aliens      map[int]*Alien

	//Private
	alive bool
}

// Operations

/*
- Generate new object with default values
Takes: City Name
Returns: New object
*/
func NewCity(name string) City {
	var e City = City{name, nil, make(map[int]*Alien), true}
	return e
}

// Check if city is alive
func (a *City) IsAlive() bool {
	return a.alive
}

// Get the city status
func (c *City) Evaluate() (int, bool) {
	return len(c.Aliens), c.IsAlive()
}

// Add an alien to the city
func (c *City) AddAlien(a *Alien) {
	c.Aliens[a.Id] = a
}

// Remove an alien from the city
func (c *City) RemoveAlien(a *Alien) {
	delete(c.Aliens, a.Id)
}

// Add a connection between cities
func (c *City) AddConnection(conn *Connection) {
	c.Connections = append(c.Connections, conn)
}

// Destro the city, connections and aliens
func (c *City) Destroy() error {

	if !c.IsAlive() {
		return fmt.Errorf("City:Destroy - Already destroyed %s", c.Name)
	}

	c.alive = false

	for _, v := range c.Connections {
		v.Destroy()
	}

	fmt.Printf("%s has been destroyed by aliens ", c.Name)
	for _, v := range c.Aliens {
		fmt.Printf("%d ", v.Id)
		v.Destroy()
	}
	fmt.Printf("!\n")

	return nil
}

// Check if the city is alive and has alien whom are trapped
func (c *City) HasTrappedAliens() (bool, error) {

	if !c.IsAlive() {
		return false, fmt.Errorf("City:Destroy - Already destroyed %s", c.Name)
	}

	// Out Connection
	var hasOutConnection bool = false
	for _, v := range c.Connections {
		if v.From.Name == c.Name && v.IsAlive() {
			hasOutConnection = true
			break
		}
	}

	// Alive Aliens
	var hasAliveAliens bool = false
	for _, v := range c.Aliens {
		if v.IsAlive() {
			hasAliveAliens = true
			break
		}
	}

	return hasAliveAliens && !hasOutConnection, nil
}
