package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Edge struct {
	node  *Node
	count uint
}

type Node struct {
	name      string
	visited bool
	neighbors []Edge
	backneighbors []Edge
}

func DFSCOUNTUNIQUE(startnodepointer *Node) uint {
	if startnodepointer.visited == true {
		return 0
	} else {
		startnodepointer.visited = true
	}
	var retval uint
	for _, edge := range startnodepointer.backneighbors {
		if edge.node.visited == true {
			continue
		} else {
			retval++
			retval += DFSCOUNTUNIQUE(edge.node)
		}
	}
	return retval
}


func DFSCOUNT(startnodepointer *Node, multiplier uint) uint {
	var retval uint
	for _, edge := range startnodepointer.neighbors {
		{
			retval+= edge.count * multiplier
			retval += DFSCOUNT(edge.node, edge.count * multiplier)
		}
	}
	return retval
}

func ParseFile() {
	start := time.Now()
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	var lines = strings.Split(string(filestring), "\n")
	var nodes = make([]Node, len(lines))
	var namehashmap = make(map[string]*Node)
	for i, line := range lines {
		var bagname = strings.Split(string(line), " bags")[0]
		node := Node{bagname, false,make([]Edge, 0, 1), make([]Edge, 0, 1)}
		nodes[i] = node
		namehashmap[bagname] = &nodes[i]
	}
	for _, line := range lines {
		if line == "" {
			continue
		}
		var bagnames = strings.Split(line, " bags contain ")
		var words = strings.Split(bagnames[1], " ")
		if words[0] == "no" {
			continue
		}
		for _, bagname := range strings.Split(bagnames[1], ", ") {
			var bagnamewords = strings.Split(bagname, " ")
			var count, _ = strconv.Atoi(bagnamewords[0])
			var name = bagnamewords[1] + " " + bagnamewords[2]
			var childNodePointer = namehashmap[name]

			var parentNodePointer = namehashmap[bagnames[0]]
			parentNodePointer.neighbors = append(parentNodePointer.neighbors, Edge{childNodePointer, uint(count)})
			childNodePointer.backneighbors = append(childNodePointer.backneighbors, Edge{parentNodePointer, 0})
		}
	}
	parse := time.Now()
	fmt.Println(DFSCOUNTUNIQUE(namehashmap["shiny gold"]))
	part1 := time.Now()
	fmt.Println(DFSCOUNT(namehashmap["shiny gold"], 1))
	part2 := time.Now()
	fmt.Println("Time needed for parsing: ", parse.Sub(start))
	fmt.Println("Time needed for part1: ", part1.Sub(parse))
	fmt.Println("Time needed for part2: ", part2.Sub(part1))
}

func main() {

	ParseFile()
}