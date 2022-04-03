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

func (c *Connection) Destroy() {
	c.alive = false

	fmt.Printf("Connection between %s & %s is destroyed", c.From.Name, c.To.Name)
}

func (c *Connection) IsAlive() bool {
	return c.alive
}
