package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanFile() map[int]int {
	counts := make(map[int]int)
	f, err := os.Open("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		num, err := strconv.Atoi(input.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(1)
		}
		counts[num]++
	}
	f.Close()
	return counts
}

func part1(counts map[int]int, target int) (int,int){
	for val, _ := range counts {
		val2 := target - val
		if counts[val2] == 1 {
			return val,val2
		}
	}
	return 0,0
}

func part2(counts map[int]int, target int) (int,int,int) {
	for val, _ := range counts {
		valsum := target - val
		var val2,val3 = part1(counts, valsum)
		if val2 != 0 && val3 != 0 {
			return val, val2, val3
		}
	}
	return 0,0,0
}

func main() {
	counts := ScanFile()
	{
		var val,val2 = part1(counts, 2020)
		fmt.Println("part 1: ", val * val2)
	}
	var val,val2,val3 = part2(counts, 2020)
	fmt.Println("part 2: ", val, val2, val3, val * val2 * val3)
}
