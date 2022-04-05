# Mad Alien Invasion

This is a simulation of aliens​ who ​are​ ​about​ ​to​ ​invade​ ​the​ ​earth, the world map consists of cities each city is connected to other cities following a specific direction (north, south, east, west).
The aliens are placed randomly on the map, and they move randomly following valid links between cities; if two or more aliens are found in one city, their connection and aliens are destroyed.

  - [Assumptions](#assumptions)
  - [Getting started](#getting-started)
  - [Setting up Dev](#setting-up-dev)
  - [Examples](#examples)
  - [Run Tests](#run-tests)

## Assumptions
* The **city** names can only contain alphabets
* The **connection/link** between two cities is not directed
* If two or more than two **aliens** are in one city the city is destroyed
* For every iteration, all aliens move unless the alien is trapped
* The simulation ends if 
  * the number of iterations are completed 
  * if all aliens are trapped


## Getting started
* First install [GoLang](https://golang.org/doc/install)

## Setting up Dev
* First parameter is NumberOfAliens, **default** is 2.
* Second parameter is FilePath, **default** is the below map.

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
Bar has been destroyed by 1 2! 
All aliens are trapped/dead, simulation is done...
Baz east=Foo
Foo west=Baz south=Qu-ux
Qu-ux north=Foo
Bee
```

## Run Tests
* Run test cases
```bash
cd internal
go test
```
* Generate coverage report
```bash
go test -coverprofile=coverage.out # Generate coverage file
go tool cover -html=coverage.out # Open coverage report
go tool cover -html=coverage.out -o coverage.html # Save the coverage report file
```