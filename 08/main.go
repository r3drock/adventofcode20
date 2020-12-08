package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type GameConsole struct {
	acc int64
	pc  int64
}

type OPCODE int

const (
	NOP OPCODE = iota
	ACC
	JMP
)

type Instruction struct {
	opcode  OPCODE
	operand int64
}

func nop(console *GameConsole, operand int64) {
	console.pc++
}

func acc(console *GameConsole, operand int64) {
	console.acc += operand
	console.pc++
}

func jmp(console *GameConsole, operand int64) {
	console.pc += operand
}

func ParseProgram() []Instruction {
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	var lines = strings.Split(string(filestring), "\n")

	var program = make([]Instruction, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		var linesplit = strings.Split(string(line), " ")
		var opcode OPCODE
		switch linesplit[0] {
		case "nop":
			opcode = NOP
		case "acc":
			opcode = ACC
		case "jmp":
			opcode = JMP
		default:
			fmt.Errorf("illegal instruction decoded %s", linesplit[0])
			os.Exit(1)
		}
		op, err := strconv.ParseInt(linesplit[1], 10, 64)
		if err != nil {
			fmt.Println("Operand parse err: %v", err)
		}

		program = append(program, Instruction{opcode: opcode, operand: op})
	}
	return program
}

func executeOneInst(console *GameConsole, opcode OPCODE, operand int64) {
	switch opcode {
	case NOP:
		nop(console, operand)
	case ACC:
		acc(console, operand)
	case JMP:
		jmp(console, operand)
	default:
		fmt.Errorf("tried executing illegal instruction %d", opcode)
		os.Exit(1)
	}
}

func Part1(program []Instruction) GameConsole {
	var console GameConsole
	var programcopy []Instruction = make([]Instruction, len(program))
	copy(programcopy, program)
	var instructionsexecuted = make(map[int64]bool)
	for !instructionsexecuted[console.pc] {
		instructionsexecuted[console.pc] = true
		executeOneInst(&console, programcopy[console.pc].opcode, programcopy[console.pc].operand)
	}
	return console
}

func TryExecuting(program []Instruction, i int, instructionlimit int32) int64 {
	if program[i].operand == 0 {
		return -1
	}
	var console GameConsole
	var programcopy = make([]Instruction, len(program))
	copy(programcopy, program)
	if programcopy[i].opcode == JMP {
		programcopy[i].opcode = NOP
	} else if programcopy[i].opcode == NOP {
		programcopy[i].opcode = JMP
	}
	var endpc = int64(len(programcopy))
	var instructions int32
	for endpc != console.pc {
		executeOneInst(&console, programcopy[console.pc].opcode, programcopy[console.pc].operand)
		if instructions == instructionlimit {
			break
		}
		instructions++
	}
	if endpc == console.pc {
		return console.acc
	}
	return -1
}

func Part2Sequential(program []Instruction) int64 {
	var result int64 = -1
	var instructionlimit int64 = int64(len(program))
	for result == -1 && instructionlimit < math.MaxInt32 {
		for i := 0; i < len(program); i++ {
			result = TryExecuting(program, i, int32(instructionlimit))
			if result != -1 {
				return result
			}
		}
		fmt.Println("Increase instructionlimit from: ", instructionlimit, " to ", instructionlimit*2)
		instructionlimit *= 2
	}
	if result != -1 {
		fmt.Print("impossible")
		os.Exit(1)
	}
	return result
}

func Part2(program []Instruction) int64 {
	var wg sync.WaitGroup
	var result int64 = -1
	var instructionlimit int64 = int64(len(program))
	for result == -1 && instructionlimit < math.MaxInt32 {
		for i := 0; i < len(program); i++ {
			wg.Add(1)
			go func(i int) {
				var acc = TryExecuting(program, i, int32(instructionlimit))
				if acc != -1 {
					result = acc
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		if result == -1 {
			fmt.Println("Increase instructionlimit from: ", instructionlimit, " to ", instructionlimit*2)
			instructionlimit *= 2
		}
	}

	return result
}

func main() {
	start := time.Now()
	var program = ParseProgram()
	parse := time.Now()
	var result1 = Part1(program).acc
	part1 := time.Now()
	var result2 = Part2Sequential(program)
	part2seq := time.Now()
	var result3 = Part2(program)
	part2 := time.Now()

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	fmt.Println("Time needed for parsing: ", parse.Sub(start))
	fmt.Println("Time needed for part1: ", part1.Sub(parse))
	fmt.Println("Time needed for part2 sequential: ", part2seq.Sub(part1))
	fmt.Println("Time needed for part2 concurrently: ", part2.Sub(part2seq))
	fmt.Fprintf(os.Stdout, "parallel speedup: %.2f", float64(part2seq.Sub(part1))/float64(part2.Sub(part2seq)))
}
