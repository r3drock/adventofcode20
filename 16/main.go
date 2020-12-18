package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const ArraySize int = 1000

func parse() [ArraySize]bool {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	input := string(bytes)
	paragraphs := strings.Split(input, "\n\n")
	//print("\n--------------\n")
	//for _, paragraph := range paragraphs {
	//	print(paragraph, "\n--------------\n")
	//}
	validvalues := [ArraySize]bool{false}
	for _, line := range strings.Split(paragraphs[0], "\n") {
		parts := strings.Split(line, ":")
		//fieldname := parts[0]
		var a, b, c, d int
		fmt.Sscanf(parts[1], "%d-%d or %d-%d", &a, &b, &c, &d)
		for i := a; i <= b; i++ {
			validvalues[i] = true
		}
		for i := c; i <= d; i++ {
			validvalues[i] = true
		}
	}
	scanningErrorRate := 0
	for i, line := range strings.Split(paragraphs[2], "\n") {
		if i == 0 {
			continue
		}
		values := strings.Split(line, ",")
		for _, val := range values {
			value, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			if value >= ArraySize {
				scanningErrorRate += value
			} else if !validvalues[value] {
				scanningErrorRate += value
			}
		}
	}
	fmt.Println(scanningErrorRate)
	return validvalues
}

func max(num int, nums ...int) int {
	for _, a := range nums {
		if num < a {
			num = a
		}
	}
	return num
}

func parse2() (int,int) {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	input := string(bytes)
	paragraphs := strings.Split(input, "\n\n")
	validvalues := [ArraySize]bool{false}
	lines := strings.Split(paragraphs[0], "\n")
	fields := make(map[string][]bool)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		var a, b, c, d int
		fmt.Sscanf(parts[1], "%d-%d or %d-%d", &a, &b, &c, &d)

		fieldname := parts[0]
		fieldvalues := make([]bool, max(a, b, c, d)+1)

		for i := a; i <= b; i++ {
			validvalues[i] = true
			fieldvalues[i] = true
		}
		for i := c; i <= d; i++ {
			validvalues[i] = true
			fieldvalues[i] = true
		}
		fields[fieldname] = fieldvalues
		//fmt.Println(line, len(fieldvalues))
		//fmt.Print(fieldvalues)
		//fmt.Scanln()
	}
	tickets := make([][]int, 0)
	scanningErrorRate := 0
outer:
	for _, line := range strings.Split(paragraphs[2], "\n") {
		if strings.HasPrefix(line, "nearby tickets:") {
			continue
		}
		values := strings.Split(line, ",")
		valid := true
		for _, val := range values {
			value, err := strconv.Atoi(val)
			if err != nil {
				continue outer
			}
			if value >= ArraySize || !validvalues[value] {
				scanningErrorRate += value
				valid = false
			}
		}
		if valid {
			ticket := make([]int, len(values))
			for i, val := range values {
				value, err := strconv.Atoi(val)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cant happen %v\n", err)
					os.Exit(1)
				}
				ticket[i] = value
			}
			tickets = append(tickets, ticket)
		}
	}

	columns := make([][]int, len(tickets[0]))
	for _, ticket := range tickets {
		for i, _ := range ticket {
			column := make([]int, len(tickets))
			columns[i] = column
		}
		break
	}
	for i, ticket := range tickets {
		for j := 0; j < len(ticket); j++ {
			columns[j][i] = ticket[j]
		}
	}
	fmt.Println(columns[0])

	columnnames := make([]string, len(columns))
	for j, column := range columns {
		candidates := make(map[string]int, len(fields))
		for i := 0; i < len(column); i++ {
			for name, field := range fields {
				fmt.Println(len(field))
				if field[column[i]] {
					candidates[name] += 1
				}
			}
		}
		for name, candidate := range candidates {
			if candidate == len(column) {
				columnnames[j] = name
			}
		}
	}
	fmt.Println("\n---------------------")
	for i, _ := range columns {
		fmt.Println(columnnames[i])
	}
	fmt.Println("\n---------------------")
	var result = 1
	for _, line := range strings.Split(paragraphs[1], "\n") {
		if len(line) <= 1 || strings.HasPrefix(line, "your ticket:") {
			continue
		}
		for i, number := range strings.Split(line, ",") {
			if strings.HasPrefix(columnnames[i], "departure") {
				val, err := strconv.Atoi(number)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
					os.Exit(1)
				}
				result *= val
				println(val)
			}
		}
	}
	return scanningErrorRate,result
}

func columnnames() {

}

func main() {
	//validvalues :=
	//parse()
	scanningErrorRate,result := parse2()
	fmt.Println(scanningErrorRate,result)
	//for i, bools := range fields {
	//	fmt.Println(i, ": ", len(bools))
	//}
}
