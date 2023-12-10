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
	typ                       rune
	row, col                  int
	connections               []rune
	visited                   bool
	distance                  int
	mainLoop, inside, outside bool
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
			if p.typ == 'S' || p.typ == '|' || p.typ == 'L' || p.typ == 'J' {
				connections = append(connections, m.pipes[p.row-1][p.col])
			}
		case 's':
			if p.typ == 'S' || p.typ == '|' || p.typ == '7' || p.typ == 'F' {
				connections = append(connections, m.pipes[p.row+1][p.col])
			}
		case 'a':
			if p.typ == 'S' || p.typ == '-' || p.typ == '7' || p.typ == 'J' {
				connections = append(connections, m.pipes[p.row][p.col-1])
			}
		case 'd':
			if p.typ == 'S' || p.typ == '-' || p.typ == 'F' || p.typ == 'L' {
				connections = append(connections, m.pipes[p.row][p.col+1])
			}
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

func (m *Map) FillOutsideStartingAt(start *Pipe) {
	toVisit := []*Pipe{start}
	for len(toVisit) > 0 {
		p := toVisit[0]
		toVisit = toVisit[1:]
		if p.mainLoop || p.outside {
			continue
		}
		p.outside = true
		if p.row > 0 {
			n := m.pipes[p.row-1][p.col]
			if !n.mainLoop && !n.outside {
				toVisit = append(toVisit, n)
			}
		}
		if p.row < len(m.pipes)-1 {
			n := m.pipes[p.row+1][p.col]
			if !n.mainLoop && !n.outside {
				toVisit = append(toVisit, n)
			}
		}
		if p.col > 0 {
			n := m.pipes[p.row][p.col-1]
			if !n.mainLoop && !n.outside {
				toVisit = append(toVisit, n)
			}
		}
		if p.col < len(m.pipes[0])-1 {
			n := m.pipes[p.row][p.col+1]
			if !n.mainLoop && !n.outside {
				toVisit = append(toVisit, n)
			}
		}
	}
}

func (m *Map) FillOutsidePipe(p *Pipe, dir rune) {
	p.visited = true

	switch p.typ {
	case '|':
		if dir == 'a' {
			if p.col > 0 {
				m.FillOutsideStartingAt(m.pipes[p.row][p.col-1])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'a')
				case 'F':
					m.FillOutsidePipe(neighbour, 'a')
				case '7':
					m.FillOutsidePipe(neighbour, '#')
				case 'L':
					m.FillOutsidePipe(neighbour, 'a')
				case 'J':
					m.FillOutsidePipe(neighbour, '#')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		} else {
			if p.col < len(m.pipes[0])-1 {
				m.FillOutsideStartingAt(m.pipes[p.row][p.col+1])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'd')
				case 'F':
					m.FillOutsidePipe(neighbour, '#')
				case '7':
					m.FillOutsidePipe(neighbour, 'd')
				case 'L':
					m.FillOutsidePipe(neighbour, '#')
				case 'J':
					m.FillOutsidePipe(neighbour, 'd')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		}
	case '-':
		if dir == 'w' {
			if p.row > 0 {
				m.FillOutsideStartingAt(m.pipes[p.row-1][p.col])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '-':
					m.FillOutsidePipe(neighbour, 'w')
				case 'F':
					m.FillOutsidePipe(neighbour, 'a')
				case '7':
					m.FillOutsidePipe(neighbour, 'd')
				case 'L':
					m.FillOutsidePipe(neighbour, '#')
				case 'J':
					m.FillOutsidePipe(neighbour, '#')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		} else {
			if p.row < len(m.pipes)-1 {
				m.FillOutsideStartingAt(m.pipes[p.row+1][p.col])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '-':
					m.FillOutsidePipe(neighbour, 's')
				case 'F':
					m.FillOutsidePipe(neighbour, '#')
				case '7':
					m.FillOutsidePipe(neighbour, '#')
				case 'L':
					m.FillOutsidePipe(neighbour, 'a')
				case 'J':
					m.FillOutsidePipe(neighbour, 'd')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		}
	case 'L':
		if dir == 'a' {
			if p.col > 0 {
				m.FillOutsideStartingAt(m.pipes[p.row][p.col-1])
			}
			if p.row < len(m.pipes)-1 {
				m.FillOutsideStartingAt(m.pipes[p.row+1][p.col])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'a')
				case '-':
					m.FillOutsidePipe(neighbour, 's')
				case 'F':
					m.FillOutsidePipe(neighbour, 'a')
				case '7':
					m.FillOutsidePipe(neighbour, '#')
				case 'J':
					m.FillOutsidePipe(neighbour, 'd')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		} else {
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'd')
				case '-':
					m.FillOutsidePipe(neighbour, 'w')
				case 'F':
					m.FillOutsidePipe(neighbour, '#')
				case '7':
					m.FillOutsidePipe(neighbour, 'd')
				case 'J':
					m.FillOutsidePipe(neighbour, '#')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		}
	case 'J':
		if dir == 'd' {
			if p.col < len(m.pipes[0])-1 {
				m.FillOutsideStartingAt(m.pipes[p.row][p.col+1])
			}
			if p.row < len(m.pipes)-1 {
				m.FillOutsideStartingAt(m.pipes[p.row+1][p.col])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'd')
				case '-':
					m.FillOutsidePipe(neighbour, 's')
				case 'F':
					m.FillOutsidePipe(neighbour, '#')
				case '7':
					m.FillOutsidePipe(neighbour, 'd')
				case 'L':
					m.FillOutsidePipe(neighbour, 'a')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		} else {
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'a')
				case '-':
					m.FillOutsidePipe(neighbour, 'w')
				case 'F':
					m.FillOutsidePipe(neighbour, 'a')
				case '7':
					m.FillOutsidePipe(neighbour, '#')
				case 'L':
					m.FillOutsidePipe(neighbour, '#')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		}
	case '7':
		if dir == 'd' {
			if p.row > 0 {
				m.FillOutsideStartingAt(m.pipes[p.row-1][p.col])
			}
			if p.col < len(m.pipes[0])-1 {
				m.FillOutsideStartingAt(m.pipes[p.row][p.col+1])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'd')
				case '-':
					m.FillOutsidePipe(neighbour, 'w')
				case 'F':
					m.FillOutsidePipe(neighbour, 'a')
				case 'J':
					m.FillOutsidePipe(neighbour, 'd')
				case 'L':
					m.FillOutsidePipe(neighbour, '#')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		} else {
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'a')
				case '-':
					m.FillOutsidePipe(neighbour, 's')
				case 'F':
					m.FillOutsidePipe(neighbour, '#')
				case 'J':
					m.FillOutsidePipe(neighbour, '#')
				case 'L':
					m.FillOutsidePipe(neighbour, 'a')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		}
	case 'F':
		if dir == 'a' {
			if p.col > 0 {
				m.FillOutsideStartingAt(m.pipes[p.row][p.col-1])
			}
			if p.row > 0 {
				m.FillOutsideStartingAt(m.pipes[p.row-1][p.col])
			}
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'a')
				case '-':
					m.FillOutsidePipe(neighbour, 'w')
				case '7':
					m.FillOutsidePipe(neighbour, 'd')
				case 'J':
					m.FillOutsidePipe(neighbour, '#')
				case 'L':
					m.FillOutsidePipe(neighbour, 'a')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		} else {
			for _, neighbour := range m.GetConnections(p) {
				if neighbour.visited {
					continue
				}
				switch neighbour.typ {
				case '|':
					m.FillOutsidePipe(neighbour, 'd')
				case '-':
					m.FillOutsidePipe(neighbour, 's')
				case '7':
					m.FillOutsidePipe(neighbour, '#')
				case 'J':
					m.FillOutsidePipe(neighbour, 'd')
				case 'L':
					m.FillOutsidePipe(neighbour, '#')
				default:
					panic("Unexpected connecting pipe!")
				}
			}
		}
	}
}

