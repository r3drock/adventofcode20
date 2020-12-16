package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type schedule struct {
	timestamp int
	pos int
}

func parseSchedules() (int, []schedule) {
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	s := bytes.Split(filestring, []byte("\n"))
	timestamp, err := strconv.Atoi(string(s[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	var schedules = make([]schedule, 0, len(s[1]))

	for i, val := range bytes.Split(s[1], []byte(",")) {
		if val[0] == 'x' {
			continue
		}
		ts, err := strconv.Atoi(string(val))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		schedules = append(schedules, schedule{ts, i})
	}
	return timestamp, schedules
}

func findEarliestBus(timestamp int, schedules []schedule) {
	shortestWaittime := math.MaxInt32
	var id int
	for _, schedule := range schedules {
		waittime := schedule.timestamp - (timestamp % schedule.timestamp)
		if waittime < shortestWaittime {
			shortestWaittime = waittime
			id = schedule.timestamp
		}
	}
	fmt.Print(shortestWaittime, id, shortestWaittime*id)
}

func findEarliestTimestamp(timestamp int, schedules []schedule) {
	builder := strings.Builder{}

println()
	for i, schedule := range schedules {
		builder.WriteString(fmt.Sprintf("t + %d mod %d = 0", schedule.pos, schedule.timestamp))
		if i < len(schedules) - 1 {
			builder.WriteString(",")
		}
	}
	wolframalphaquery := builder.String()
	fmt.Println("Query for Wolframalpha: ", wolframalphaquery)
}

func main() {
	timestamp, schedules := parseSchedules()
	findEarliestBus(timestamp, schedules)
	findEarliestTimestamp(timestamp, schedules)
}
