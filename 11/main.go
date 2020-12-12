package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func parseLayout() [][]byte {
	filestring, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	var layout = make([][]byte, 0, len(filestring))

	for i, line := range bytes.Split(filestring, []byte("\n")) {
		if string(line) == "" {
			continue
		}
		layout = append(layout, make([]byte, 0, len(line)))
		for _, c := range line {
			if c == byte('L') || c == byte('.') {
				layout[i] = append(layout[i], c)
			}
		}
	}

	return layout
}

func print2D(layout [][]byte) {
	for _, row := range layout {
		for j := 0; j < len(row); j++ {
			fmt.Printf("%c", row[j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkSurroundingSeats(y int, x int, layout [][]byte) int {
	var numberseats int
	for ydelta := -1; ydelta <= 1; ydelta++ {
		for xdelta := -1; xdelta <= 1; xdelta++ {
			if xdelta == 0 && ydelta == 0 ||
				x+xdelta < 0 || y+ydelta < 0 ||
				x+xdelta >= len(layout[0]) || y+ydelta >= len(layout) {
				continue
			}
			if layout[y+ydelta][x+xdelta] == '#' {
				numberseats++
			}
		}
	}
	return numberseats
}

func checkline(y int, x int, layout [][]byte, ydir int, xdir int) int {
	for i := 1; (i*ydir)+y < len(layout) && (i*ydir)+y >= 0 && (i*xdir)+x < len(layout[0]) && (i*xdir)+x >= 0; i++ {
		if layout[(i*ydir)+y][(i*xdir)+x] == '#' {
			return 1
		} else if layout[(i*ydir)+y][(i*xdir)+x] == 'L' {
			return 0
		}
	}
	return 0
}

func checkSurroundingSeatsFar(y int, x int, layout [][]byte) int {
	var numberseats int
	for ydir := -1; ydir <= 1; ydir++ {
		for xdir := -1; xdir <= 1; xdir++ {
			if xdir == 0 && ydir == 0 {
				continue
			}
			numberseats += checkline(y, x, layout, ydir, xdir)
		}
	}
	return numberseats
}

func countSeats(a [][]byte, b [][]byte) int {
	seatcount := 0
outer:
	for y := 0; y < len(a); y++ {
		for x := 0; x < len(a[0]); x++ {
			if a[y][x] != b[y][x] {
				seatcount = -1
				break outer
			}
			if a[y][x] == '#' {
				seatcount++
			}
		}
	}
	return seatcount
}

type surround func(int, int, [][]byte) int

func applyrules(layout [][]byte, fn surround, seatThreshold int) int {
	var oldlayout = make([][]byte, len(layout))
	for i, row := range layout {
		oldlayout[i] = make([]byte, len(row))
		copy(oldlayout[i], row)
	}
	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[0]); x++ {
			switch oldlayout[y][x] {
			case '.':
				continue
			case 'L':
				if fn(y, x, oldlayout) == 0 {
					layout[y][x] = '#'
				}
			case '#':
				if fn(y, x, oldlayout) >= seatThreshold {
					layout[y][x] = 'L'
				}
			}
		}
	}
	return countSeats(oldlayout, layout)
}
func part1(layout [][]byte) {
	for true {
		seatcount := applyrules(layout, checkSurroundingSeats, 4)
		if seatcount != -1 {
			println(seatcount)
			break
		}
	}
}
func part2(layout [][]byte) {
	for true {
		seatcount := applyrules(layout, checkSurroundingSeatsFar, 5)
		if seatcount != -1 {
			println(seatcount)
			break
		}
	}
}
func main() {
	start := time.Now()
	layout := parseLayout()
	parse := time.Now()
	var oldlayout = make([][]byte, len(layout))
	for i, row := range layout {
		oldlayout[i] = make([]byte, len(row))
		copy(oldlayout[i], row)
	}
	part1(layout)
	part1 := time.Now()
	//part2(oldlayout)
	f, err := os.OpenFile("a.gif", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f, oldlayout)
	part2 := time.Now()
	fmt.Println("Time needed for parsing: ", parse.Sub(start))
	fmt.Println("Time needed for part1: ", part1.Sub(parse))
	fmt.Println("Time needed for part2: ", part2.Sub(part1))
}

var palette = []color.Color{color.White, color.Black, color.RGBA{R: 222, G: 122, B: 0, A: 1}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	redIndex   = 2
)

func lissajous(out io.Writer, layout [][]byte) {
	const (
		delay = 12 // delay between frames in 10ms units
	)
	anim := gif.GIF{LoopCount: -1}
	for true {
		rect := image.Rect(0, 0, 2*len(layout[0]), 2*len(layout))
		img := image.NewPaletted(rect, palette)
		fillFrame(img, layout)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		seatcount := applyrules(layout, checkSurroundingSeats, 4)
		if seatcount != -1 {
			println(seatcount)
			rect := image.Rect(0, 0, 2*len(layout[0]), 2*len(layout))
			img := image.NewPaletted(rect, palette)
			fillFrame(img, layout)
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
			break
		}
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
func fillFrame(img *image.Paletted, layout [][]byte) {
	for y, _ := range layout {
		for x := 0; x < len(layout[y]); x++ {
			var colorIndex uint8
			switch layout[y][x] {
			case '.':
				colorIndex = whiteIndex
			case 'L':
				colorIndex = blackIndex
			case '#':
				colorIndex = redIndex
			}
			img.SetColorIndex(x*2, y*2, colorIndex)
			img.SetColorIndex(x*2+1, y*2, colorIndex)
			img.SetColorIndex(x*2, y*2+1, colorIndex)
			img.SetColorIndex(x*2+1, y*2+1, colorIndex)
		}
	}
}