func (m *Map) FillOutside() {
	found := false
	var startF *Pipe
	for row := 0; row < len(m.pipes); row++ {
		for col := 0; col < len(m.pipes[0]); col++ {
			p := m.pipes[row][col]
			if p.mainLoop {
				if p.typ != 'F' {
					panic(fmt.Sprintf("Assertion failed for type equal to F, instead found %s", string(p.typ)))
				}
				startF = p
				fmt.Printf("startF FOUND, startF.row, startF.col: (%d, %d)\n", startF.row, startF.col)
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	m.FillOutsidePipe(startF, 'a')
}

func (p *Pipe) ReplaceStart() {
	if p.typ != 'S' {
		panic("Tried to replace wrong pipe as start!")
	}
	c := p.connections
	if c[0] == 'a' {
		if c[1] == 'w' {
			p.typ = 'J'
		} else if c[1] == 's' {
			p.typ = '7'
		} else if c[1] == 'd' {
			p.typ = '-'
		}
	} else if c[0] == 'w' {
		if c[1] == 'a' {
			p.typ = 'J'
		} else if c[1] == 's' {
			p.typ = '|'
		} else if c[1] == 'd' {
			p.typ = 'L'
		}
	} else if c[0] == 's' {
		if c[1] == 'a' {
			p.typ = '7'
		} else if c[1] == 'w' {
			p.typ = '|'
		} else if c[1] == 'd' {
			p.typ = 'F'
		}
	} else if c[0] == 'd' {
		if c[1] == 'a' {
			p.typ = '-'
		} else if c[1] == 's' {
			p.typ = 'F'
		} else if c[1] == 'w' {
			p.typ = 'L'
		}
	}
}

func (m *Map) FillMainLoop() {
	start := m.pipes[m.startRow][m.startCol]
	start.mainLoop = true
	start.ReplaceStart()
	toVisit := []*Pipe{start}

	numMainLoop := 1

	for len(toVisit) > 0 {
		p := toVisit[0]
		toVisit = toVisit[1:]
		for _, neighbour := range m.GetConnections(p) {
			if !neighbour.mainLoop {
				neighbour.mainLoop = true
				//fmt.Printf("Adding (%d,%d) to mainLoop\n", neighbour.row, neighbour.col)
				toVisit = append(toVisit, neighbour)
				numMainLoop++
			}
		}
	}

	fmt.Printf("numMainLoop: %d\n", numMainLoop)
}

func (m *Map) GetEnclosed() int {
	enclosed := 0
	outside := 0
	mainLoop := 0

	m.FillMainLoop()
	m.FillOutside()

	for _, row := range m.pipes {
		for _, p := range row {
			if !p.outside && !p.mainLoop {
				enclosed++
			}
			if p.outside {
				outside++
			}
			if p.mainLoop {
				mainLoop++
			}
		}
	}
	fmt.Printf("outside: %d\n", outside)
	fmt.Printf("mainLoop: %d\n", mainLoop)

	return enclosed
}

func (m *Map) String() string {
	s := ""
	for _, row := range m.pipes {
		for _, col := range row {
			if col.mainLoop {
				s += col.String()
			} else if col.outside {
				s += "O"
			} else {
				s += "X"
			}
		}
		s += "\n"
	}
	return s
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

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

	fmt.Printf("GetEnclosed: %d\n", m.GetEnclosed())
	//fmt.Printf("%v\n", m)
}

func ReadFakeInput() []string {
	input := `..........
.F------7.
.|S----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........`

	return strings.Split(input, "\n")
}

func ReadFakeInput2() []string {
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
