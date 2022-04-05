# Mad Alien Invasion

This is a simulation of aliens​ ​are​ ​about​ ​to​ ​invade​ ​the​ ​earth, the world map consists of cities each city is connected to other citities following a specific direction (north, south, east, west).
The aliens are placed randomy in the map, and they move randomly following valid links between cities; if two or more aliens are found in one city the city, its connection and aliens are destroyed.

- [Mad Alien Invasion](#mad-alien-invasion)
  - [Assumptions](#assumptions)
  - [Getting started](#getting-started)
  - [Setting up Dev](#setting-up-dev)
  - [Examples](#examples)
  - [Run Tests](#run-tests)

## Assumptions
* The **city** names can only contain alphabits
* The **connection/link** between two cities are not directed
* If two or more than two **aliens** are in one city the city is destroyed
* For every iteration all aliens moves unless the alien is trapped
* The simulation ends if 
  * the number of iterations are completed 
  * or if all aliens are trapped
  * or all cities are destroyed


## Getting started
* First install [GoLang](https://golang.org/doc/install)

## Setting up Dev
* First parameter is NumberOfAliens, the **default** is 5.
* Second parameter is FilePath, the **default** is the below map:

```bash
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
Bee east=Bar
Baz east=Foo
Qu-ux north=Foo
```
* Execute the simulation
```bash
go run cmd/main.go [NumberOfAliens] [FilePath]
```

## Examples
* Run the below command
```bash
go run cmd/main.go 10 
```
* Output example
```bash
Bar has been destroyed! 
All aliens are trapped/dead, simulation is done...
Baz east=Foo
Foo west=Baz south=Qu-ux
Qu-ux north=Foo
Bee
```

## Run Tests
```bash
cd internal
go test
go test -coverprofile=coverage.out 
```