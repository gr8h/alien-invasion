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

func NewCity(name string) City {
	var e City = City{name, nil, make(map[int]*Alien), true}
	return e
}

func (a *City) IsAlive() bool {
	return a.alive
}

func (c *City) Evaluate() (int, bool) {
	return len(c.Aliens), c.IsAlive()
}

func (c *City) AddAlien(a *Alien) {
	c.Aliens[a.Id] = a
}

func (c *City) RemoveAlien(a *Alien) {
	delete(c.Aliens, a.Id)
}

func (c *City) AddConnection(conn *Connection) {
	c.Connections = append(c.Connections, conn)
}

func (c *City) Destroy() error {

	if !c.IsAlive() {
		return fmt.Errorf("City:Destroy - Already destroyed %s", c.Name)
	}

	c.alive = false

	for _, v := range c.Connections {
		v.Destroy()
	}

	for _, v := range c.Aliens {
		v.Destroy()
	}

	fmt.Printf("%s has been destroyed! \n", c.Name)

	return nil
}

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
