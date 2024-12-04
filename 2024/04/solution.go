package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type WordSearch struct {
	x, m, a, s [][]bool
	xmas       [][]rune
}

func (w *WordSearch) AddLine(l string) {
	n := len(l)
	x := make([]bool, n)
	m := make([]bool, n)
	a := make([]bool, n)
	s := make([]bool, n)
	xmas := make([]rune, n)
	for i, c := range l {
		switch c {
		case 'X':
			x[i] = true
		case 'M':
			m[i] = true
		case 'A':
			a[i] = true
		case 'S':
			s[i] = true
		}
		xmas[i] = c
	}
	w.x = append(w.x, x)
	w.m = append(w.m, m)
	w.a = append(w.a, a)
	w.s = append(w.s, s)
	w.xmas = append(w.xmas, xmas)
}

func (w *WordSearch) isLetter(c rune, i, j int) bool {
	if i < 0 || j < 0 || i >= len(w.xmas) || j >= len(w.xmas[0]) {
		return false
	}
	if w.xmas[i][j] == c {
		return true
	}
	return false
}

func translateDirToDelta(dir string) (int, int) {
	var di, dj int
	switch dir {
	case "NW":
		di = -1
		dj = -1
	case "N":
		di = 0
		dj = -1
	case "NE":
		di = 1
		dj = -1
	case "W":
		di = -1
		dj = 0
	case "E":
		di = 1
		dj = 0
	case "SW":
		di = -1
		dj = 1
	case "S":
		di = 0
		dj = 1
	case "SE":
		di = 1
		dj = 1
	}
	return di, dj
}

func (w *WordSearch) isMAS(i, j int, dir string) bool {
	di, dj := translateDirToDelta(dir)

	i += di
	j += dj
	if w.isLetter('M', i, j) {
		i += di
		j += dj
		if w.isLetter('A', i, j) {
			i += di
			j += dj
			if w.isLetter('S', i, j) {
				return true
			}
		}
	}

	return false
}

func (w *WordSearch) countXmasAt(i, j int) int {
	num := 0
	if w.isMAS(i, j, "NW") {
		num++
	}
	if w.isMAS(i, j, "N") {
		num++
	}
	if w.isMAS(i, j, "NE") {
		num++
	}
	if w.isMAS(i, j, "W") {
		num++
	}
	if w.isMAS(i, j, "E") {
		num++
	}
	if w.isMAS(i, j, "SW") {
		num++
	}
	if w.isMAS(i, j, "S") {
		num++
	}
	if w.isMAS(i, j, "SE") {
		num++
	}
	return num
}

func (w *WordSearch) countXmas() int {
	num := 0
	for i, row := range w.xmas {
		for j, x := range row {
			if x == 'X' {
				num += w.countXmasAt(i, j)
			}
		}
	}
	return num
}

func applyDirToXY(dir string, x, y int) (int, int) {
	dx, dy := translateDirToDelta(dir)
	return x + dx, y + dy
}

func (w *WordSearch) cornersAreMS(i1, j1, i2, j2 int) bool {
	if w.isLetter('M', i1, j1) && w.isLetter('S', i2, j2) {
		return true
	}
	if w.isLetter('S', i1, j1) && w.isLetter('M', i2, j2) {
		return true
	}
	return false
}

func (w *WordSearch) isCrossMASAt(i, j int) bool {
	i1a, j1a := applyDirToXY("NW", i, j)
	i1b, j1b := applyDirToXY("SE", i, j)
	i2a, j2a := applyDirToXY("NE", i, j)
	i2b, j2b := applyDirToXY("SW", i, j)

	if w.cornersAreMS(i1a, j1a, i1b, j1b) {
		if w.cornersAreMS(i2a, j2a, i2b, j2b) {
			return true
		}
	}

	return false
}

func (w *WordSearch) countCrossMAS() int {
	num := 0
	for i, row := range w.xmas {
		for j, x := range row {
			if x == 'A' {
				if w.isCrossMASAt(i, j) {
					num++
				}
			}
		}
	}
	return num
}

func a(input *bufio.Scanner) {
	var n int
	w := &WordSearch{}

	for n = 0; input.Scan(); n++ {
		l := input.Text()
		w.AddLine(l)
	}

	fmt.Println(w.countXmas())
}

func b(input *bufio.Scanner) {
	var n int
	w := &WordSearch{}

	for n = 0; input.Scan(); n++ {
		l := input.Text()
		w.AddLine(l)
	}

	fmt.Println(w.countCrossMAS())
}

func main() {
	real := flag.Bool("real", false, "Whether to use the real input")
	runA := flag.Bool("a", false, "Run program a")
	runB := flag.Bool("b", false, "Run program a")

	flag.Parse()

	fileName := "sample.txt"
	if *real {
		fileName = "input.txt"
	}

	input := ReadInput(fileName)

	if !*runA && !*runB {
		panic("Did not specify a or b")
	}

	if *runA && *runB {
		panic("Specified both a and b")
	}

	if *runA {
		a(input)
	}
	if *runB {
		b(input)
	}
}

func ReadInput(fileName string) *bufio.Scanner {
	input, err := os.Open(fileName)
	if err != nil {
		panic("could not open")
	}

	return bufio.NewScanner(input)
}
