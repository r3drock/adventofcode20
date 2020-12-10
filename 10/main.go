package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ParseRatings() []int {
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	var lines = strings.Split(string(filestring), "\n")

	var ratings = make([]int, 1, len(lines))
	ratings[0] = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		rating, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		ratings = append(ratings, rating)
	}
	sort.Ints(ratings)
	ratings = append(ratings, ratings[len(ratings)-1] + 3)
	return ratings
}

func FindDifferences(ratings []int) (int,int) {
	var onedifferences int
	var threedifferences int
	for i := 0; i < len(ratings)-1; i++ {
		var rating = ratings[i]
		if ratings[i+1] == rating + 1 {
			onedifferences++
		} else if ratings[i+1] == rating + 3 {
			threedifferences++
		}
	}
	return onedifferences, threedifferences
}

func FindArangements(ratings []int, i int, arrangementcounts []int) int {
	if i == len(ratings) || i == len(ratings) -1 {
		return 1
	} else if i > len(ratings) {
		return 0
	}
	var retval = 0
	var rating = ratings[i]
	for j := 1; j <= 3; j++ {
		if i+j < len(ratings) {
			if ratings[i+j] <= rating+3 {
				if arrangementcounts[i+j] != -1 {
					retval += arrangementcounts[i+j]
				} else {
					retval += FindArangements(ratings, i+j, arrangementcounts)
				}
			}
		}
	}
	arrangementcounts[i] = retval
	return retval
}

func main() {
	start := time.Now()
	var ratings = ParseRatings()
	parse := time.Now()
	fmt.Println(ratings)
	onedifferences, threedifferences := FindDifferences(ratings)
	fmt.Println(onedifferences, threedifferences, onedifferences * threedifferences)
	part1 := time.Now()
	var arrangementcounts = make([]int, len(ratings))
	for i := 0; i < len(arrangementcounts); i++ {
		arrangementcounts[i] = -1
	}
	fmt.Println(FindArangements(ratings, 0, arrangementcounts))
	part2 := time.Now()
	fmt.Println("Time needed for parsing: ", parse.Sub(start))
	fmt.Println("Time needed for part1: ", part1.Sub(parse))
	fmt.Println("Time needed for part2: ", part2.Sub(part1))
}
