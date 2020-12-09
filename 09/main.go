package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
    "container/list"
)

func ParseProgram() []int {
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	var lines = strings.Split(string(filestring), "\n")

	var numbers = make([]int, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func FindFirstInvalid(numbers []int, preamblelength int) int {
	var lastnumbersFIFO = list.New()
	var lastnumbershashmap = make(map[int]bool)
	for i := 0; i < preamblelength; i++ {
		lastnumbersFIFO.PushBack(numbers[i])
		lastnumbershashmap[numbers[i]] = true
	}

	for i := preamblelength; i < len(numbers); i++ {
		var valid bool
		{ //is valid?
			for e := lastnumbersFIFO.Front(); e != nil; e = e.Next() {
				secondnumber := numbers[i] - e.Value.(int)
				if lastnumbershashmap[secondnumber] {
					valid = true
					break
				}
			}
		}
		if !valid {
			return numbers[i]
		}
		var frontelement = lastnumbersFIFO.Front()
		var frontelementvalue = frontelement.Value.(int)
		lastnumbersFIFO.Remove(frontelement)
		lastnumbersFIFO.PushBack(numbers[i])
		delete(lastnumbershashmap, frontelementvalue)
		lastnumbershashmap[numbers[i]] = true
	}
	 return -1
}

func FindSmallestandLargest(numbers []int, left int, right int) (int, int){
	var smallest = math.MaxInt32
	var largest = math.MinInt32
	for i := left; i <= right; i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		} else if numbers[i] > largest {
			largest = numbers[i]
		}
	}
	return smallest, largest
}

func FindContigouus(numbers []int, sum int) int {
	var numsums = make([][]int,0, len(numbers))
	numsums = append(numsums, make([]int, len(numbers)))
	for i := 0; i < len(numbers); i++ {
		numsums[0][i] = numbers[i]
	}
	for j := 0; j < len(numbers); j++ {
		numsums = append(numsums, make([]int, len(numsums[j])-1))
		for i := 0; i < len(numsums[j])-1; i++ {
				numsums[j+1][i] = numsums[j][i] + numsums[0][i+j+1]
				if numsums[j+1][i] == sum{
					smallest, largest := FindSmallestandLargest(numbers, i,i+j+1)
					return smallest + largest
				}
		}
	}
	for _, numsum := range numsums {
		fmt.Println(numsum)
	}
	return -1
}

func main() {
	var numbers = ParseProgram()
	var invalid = FindFirstInvalid(numbers, 25)
	fmt.Println(invalid)
	fmt.Println(FindContigouus(numbers, invalid))
}
