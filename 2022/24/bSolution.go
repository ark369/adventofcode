package main

import (
	"fmt"
	"os"
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

type Blizzard struct {
	x, y int
	dir  rune
}

type Map struct {
	blizzards     []*Blizzard
	width, height int
	cells         [][]int
}

func (m *Map) String() string {
	s := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			cell := m.cells[y][x]
			if x == 0 || y == 0 || x == m.width-1 || y == m.height-1 {
				if cell == 1 {
					s += "#"
				} else if cell == 0 {
					s += "."
				} else {
					panic("unexpected blizzard in wall")
				}
				continue
			}
			if cell == 0 {
				s += "."
			} else {
				s += fmt.Sprintf("%d", cell)
			}
		}
		s += "\n"
	}
	//s += fmt.Sprintf("\nPossible moves: %v\n\n", g.PossibleMoves())
	return s
}

type GameState struct {
	time, playerX, playerY int
	end1, start1           bool
}

func (g *GameState) End1(targetHeight int) bool {
	return !g.end1 && g.playerY == targetHeight-1
}

func (g *GameState) Start1() bool {
	return g.end1 && !g.start1 && g.playerY == 0
}

func (g *GameState) PathFound(targetHeight int) bool {
	return g.end1 && g.start1 && g.playerY == targetHeight-1
}

func MakeGameStateAndMap(input []string) (*GameState, *Map) {
	g := &GameState{}
	g.time = 0
	m := &Map{}
	m.width = len(input[0])
	m.height = len(input)
	m.cells = make([][]int, m.height)

	for y, l := range input {
		m.cells[y] = make([]int, m.width)
		for x, c := range l {
			switch c {
			case '#':
				m.cells[y][x] = 1
			case '<':
				fallthrough
			case '^':
				fallthrough
			case '>':
				fallthrough
			case 'v':
				b := &Blizzard{x, y, c}
				m.blizzards = append(m.blizzards, b)
				m.cells[y][x]++
			case '.':
				if y == 0 {
					g.playerX = x
					g.playerY = 0
				}
			}
		}
	}
	return g, m
}

func (m *Map) NextBlizzardXY(b *Blizzard) (int, int) {
	x := b.x
	y := b.y
	switch b.dir {
	case '<':
		x -= 1
		if x == 0 {
			x = m.width - 2
		}
	case '^':
		y -= 1
		if y == 0 {
			y = m.height - 2
		}
	case '>':
		x += 1
		if x == m.width-1 {
			x = 1
		}
	case 'v':
		y += 1
		if y == m.height-1 {
			y = 1
		}
	}
	return x, y
}

func (m *Map) Clone() *Map {
	m2 := &Map{}
	m2.width = m.width
	m2.height = m.height
	for _, b := range m.blizzards {
		m2.blizzards = append(m2.blizzards, b)
	}

	m2.cells = make([][]int, m.height)
	for y, l := range m.cells {
		m2.cells[y] = make([]int, m.width)
		for x, c := range l {
			m2.cells[y][x] = c
		}
	}

	return m2
}

func (m2 *Map) Shift() *Map {
	m := m2.Clone()
	for _, b := range m.blizzards {
		x, y := m.NextBlizzardXY(b)
		m.cells[b.y][b.x]--
		b.x = x
		b.y = y
		m.cells[y][x]++
	}
	return m
}
func (g *GameState) PossibleMoves(m *Map) []rune {
	moves := []rune{}
	x := g.playerX
	y := g.playerY
	if !g.end1 || (g.end1 && g.start1) {
		if y < m.height-1 && m.cells[y+1][x] == 0 {
			moves = append(moves, 'v')
		}
		if x < m.width-1 && m.cells[y][x+1] == 0 {
			moves = append(moves, '>')
		}
		if m.cells[y][x] == 0 {
			moves = append(moves, '+')
		}
		if y > 0 && m.cells[y-1][x] == 0 {
			moves = append(moves, '^')
		}
		if x > 0 && m.cells[y][x-1] == 0 {
			moves = append(moves, '<')
		}
	} else {
		if x > 0 && m.cells[y][x-1] == 0 {
			moves = append(moves, '<')
		}
		if y > 0 && m.cells[y-1][x] == 0 {
			moves = append(moves, '^')
		}
		if m.cells[y][x] == 0 {
			moves = append(moves, '+')
		}
		if x < m.width-1 && m.cells[y][x+1] == 0 {
			moves = append(moves, '>')
		}
		if y < m.height-1 && m.cells[y+1][x] == 0 {
			moves = append(moves, 'v')
		}
	}
	return moves
}

