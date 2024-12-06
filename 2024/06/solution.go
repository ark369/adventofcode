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

type Solver struct {
	m    [][]rune
	ms   [][][]rune
	x, y int
	dir  rune
}

func MakeSolver(input *bufio.Scanner) *Solver {
	s := &Solver{}
	s.m = [][]rune{}
	s.ms = [][][]rune{}
	x := 0
	for input.Scan() {
		l := input.Text()
		row := []rune{}
		rows := [][]rune{}
		for y, c := range l {
			rowss := []rune{}
			if c != '.' {
				rowss = append(rowss, c)
			}
			row = append(row, c)
			switch c {
			case '^':
				fallthrough
			case 'v':
				fallthrough
			case '<':
				fallthrough
			case '>':
				s.x = x
				s.y = y
				s.dir = c
			}
			rows = append(rows, rowss)
		}
		s.m = append(s.m, row)
		s.ms = append(s.ms, rows)
		x++
	}
	return s
}

func (s *Solver) OutOfBounds(x, y int) bool {
	if x < 0 || y < 0 || x >= len(s.ms) || y >= len(s.ms[0]) {
		return true
	}
	return false
}

func (s *Solver) BoundaryAt(x, y int) bool {
	if s.OutOfBounds(x, y) {
		return false
	}
	if s.m[x][y] == '#' {
		return true
	}
	return false
}

func (s *Solver) Move() (bool, bool) {
	newX := s.x
	newY := s.y
	switch s.dir {
	case '^':
		newX -= 1
		if s.BoundaryAt(newX, newY) {
			s.dir = '>'
			newX += 1
			newY += 1
		}
	case 'v':
		newX += 1
		if s.BoundaryAt(newX, newY) {
			s.dir = '<'
			newX -= 1
			newY -= 1
		}
	case '<':
		newY -= 1
		if s.BoundaryAt(newX, newY) {
			s.dir = '^'
			newX -= 1
			newY += 1
		}
	case '>':
		newY += 1
		if s.BoundaryAt(newX, newY) {
			s.dir = 'v'
			newX += 1
			newY -= 1
		}
	}

	s.m[s.x][s.y] = 'X'

	if s.OutOfBounds(newX, newY) {
		return true, true
	}

	visited := false
	if s.m[newX][newY] == 'X' {
		visited = true
	}
	s.m[newX][newY] = s.dir
	s.x = newX
	s.y = newY

	//fmt.Println(s)
	return visited, false
}

func (s *Solver) CountVisited() int {
	v := 0
	var visited, done bool
	for ; !done; visited, done = s.Move() {
		if !visited {
			v++
		}
	}
	return v
}

func (s *Solver) Contains(x, y int, dir rune) bool {
	if s.OutOfBounds(x, y) {
		return false
	}
	for _, d := range s.ms[x][y] {
		if d == dir {
			return true
		}
	}
	return false
}

func Clone(s *Solver) *Solver {
	ss := &Solver{}
	ss.ms = [][][]rune{}
	for _, row := range s.ms {
		newRow := [][]rune{}
		for _, rows := range row {
			newRows := []rune{}
			newRows = append(newRows, rows...)
			newRow = append(newRow, newRows)
		}
		ss.ms = append(ss.ms, newRow)
	}
	ss.x = s.x
	ss.y = s.y
	ss.dir = s.dir
	return ss
}

func (s *Solver) Stuck() bool {
	max := 100000
	curr := 0
	for curr < max {
		s.Moves()
		if s.OutOfBounds(s.x, s.y) {
			return false
		}
		curr++
	}
	return true
}

func (s *Solver) Moves() {
	x := s.x
	y := s.y
	switch s.dir {
	case '^':
		x -= 1
		if s.Contains(x, y, '#') {
			s.dir = '>'
			return
		}
	case 'v':
		x += 1
		if s.Contains(x, y, '#') {
			s.dir = '<'
			return
		}
	case '<':
		y -= 1
		if s.Contains(x, y, '#') {
			s.dir = '^'
			return
		}
	case '>':
		y += 1
		if s.Contains(x, y, '#') {
			s.dir = 'v'
			return
		}
	}

	s.x = x
	s.y = y
}

func (s *Solver) CountObsructions() int {
	v := 0
	for x := range s.ms {
		for y := range s.ms[x] {
			if x == s.x && y == s.y {
				continue
			}
			if s.Contains(x, y, '#') {
				continue
			}
			ss := Clone(s)
			ss.ms[x][y] = append(ss.ms[x][y], '#')
			if ss.Stuck() {
				v++
			}
		}
	}
	return v
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input)

	fmt.Println(s.CountVisited())
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input)

	fmt.Println(s.CountObsructions())
}

func main() {
	real := flag.Bool("real", false, "Whether to use the real input")
	runA := flag.Bool("a", false, "Run program a")
	runB := flag.Bool("b", false, "Run program b")

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
