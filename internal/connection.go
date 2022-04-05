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

func NewConnection(from *City, to *City, direction string) Connection {

	_, _, err := validateConnection(from, to, direction)
	check(err)

	var e Connection = Connection{from, to, direction, true}
	return e
}

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

func (c *Connection) IsAlive() bool {
	return c.alive
}

func validateConnection(from *City, to *City, direction string) (bool, *Connection, error) {

	notvalid, conn := from.contains(to)
	if notvalid {
		return false, conn, fmt.Errorf(fmt.Sprintf("validateConnection f - Connection already exist %s %s %s.", conn.From.Name, conn.Direction, conn.To.Name))
	}

	return true, nil, nil
}

func (from *City) contains(city *City) (bool, *Connection) {
	for _, a := range from.Connections {
		if a.To == city {
			return true, a
		}
	}
	return false, nil
}
