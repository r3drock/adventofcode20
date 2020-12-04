package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile() []string {
	data, err := ioutil.ReadFile("input2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")

}

func stepthroughforest(lines []string, xstep int, ystep int) int {
	var width, height int = len(lines[0]), len(lines)

	var x, y int
	var treecount int
	for y < height {
		if lines[y][x] == '#' {
			treecount++
		}
		x += xstep
		x %= width
		y += ystep
	}
	return treecount
}

func main() {
	var lines = ReadFile()
	fmt.Println(stepthroughforest(lines, 3, 1))

	fmt.Println(stepthroughforest(lines, 1, 1) *
		stepthroughforest(lines, 3, 1) *
		stepthroughforest(lines, 5, 1) *
		stepthroughforest(lines, 7, 1) *
		stepthroughforest(lines, 1, 2))

}
