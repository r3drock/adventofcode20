package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"runtime/pprof"
	"strings"
	"sync"
	"time"
)

func Max(num int, nums ...int) int {
	for _, a := range nums {
		if num < a {
			num = a
		}
	}
	return num
}

func getXD(vector []int, cubelen int, coords ...int) int {
	bias := (len(vector) + 1) / 2
	idx := bias
	for i, coord := range coords {
		idx += coord * int(math.Pow(float64(cubelen), float64(i)))
	}
	return vector[idx]
}

func get4D(vector []int, cubelen int, x int , y int, z int , w int) int {
	cubelenstart := cubelen
	y*=cubelen
	cubelen *= cubelenstart
	z*=cubelen
	cubelen *= cubelenstart
	w*=cubelen
	bias := (len(vector) + 1) / 2
	return vector[x+y+z+w+bias]
}

func setXD(vector []int, cubelen int, val int, coords ...int) {
	bias := (len(vector) + 1) / 2
	idx := bias
	for i, coord := range coords {
		idx += coord * int(math.Pow(float64(cubelen), float64(i)))
	}
	vector[idx] = val
}

func getActiveNeighbourCount4D(vector []int, X int, Y int, Z int, W int, cubelen int) int {
	sum := 0
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						//||x < -cubelen/2 || y < -cubelen/2 || z < -cubelen/2 ||
						//	x > cubelen/2 || y > cubelen/2 || z > cubelen/2
						continue
					}
					sum += get4D(vector, cubelen, X+x, Y+y, Z+z, W+w)
				}
			}
		}
	}
	return sum
}

func get3D(vector []int, x int, y int, z int, cubelen int) int {
	bias := (len(vector) + 1) / 2
	return vector[x+(y*cubelen)+(z*cubelen*cubelen)+bias]
}
func set3D(vector []int, x int, y int, z int, cubelen int, val int) {
	bias := (len(vector) + 1) / 2
	vector[x+(y*cubelen)+(z*cubelen*cubelen)+bias] = val
}

func parseStart(cycles int) ([]int, []int, int) {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}
	input := string(bytes)
	lines := strings.Split(input, "\n")

	zsize, ysize, xsize := 1, len(lines), len(lines[0])
	if xsize != ysize {
		fmt.Fprintf(os.Stderr, "no cube as input %d %d %s", ysize, xsize, "\n")
		os.Exit(1)
	}
	cubelen := Max(xsize, ysize, zsize)
	if cubelen%2 == 0 {
		cubelen++
	}
	cubelen += 2*cycles + 2
	upperboundsize := cubelen * cubelen * cubelen
	var pocketdimension = make([]int, upperboundsize+1)
	var pocketdimension4D = make([]int, (upperboundsize*cubelen)+1)
	y := -len(lines) / 2
	for _, line := range lines {
		x := -len(line) / 2
		for _, c := range line {
			if c == '#' {
				set3D(pocketdimension, x, y, 0, cubelen, 1)
				setXD(pocketdimension4D, cubelen, 1, x, y, 0, 0)
			}
			x++
		}
		y++
	}
	return pocketdimension, pocketdimension4D, cubelen
}

func getActiveNeighbourCount(pocketdimension []int, X int, Y int, Z int, cubelen int) int {
	sum := 0
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x == 0 && y == 0 && z == 0 {
					//||x < -cubelen/2 || y < -cubelen/2 || z < -cubelen/2 ||
					//	x > cubelen/2 || y > cubelen/2 || z > cubelen/2
					continue
				}
				sum += get3D(pocketdimension, X+x, Y+y, Z+z, cubelen)
			}
		}
	}
	return sum
}

func cycle(pocketdimension []int, cubelen int) {
	var oldpocketdimension = make([]int, len(pocketdimension))
	copy(oldpocketdimension, pocketdimension)
	for z := 1 - (cubelen / 2); z <= (cubelen/2)-1; z++ {
		for y := 1 - (cubelen / 2); y <= (cubelen/2)-1; y++ {
			for x := 1 - (cubelen / 2); x <= (cubelen/2)-1; x++ {
				count := getActiveNeighbourCount(oldpocketdimension, x, y, z, cubelen)
				if get3D(oldpocketdimension, x, y, z, cubelen) == 1 {
					if !(count == 2 || count == 3) {
						set3D(pocketdimension, x, y, z, cubelen, 0)
					}
				} else {
					if count == 3 {
						set3D(pocketdimension, x, y, z, cubelen, 1)
					}
				}
			}
		}
	}
}

func cycle4D(pocketdimension []int, oldpocketdimension []int, w int, cubelen int, wg *sync.WaitGroup) {
	for z := 1 - (cubelen / 2); z <= (cubelen/2)-1; z++ {
		for y := 1 - (cubelen / 2); y <= (cubelen/2)-1; y++ {
			for x := 1 - (cubelen / 2); x <= (cubelen/2)-1; x++ {
				count := getActiveNeighbourCount4D(oldpocketdimension, x, y, z, w, cubelen)
				if get4D(oldpocketdimension, cubelen, x, y, z, w) == 1 {
					if !(count == 2 || count == 3) {
						setXD(pocketdimension, cubelen, 0, x, y, z, w)
					}
				} else {
					if count == 3 {
						setXD(pocketdimension, cubelen, 1, x, y, z, w)
					}
				}
			}
		}
	}
	wg.Done()
}

const cycles = 6

func part1(pocketdimension []int, cubelen int) int {
	for i := 0; i < cycles; i++ {
		cycle(pocketdimension, cubelen)

	}
	sum := 0
	for i := 0; i < len(pocketdimension); i++ {
		sum += pocketdimension[i]
	}
	return sum
}
func part2(pocketdimension4D []int, cubelen int) int {
	for i := 0; i < cycles; i++ {
		var oldpocketdimension4D = make([]int, len(pocketdimension4D))
		copy(oldpocketdimension4D, pocketdimension4D)
		wg := new(sync.WaitGroup)
		for w := 1 - (cubelen / 2); w <= (cubelen/2)-1; w++ {
			wg.Add(1)
			go cycle4D(pocketdimension4D, oldpocketdimension4D, w, cubelen, wg)
		}
		wg.Wait()
	}
	sum := 0
	for i := 0; i < len(pocketdimension4D); i++ {
		sum += pocketdimension4D[i]
	}
	return sum
}

func main() {
	start := time.Now()
	pocketdimension, pocketdimension4D, cubelen := parseStart(cycles)
	parse := time.Now()
	fmt.Println( part1(pocketdimension, cubelen))
	part1 := time.Now()
	proffile,err := os.Create("harper.prof")
	defer proffile.Close()
	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}
	pprof.StartCPUProfile(proffile)
	fmt.Println( part2(pocketdimension4D, cubelen))
	part2 := time.Now()
	pprof.StopCPUProfile()
	fmt.Println("Time needed for parsing: ", parse.Sub(start))
	fmt.Println("Time needed for part1: ", part1.Sub(parse))
	fmt.Println("Time needed for part2: ", part2.Sub(part1))
}
