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

type Pipe struct {
	typ         rune
	row, col    int
	connections []rune
	visited     bool
	distance    int
}

func (p *Pipe) AddConnection(c rune) {
	p.connections = append(p.connections, c)
}

func (p *Pipe) String() string {
	return string(p.typ)
}

type Map struct {
	startRow, startCol int
	pipes              [][]*Pipe
}

func (m *Map) ConnectDown(row, col int) {
	if row < len(m.pipes)-1 {
		m.pipes[row+1][col].AddConnection('w')
	}
}

func (m *Map) ConnectUp(row, col int) {
	if row > 0 {
		m.pipes[row-1][col].AddConnection('s')
	}
}

func (m *Map) ConnectLeft(row, col int) {
	if col > 0 {
		m.pipes[row][col-1].AddConnection('d')
	}
}

func (m *Map) ConnectRight(row, col int) {
	if col < len(m.pipes[0])-1 {
		m.pipes[row][col+1].AddConnection('a')
	}
}

func (m *Map) AddRow(row int, s string) {
	for col, c := range s {
		m.pipes[row][col].typ = c
		switch c {
		case '.':
		case 'S':
			m.startCol = col
			m.startRow = row
		case '|':
			m.ConnectUp(row, col)
			m.ConnectDown(row, col)
		case '-':
			m.ConnectLeft(row, col)
			m.ConnectRight(row, col)
		case 'L':
			m.ConnectUp(row, col)
			m.ConnectRight(row, col)
		case 'J':
			m.ConnectUp(row, col)
			m.ConnectLeft(row, col)
		case '7':
			m.ConnectDown(row, col)
			m.ConnectLeft(row, col)
		case 'F':
			m.ConnectDown(row, col)
			m.ConnectRight(row, col)
		}
	}
}

func (m *Map) GetConnections(p *Pipe) []*Pipe {
	connections := []*Pipe{}

	for _, c := range p.connections {
		switch c {
		case 'w':
			connections = append(connections, m.pipes[p.row-1][p.col])
		case 's':
			connections = append(connections, m.pipes[p.row+1][p.col])
		case 'a':
			connections = append(connections, m.pipes[p.row][p.col-1])
		case 'd':
			connections = append(connections, m.pipes[p.row][p.col+1])
		}
	}
	return connections
}

func (m *Map) GetFarthest() int {
	start := m.pipes[m.startRow][m.startCol]
	start.visited = true
	toVisit := []*Pipe{start}
	maxDistance := 0

	for len(toVisit) > 0 {
		p := toVisit[0]
		toVisit = toVisit[1:]
		if p.distance > maxDistance {
			maxDistance = p.distance
		}
		for _, neighbour := range m.GetConnections(p) {
			if !neighbour.visited {
				neighbour.visited = true
				neighbour.distance = p.distance + 1
				toVisit = append(toVisit, neighbour)
			}
		}
	}

	return maxDistance
}

func (m *Map) GetEnclosed() int {
	enclosed := 0

	return enclosed
}

func (m *Map) String() string {
	s := ""
	for i, row := range m.pipes {
		for j, col := range row {
			if i == m.startRow && j == m.startCol {
				s += "S"
			} else if col == nil {
				s += "."
			} else {
				s += col.String()
			}
		}
		s += "\n"
	}
	return s
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	m := &Map{}
	m.pipes = make([][]*Pipe, len(input))
	for row := 0; row < len(m.pipes); row++ {
		m.pipes[row] = make([]*Pipe, len(input[0]))
		for col := 0; col < len(m.pipes[row]); col++ {
			p := &Pipe{}
			p.row = row
			p.col = col
			m.pipes[row][col] = p
		}
	}

	for i, l := range input {
		m.AddRow(i, l)
	}

	//fmt.Printf("%v\n", m)
	fmt.Printf("GetEnclosed: %d\n", m.GetEnclosed())
}

func ReadFakeInput() []string {
	input := `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
