package internal

import (
	"fmt"
)

// Struct
type Connection struct {
	From      *City
	To        *City
	Direction string

	//Private
	alive bool
}

// Operations

/*
- Generate new object with default values
Parameters: City fron, City To and Direction (north, south, west, east)
Returns: New object
*/
func NewConnection(from *City, to *City, direction string) Connection {
	var e Connection = Connection{from, to, direction, true}
	return e
}

// Destory the connection
func (c *Connection) Destroy() error {

	if !c.IsAlive() {
		return fmt.Errorf("Connection:Destroy - Already destroyed.")
	}

	c.alive = false

	if showExtraMessages {
		fmt.Printf("Connection between %s & %s is destroyed. \n", c.From.Name, c.To.Name)
	}

	return nil
}

// Check if the connection is alive
func (c *Connection) IsAlive() bool {
	return c.alive
}
