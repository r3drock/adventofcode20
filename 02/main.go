package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func part2() int64 {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		os.Exit(1)
	}
	var validpasswordcount int64

	for _, line := range strings.Split(string(data), "\n") {
		var splitted = strings.Split(line, ":")
		if len(splitted) != 2 {
			break
		}
		var left, password = splitted[0], splitted[1]
		splitted = strings.Split(left, " ")
		var passwordrange = splitted[0]
		var passwordchar uint8 = []uint8(splitted[1])[0]
		splitted = strings.Split(passwordrange, "-")
		var minindex, _ = strconv.Atoi(splitted[0])
		var maxindex, _ = strconv.Atoi(splitted[1])

		if (password[minindex] == passwordchar) != (password[maxindex] == passwordchar) {
			validpasswordcount++
		}
	}
	return validpasswordcount
}
func part1() int64 {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		os.Exit(1)
	}
	var validpasswordcount int64

	for _, line := range strings.Split(string(data), "\n") {
		var splitted = strings.Split(line, ":")
		if len(splitted) != 2 {
			break
		}
		var left, password = splitted[0], splitted[1]
		splitted = strings.Split(left, " ")
		if len(splitted) != 2 {
			fmt.Fprintf(os.Stderr, "Wrong format\n", err)
			os.Exit(1)
		}
		var passwordrange = splitted[0]
		var passwordchar rune = []rune(splitted[1])[0]
		splitted = strings.Split(passwordrange, "-")
		if len(splitted) != 2 {
			fmt.Fprintf(os.Stderr, "Wrong format\n", err)
			os.Exit(1)
		}
		var passwordmin, err = strconv.Atoi(splitted[0])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(1)
		}
		var passwordmax, err2 = strconv.Atoi(splitted[1])
		if err2 != nil {
			// handle error
			fmt.Println(err2)
			os.Exit(1)
		}
		var counts = make(map[rune]int)
		for _, c := range password[1:] {
			counts[c]++
		}
		if counts[passwordchar] >= passwordmin && counts[passwordchar] <= passwordmax {
			validpasswordcount++
		}
	}
	return validpasswordcount
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
