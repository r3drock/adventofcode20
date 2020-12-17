package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput() string{
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return string(filestring)
}

func genMasks(maskstring string) (uint64, uint64) {
	var ormask uint64
	var andmask uint64
	for _, c := range maskstring {
		switch c {
		case 'X':
			ormask |= 0
			andmask |= 1
		case '0':
			ormask |= 0
			andmask |= 0
		case '1':
			ormask |= 1
			andmask |= 1
		default:
			_, _ = fmt.Fprintf(os.Stderr, "wrong mask format!\n")
			os.Exit(1)
		}
		ormask <<= 1
		andmask <<= 1
	}
	ormask >>= 1
	andmask >>= 1
	return ormask, andmask
}

func part1(filestring string) {
	memory := make(map[int]uint64)
	var ormask, andmask uint64
	for _, line := range strings.Split(filestring,"\n") {
		if len(line) < 2 {
			continue
		}
		if line[1] == 'a' {
			ormask, andmask = genMasks(line[7:])
		} else if line[1] == 'e' {
			memidx, err := strconv.Atoi(strings.Split(line[4:], "]")[0])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			signedval, err := strconv.ParseInt(strings.Split(line, "= ")[1],10, 64)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			val := uint64(signedval)
			val |= ormask
			val &= andmask
			memory[memidx] = val
		} else {
			continue
		}
	}
	var sum uint64
	for _, val := range memory {
		sum += val
	}
	fmt.Println(sum)
}

func genMultipleVals(maskstring string, val uint64) []uint64 {
	var vals = make([]uint64,1)
	vals[0] = val
	var currentbit uint64 = 1 << 35
	for j := 0; currentbit > 0; j++ {
		switch maskstring[j] {
		case 'X':
			previouslen := len(vals)
			for i := 0; i < previouslen; i++ {
				val1 := vals[i]
				vals[i] &= math.MaxUint64 ^ currentbit
				val1 |= currentbit
				vals = append(vals, val1)
			}
		case '0':
		case '1':
			for i := 0; i < len(vals); i++ {
				vals[i] |= currentbit
			}
		default:
			_, _ = fmt.Fprintf(os.Stderr, "wrong mask format!\n")
			os.Exit(1)
		}
		currentbit >>= 1
	}
	return vals
}

func part2(filestring string) {
	memory := make(map[uint64]uint64)
	var lastMask string
	for _, line := range strings.Split(filestring,"\n") {
		if len(line) < 2 {
			continue
		}
		if line[1] == 'a' {
			lastMask = line[7:]
		} else if line[1] == 'e' {
			tempmemidx, err := strconv.Atoi(strings.Split(line[4:], "]")[0])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			var memidx = uint64(tempmemidx)
			signedval, err := strconv.ParseInt(strings.Split(line, "= ")[1],10, 64)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			val := uint64(signedval)
			vals := genMultipleVals(lastMask, memidx)
			for i := 0; i < len(vals); i++ {
				memory[vals[i]] = val
			}

		} else {
			continue
		}
	}
	var sum uint64
	for _, val := range memory {
		sum += val
	}
	fmt.Println(sum)
}

func main() {
	input := parseInput()
	part1(input)
	part2(input)
}