func Traverse() {
	g := states[0]
	states = states[1:]
	num := len(maps)
	if g.time > num-1 {
		maps = append(maps, maps[num-1].Shift())
		maps[num-1] = nil
	}
	currMap := maps[g.time]
	if g.PathFound(currMap.height) {
		fmt.Printf("PATH FOUND: time: %d\n", g.time)
		os.Exit(0)
	}
	possibleMoves := g.PossibleMoves(currMap)
	immediateReturn := false
	for _, dir := range possibleMoves {
		x := g.playerX
		y := g.playerY
		g2 := &GameState{}
		g2.time = g.time + 1
		g2.end1 = g.end1
		g2.start1 = g.start1
		switch dir {
		case '<':
			x -= 1
		case '^':
			y -= 1
		case '>':
			x += 1
		case 'v':
			y += 1
		}
		g2.playerX = x
		g2.playerY = y
		if !g2.end1 && y == 0 && dir == '^' {
			continue
		}
		if g2.end1 && !g2.start1 && y == currMap.height-1 && dir == 'v' {
			continue
		}
		if g2.end1 && !g2.start1 && g2.Start1() {
			g2.start1 = true
			fmt.Printf("len(states): %d\n", len(states))
			states = []*GameState{}
			fmt.Printf("len(states): %d\n", len(states))
			fmt.Printf("START1 FOUND: time: %d\n", g.time)
			immediateReturn = true
		}
		if !g2.end1 && g2.End1(currMap.height) {
			g2.end1 = true
			fmt.Printf("len(states): %d\n", len(states))
			states = []*GameState{}
			fmt.Printf("len(states): %d\n", len(states))
			fmt.Printf("END1 FOUND: time: %d\n", g.time)
			immediateReturn = true
		}
		if _, ok := seen[g2.K()]; !ok {
			if g2.time > 330 {
				//fmt.Printf("g2: %v\n", g2)
			}
			seen[g2.K()] = true
			states = append(states, g2)
			if immediateReturn {
				return
			}
		}
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Print(g *GameState, m *Map) {
	s := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			player := false
			if g.playerX == x && g.playerY == y {
				s += "E"
				player = true
			}
			cell := m.cells[y][x]
			if x == 0 || y == 0 || x == m.width-1 || y == m.height-1 {
				if cell == 1 {
					if player {
						panic(fmt.Sprintf("Player collision with wall at x,y: %d, %d, time: %d", x, y, g.time))
					}
					s += "#"
				} else if cell == 0 && !player {
					s += "."
				}
				continue
			}
			if cell == 0 && !player {
				s += "."
			} else if cell > 0 {
				if player {
					panic(fmt.Sprintf("Player collision with blizzard at x,y: %d, %d, time: %d", x, y, g.time))
				}
				s += fmt.Sprintf("%d", cell)
			}
		}
		s += "\n"
	}
	s += "\n"
	//s += fmt.Sprintf("\nPossible moves: %v\n\n", g.PossibleMoves())
	fmt.Printf(s)

}

func (g *GameState) K() string {
	return fmt.Sprintf("%d,%d,%d", g.playerX, g.playerY, g.time)
}

var maps []*Map = []*Map{}
var startX, startY int
var states []*GameState = []*GameState{}
var seen map[string]bool

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	initial, currMap := MakeGameStateAndMap(input)
	startX = initial.playerX
	startY = initial.playerY
	maps = append(maps, currMap.Shift())
	states = append(states, initial)
	seen = make(map[string]bool)
	seen[initial.K()] = true
	//fmt.Printf("%s\n", currMap)
	//fmt.Printf("%s\n", maps[0])
	for len(states) > 0 {
		Traverse()
	}
	fmt.Printf("DONE\n")
}

func ReadFakeInput() []string {
	input := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
