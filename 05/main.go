package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile() []string {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")

}
func pow2(exponent int) int {
	var retval int = 1;
	for exponent > 0 {
		exponent--
		retval *= 2;
	}
	return retval
}
func findRow(s string) int {
	var min, max int = 0, 127
	for i := 6; i >= 0; i-- {
		if s[6-i] == 'B' {
			min += pow2(i)
		} else if s[6-i] == 'F' {
			max -= pow2(i)
		} else {
			fmt.Fprintf(os.Stderr,"ERROR: 6-%d: %d s[6-%d]: %c \n", i, 6-i, i, s[6-i])
			os.Exit(1)
		}
	}
	if min != max {
		fmt.Errorf("ERROR min: %d max: %d\n", min, max)
		os.Exit(1)
	}
	return min
}
func findColumn(s string) int {
	var min, max int = 0, 7
	for i := 2; i >= 0; i-- {
		if s[9-i] == 'R' {
			min += pow2(i)
		} else if s[9-i] == 'L' {
			max -= pow2(i)
		} else {
			fmt.Fprintf(os.Stderr,"ERROR: 6-%d: %d s[6-%d]: %c \n", i, 9-i, i, s[9-i])
			os.Exit(1)
		}
	}
	if min != max {
		fmt.Fprintf(os.Stderr,"ERROR min: %d max: %d\n", min, max)
		os.Exit(1)
	}
	return min
}

func calcSeatId(row int, column int) int {
	return row * 8 + column
}

func FindSeat(ids []bool) int {
	for i := 0; i < (len(ids)-2); i++ {
		if ids[i] && (!ids[i+1]) && ids[i+2] {
			return i + 1
		}
	}
	return -1
}

func main() {
	var lines = ReadFile()
	var highestId int
	var ids = make([]bool,128*8)
	for _, line := range lines {
		if len(line) <= 1 {
			continue
		}
		var row = findRow(line)
		var column = findColumn(line)
		var id = calcSeatId(row, column)
		ids[id] = true
		if id > highestId {
			highestId = id
		}
		//fmt.Println(id, row, column)
	}

	fmt.Println(highestId)
	fmt.Println(FindSeat(ids))
}
