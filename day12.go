package main

import (
	"fmt"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Ship struct {
	position  Point
	direction Point
	waypoint  Point
}

var translations = map[string]Point{
	"N": {0, 1},
	"S": {0, -1},
	"E": {1, 0},
	"W": {-1, 0},
}

var rotations = map[string][]int{
	"L90":  []int{0, -1, 1, 0},
	"L180": []int{-1, 0, 0, -1},
	"L270": []int{0, 1, -1, 0},
	"R90":  []int{0, 1, -1, 0},
	"R180": []int{-1, 0, 0, -1},
	"R270": []int{0, -1, 1, 0},
}

func solveDay12Example() {
	instructions := []string{"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	moveShipWithWaypoint(instructions)
}

func solveDay12Part1() {
	instructions := getDataFromFile("day12")
	moveShip(instructions)
}

func moveShip(instructions []string) {
	ship := &Ship{Point{0, 0}, Point{1, 0}, Point{0, 0}}
	for _, xmd := range instructions {
		ship.move(xmd)
	}
	fmt.Printf("Ship is at position (%d, %d) Distance from start: %d\n", ship.position.x, ship.position.y, abs(ship.position.x)+abs(ship.position.y))
}

func rotate(d Point, rotation []int) Point {
	return Point{d.x*rotation[0] + d.y*rotation[1], d.x*rotation[2] + d.y*rotation[3]}
}

func (ship *Ship) move(cmd string) {
	key := cmd[0:1]
	value, _ := strconv.Atoi(cmd[1:])
	if key == "F" {
		ship.position = Point{ship.position.x + value*ship.direction.x, ship.position.y + value*ship.direction.y}
	} else if key == "R" || key == "L" {
		ship.direction = rotate(ship.direction, rotations[cmd])
	} else {
		ship.position = Point{ship.position.x + value*translations[key].x, ship.position.y + value*translations[key].y}
	}
}

func solveDay12Part2() {
	instructions := getDataFromFile("day12")
	moveShipWithWaypoint(instructions)
}

func moveShipWithWaypoint(instructions []string) {
	ship := &Ship{Point{0, 0}, Point{1, 0}, Point{10, 1}}
	for _, xmd := range instructions {
		ship.moveWithWaypoint(xmd)
	}
	fmt.Printf("Ship is at position (%d, %d) Distance from start: %d\n", ship.position.x, ship.position.y, abs(ship.position.x)+abs(ship.position.y))
}

func (ship *Ship) moveWithWaypoint(cmd string) {
	key := cmd[0:1]
	value, _ := strconv.Atoi(cmd[1:])
	if key == "F" {
		ship.position = Point{ship.position.x + value*ship.waypoint.x, ship.position.y + value*ship.waypoint.y}
	} else if key == "R" || key == "L" {
		ship.waypoint = rotate(ship.waypoint, rotations[cmd])
	} else {
		ship.waypoint = Point{ship.waypoint.x + value*translations[key].x, ship.waypoint.y + value*translations[key].y}
	}
}
