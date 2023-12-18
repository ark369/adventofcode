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

func (m *Map) Clone() *Map {
	newMap := &Map{}
	for _, row := range m.cells {
		newRow := make([]*Cell, len(row))
		for i, cell := range row {
			newCell := MakeCell(cell.typ)
			newRow[i] = newCell
		}
		newMap.cells = append(newMap.cells, newRow)
	}
	return newMap
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

func (d Dir) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		return fmt.Sprintf("%d", int(d))
	}
}

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

	base := &Map{}
	var m *Map
	best := 0

	for _, l := range input {
		base.AddRow(l)
	}
	//top
	for col := 0; col < base.Cols(); col++ {
		m = base.Clone()
		row := 0
		dir := DOWN
		seen := make(map[Seen]bool)
		m.Light(row, col, dir, seen)
		energized := m.NumEnergized()
		if energized > best {
			best = energized
			fmt.Printf("New best: %d, row: %d, col: %d, dir: %s\n", best, row, col, dir)
		}
	}
	//bottom
	for col := 0; col < base.Cols(); col++ {
		m = base.Clone()
		row := m.Rows() - 1
		dir := UP
		seen := make(map[Seen]bool)
		m.Light(row, col, dir, seen)
		energized := m.NumEnergized()
		if energized > best {
			best = energized
			fmt.Printf("New best: %d, row: %d, col: %d, dir: %s\n", best, row, col, dir)
		}
	}
	//left
	for row := 0; row < base.Rows(); row++ {
		m = base.Clone()
		col := m.Cols() - 1
		dir := LEFT
		seen := make(map[Seen]bool)
		m.Light(row, col, dir, seen)
		energized := m.NumEnergized()
		if energized > best {
			best = energized
			fmt.Printf("New best: %d, row: %d, col: %d, dir: %s\n", best, row, col, dir)
		}
	}
	//right
	for row := 0; row < base.Rows(); row++ {
		m = base.Clone()
		col := 0
		dir := RIGHT
		seen := make(map[Seen]bool)
		m.Light(row, col, dir, seen)
		energized := m.NumEnergized()
		if energized > best {
			best = energized
			fmt.Printf("New best: %d, row: %d, col: %d, dir: %s\n", best, row, col, dir)
		}
	}

	fmt.Printf("best: %d", best)
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
