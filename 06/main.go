package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func answeredallyes(answercounts map[rune]int, size int) int {
	var retval int
	for _, counts := range answercounts {
		if counts == size {
			retval++
		}
	}
	return retval
}

func ParseFile() (int,int) {
	lines, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	var retval,retval2 int
	for _, grouplines := range strings.Split(string(lines), "\n\n") {
		var size = 1
		for _, c := range grouplines {
			if c == '\n' {
				size++
			}
		}
		var groupanswercounts = make(map[rune]int)
		for _, personanswers := range strings.Split(grouplines, "\n") {
			for _, answer := range personanswers {
				groupanswercounts[answer]++
			}
		}
		retval += len(groupanswercounts)
		retval2 += answeredallyes(groupanswercounts, size)
	}
	return retval,retval2
}

func main() {
	var part1sol, part2sol = ParseFile()
	fmt.Printf("%d\n%d\n", part1sol, part2sol)
}
