package internal

import (
	"fmt"
)

// Const
const North string = "north"
const East string = "east"
const South string = "south"
const West string = "west"

var oppositeDirection = map[string]string{
	North: South,
	East:  West,
	South: North,
	West:  East,
}

// Struct
type Connection struct {
	From      *City
	To        *City
	Direction string

	//Private
	alive bool
}

// Operations

func NewConnection(from *City, to *City, direction string) Connection {
	var e Connection = Connection{from, to, direction, true}
	return e
}

func (c *Connection) Destroy() error {

	if !c.IsAlive() {
		return fmt.Errorf("Connection:Destroy - Already destroyed.")
	}

	c.alive = false

	fmt.Printf("Connection between %s & %s is destroyed. \n", c.From.Name, c.To.Name)

	return nil
}

func (c *Connection) IsAlive() bool {
	return c.alive
}
