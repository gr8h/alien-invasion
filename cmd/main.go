package main

import (
	"alien-invasion/internal/alien"
	"fmt"
)

func main() {
	var e = alien.New(5)
	fmt.Println(e.Id)

	e.Move()
	fmt.Println(e.Id)
}
