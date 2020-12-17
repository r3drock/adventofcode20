package main

import (
	"fmt"
	"time"
)

type number struct {
	turnsspoken [2]int64
	amountspoken int64
}
func calc(startnumbers []int64, limit int64) int64 {

	var lastnumspoken int64
	protocol := make(map[int64]number)
	for i, startnumber := range startnumbers {
		if val, ok := protocol[startnumber]; ok {
			val.turnsspoken[0] = val.turnsspoken[1]
			val.turnsspoken[1] = int64(i+1)
			val.amountspoken++
			protocol[startnumber] = val
		} else {
			protocol[startnumber] = number{
				turnsspoken: [2]int64{0,int64(i+1)},
				amountspoken: 1,
			}
		}
		lastnumspoken = startnumber
	}

	for i := int64(len(startnumbers) + 1); i <= limit; i++ {
		if protocol[lastnumspoken].amountspoken == 1 {
			if val, ok := protocol[0]; ok {
				val.turnsspoken[0] = val.turnsspoken[1]
				val.turnsspoken[1] = int64(i)
				val.amountspoken++
				protocol[0] = val
				lastnumspoken = 0
			}
		} else {
			if val, ok := protocol[lastnumspoken]; ok {
				difference := val.turnsspoken[1] - val.turnsspoken[0]
				if val, ok := protocol[difference]; ok {
					val.turnsspoken[0] = val.turnsspoken[1]
					val.turnsspoken[1] = i
					val.amountspoken++
					protocol[difference] = val
				} else {
					protocol[difference] = number{
						turnsspoken:  [2]int64{0,i},
						amountspoken: 1,
					}
				}
				lastnumspoken = difference
			}
		}
	}
	return lastnumspoken
}
func main() {
	start := time.Now()
	startnumbers := []int64{19, 20, 14, 0, 9, 1}
	println(calc(startnumbers, 2020))
	part1 := time.Now()
	println(calc(startnumbers, 30000000))
	part2 := time.Now()
	fmt.Println("Time needed for part1: ", part1.Sub(start))
	fmt.Println("Time needed for part2: ", part2.Sub(part1))
}
