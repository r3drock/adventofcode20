package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

const (
	North int = 0
	West = 270
	South = 180
	East = 90
	Left = iota + 267
	Right
	Forward
)

type instruction struct {
	action int
	value  int
}

type shipState struct {
	x      int
	y      int
	xway int
	yway int
	facing int
}

func parseInstructions() []instruction {
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	var instructions = make([]instruction, 0, len(filestring))

	for _, line := range bytes.Split(filestring, []byte("\n")) {
		if string(line) == "" {
			continue
		}
		var action int
		switch line[0] {
		case 'N':
			action = North
		case 'S':
			action = South
		case 'W':
			action = West
		case 'E':
			action = East
		case 'L':
			action = Left
		case 'R':
			action = Right
		case 'F':
			action = Forward
		}
		value, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			fmt.Errorf("Wrong format in input!")
			os.Exit(1)
		}

		instructions = append(instructions, instruction{action: action, value: value})
	}

	return instructions
}

func rotate(x float64, y float64, degree float64) (int,int) {
	rad := degree * (math.Pi/ 180.0)
	return int(math.Round(x* math.Cos(rad) -y* math.Sin(rad))), int(math.Round(x* math.Sin(rad) + y* math.Cos(rad)))
}


func moveWaypoint(ship *shipState, inst instruction) {
	switch inst.action {
	case North:
		ship.yway += inst.value
	case South:
		ship.yway -= inst.value
	case West:
		ship.xway -= inst.value
	case East:
		ship.xway += inst.value
	case Left:
		ship.xway, ship.yway = rotate(float64(ship.xway), float64(ship.yway), float64(inst.value))
	case Right:
		ship.xway, ship.yway = rotate(float64(ship.xway), float64(ship.yway), float64(-inst.value))

	case Forward:
		ship.x += inst.value * ship.xway
		ship.y += inst.value * ship.yway
	}
}

func move(ship *shipState, inst instruction) {
	switch inst.action {
	case North:
		ship.y += inst.value
	case South:
		ship.y -= inst.value
	case West:
		ship.x -= inst.value
	case East:
		ship.x += inst.value
	case Left:
		ship.facing += 360
		ship.facing -= inst.value
		ship.facing %= 360

	case Right:
		ship.facing += inst.value
		ship.facing %= 360

	case Forward:
		switch ship.facing {
		case North:
			ship.y += inst.value
		case South:
			ship.y -= inst.value
		case West:
			ship.x -= inst.value
		case East:
			ship.x += inst.value
		}
	}
}

func moveInstructions(instructions []instruction, fn func(*shipState, instruction)){
	ship := shipState{0, 0, 10, 1, East}
	for _, inst := range instructions {
		fn(&ship, inst)
	}
	fmt.Println(Manhattan(ship.x, ship.y))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Manhattan(x int, y int) int {
	return Abs(x) + Abs(y)
}

func main() {
	instructions := parseInstructions()
	moveInstructions(instructions, move)
	moveInstructions(instructions, moveWaypoint)
}
