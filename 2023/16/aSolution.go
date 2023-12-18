package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Cell struct {
	typ       rune
	energized bool
}

func MakeCell(typ rune) *Cell {
	c := &Cell{}
	c.typ = typ
	return c
}

type Map struct {
	cells [][]*Cell
}

func (m *Map) AddRow(l string) {
	row := make([]*Cell, len(l))
	for i, c := range l {
		cell := MakeCell(c)
		row[i] = cell
	}
	m.cells = append(m.cells, row)
}

type Seen struct {
	row, col int
	dir      Dir
}

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)

func (m *Map) Shine() {
	seen := make(map[Seen]bool)
	m.Light(0, 0, RIGHT, seen)
}

func (m *Map) Cols() int {
	return len(m.cells[0])
}

func (m *Map) Rows() int {
	return len(m.cells)
}

func (m *Map) Light(row, col int, dir Dir, seen map[Seen]bool) {
	s := Seen{row, col, dir}
	if _, ok := seen[s]; ok {
		return
	} else {
		seen[s] = true
	}
	c := m.cells[row][col]
	c.energized = true
	switch dir {
	case UP:
		switch c.typ {
		case '.':
			if row > 0 {
				m.Light(row-1, col, UP, seen)
			}
		case '\\':
			if col > 0 {
				m.Light(row, col-1, LEFT, seen)
			}
		case '/':
			if col < m.Cols()-1 {
				m.Light(row, col+1, RIGHT, seen)
			}
		case '|':
			if row > 0 {
				m.Light(row-1, col, UP, seen)
			}
		case '-':
			if col > 0 {
				m.Light(row, col-1, LEFT, seen)
			}
			if col < m.Cols()-1 {
				m.Light(row, col+1, RIGHT, seen)
			}
		}
	case DOWN:
		switch c.typ {
		case '.':
			if row < m.Rows()-1 {
				m.Light(row+1, col, DOWN, seen)
			}
		case '\\':
			if col < m.Cols()-1 {
				m.Light(row, col+1, RIGHT, seen)
			}
		case '/':
			if col > 0 {
				m.Light(row, col-1, LEFT, seen)
			}
		case '|':
			if row < m.Rows()-1 {
				m.Light(row+1, col, DOWN, seen)
			}
		case '-':
			if col < m.Cols()-1 {
				m.Light(row, col+1, RIGHT, seen)
			}
			if col > 0 {
				m.Light(row, col-1, LEFT, seen)
			}
		}
	case LEFT:
		switch c.typ {
		case '.':
			if col > 0 {
				m.Light(row, col-1, LEFT, seen)
			}
		case '\\':
			if row > 0 {
				m.Light(row-1, col, UP, seen)
			}
		case '/':
			if row < m.Rows()-1 {
				m.Light(row+1, col, DOWN, seen)
			}
		case '|':
			if row > 0 {
				m.Light(row-1, col, UP, seen)
			}
			if row < m.Rows()-1 {
				m.Light(row+1, col, DOWN, seen)
			}
		case '-':
			if col > 0 {
				m.Light(row, col-1, LEFT, seen)
			}
		}
	case RIGHT:
		switch c.typ {
		case '.':
			if col < m.Cols()-1 {
				m.Light(row, col+1, RIGHT, seen)
			}
		case '\\':
			if row < m.Rows()-1 {
				m.Light(row+1, col, DOWN, seen)
			}
		case '/':
			if row > 0 {
				m.Light(row-1, col, UP, seen)
			}
		case '|':
			if row < m.Rows()-1 {
				m.Light(row+1, col, DOWN, seen)
			}
			if row > 0 {
				m.Light(row-1, col, UP, seen)
			}
		case '-':
			if col < m.Cols()-1 {
				m.Light(row, col+1, RIGHT, seen)
			}
		}
	}
}

func (m *Map) NumEnergized() int {
	sum := 0
	for _, row := range m.cells {
		for _, cell := range row {
			if cell.energized {
				sum++
			}
		}
	}
	return sum
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	m := &Map{}

	for _, l := range input {
		m.AddRow(l)
	}
	m.Shine()

	fmt.Printf("m.NumEnergized(): %d", m.NumEnergized())
}

func ReadFakeInput() []string {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
