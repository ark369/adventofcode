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

type Blizzard struct {
	x, y int
	dir  rune
}

type Wall struct {
	x, y int
}

type Cell struct {
	x, y      int
	wall      *Wall
	blizzards []*Blizzard
}

func (c *Cell) RemoveBlizzard(rem *Blizzard) {
	for i, b := range c.blizzards {
		if b == rem {
			c.blizzards[i] = c.blizzards[len(c.blizzards)-1]
			break
		}
	}
	// Chop the end off
	c.blizzards = c.blizzards[:len(c.blizzards)-1]
}

func (c *Cell) IsEmpty() bool {
	return c.wall == nil && len(c.blizzards) == 0
}

type Map struct {
	blizzards     []*Blizzard
	width, height int
	walls         []*Wall
	cells         [][]*Cell
}

type GameState struct {
	time, playerX, playerY int
	path                   string
}

func (g *GameState) PathFound(targetHeight int) bool {
	return g.playerY == targetHeight-1
}

func MakeGameStateAndMap(input []string) (*GameState, *Map) {
	g := &GameState{}
	g.time = 0
	m := &Map{}
	m.width = len(input[0])
	m.height = len(input)
	m.cells = make([][]*Cell, m.height)
	for i, l := range input {
		m.cells[i] = make([]*Cell, m.width)
		for j, c := range l {
			cell := &Cell{}
			cell.x = j
			cell.y = i
			m.cells[i][j] = cell
			switch c {
			case '#':
				wall := &Wall{j, i}
				m.walls = append(m.walls, wall)
				cell.wall = wall
			case '<':
				fallthrough
			case '^':
				fallthrough
			case '>':
				fallthrough
			case 'v':
				b := &Blizzard{j, i, c}
				m.blizzards = append(m.blizzards, b)
				cell.blizzards = append(cell.blizzards, b)
			case '.':
				if i == 0 {
					g.playerX = j
					g.playerY = 0
				}
			}
		}
	}
	return g, m
}

func (m *Map) NextBlizzardCell(b *Blizzard) *Cell {
	x := b.x
	y := b.y
	var cell *Cell
	switch b.dir {
	case '<':
		x -= 1
		cell = m.cells[y][x]
		if cell.wall != nil {
			cell = m.cells[y][m.width-2]
		}
	case '^':
		y -= 1
		cell = m.cells[y][x]
		if cell.wall != nil {
			cell = m.cells[m.height-2][x]
		}
	case '>':
		x += 1
		cell = m.cells[y][x]
		if cell.wall != nil {
			cell = m.cells[y][1]
		}
	case 'v':
		y += 1
		cell = m.cells[y][x]
		if cell.wall != nil {
			cell = m.cells[1][x]
		}
	}
	return cell
}

func (m *Map) Shift() {
	for _, b := range m.blizzards {
		startCell := m.cells[b.y][b.x]
		nextCell := m.NextBlizzardCell(b)
		nextCell.blizzards = append(nextCell.blizzards, b)
		startCell.RemoveBlizzard(b)
		b.x = nextCell.x
		b.y = nextCell.y
	}
}
func (g *GameState) PossibleMoves(m *Map) []string {
	moves := []string{}
	x := g.playerX
	y := g.playerY
	if y < m.height-1 && m.cells[y+1][x].IsEmpty() {
		moves = append(moves, "v")
	}
	if x < m.width-1 && m.cells[y][x+1].IsEmpty() {
		moves = append(moves, ">")
	}
	if y > 0 && m.cells[y-1][x].IsEmpty() {
		moves = append(moves, "^")
	}
	if x > 0 && m.cells[y][x-1].IsEmpty() {
		moves = append(moves, "<")
	}
	if m.cells[y][x].IsEmpty() {
		moves = append(moves, "+")
	}
	return moves
}

func (m *Map) String() string {
	s := ""
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			cell := m.cells[i][j]
			if cell.wall != nil {
				s += "#"
			} else if len(cell.blizzards) > 1 {
				s += fmt.Sprintf("%d", len(cell.blizzards))
			} else if len(cell.blizzards) == 1 {
				s += string(cell.blizzards[0].dir)
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	//s += fmt.Sprintf("\nPossible moves: %v\n\n", g.PossibleMoves())
	return s
}

func main() {
  // COULD NOT ACTUALLY SOLVE THE MAIN INPUT IN GO PLAYGROUND
  // Needs optimizations, some possibilities:
  // * pseudo DFS instead of current BFS, might need to memo-ize the maps per time
  // * More compact representation of current state
  // * More compact representation of map (just blizzards and walls?)
  // * Pruning of possible states?
  // * Identification of duplicate map states?

  //input := ReadInput()
	input := ReadFakeInput()

	initial, currMap := MakeGameStateAndMap(input)
	currTime := 0
	states := []*GameState{initial}
	currMap.Shift()

	counter := 0

	for len(states) > 0 {
		counter++
		g := states[0]
		if g.PathFound(currMap.height) {
			fmt.Printf("PATH FOUND: counter: %d, time: %d, path: %s\n", counter, g.time, g.path)
			break
		}
		states = states[1:]
		if g.time > currTime {
			currMap.Shift()
			currTime++
		}
		possibleMoves := g.PossibleMoves(currMap)
		for _, m := range possibleMoves {
			next := &GameState{}
			next.path = g.path + m
			next.time = currTime + 1
			x := g.playerX
			y := g.playerY
			switch m {
			case "+":
				next.playerX = x
				next.playerY = y
			case "<":
				next.playerX = x - 1
				next.playerY = y
			case "^":
				next.playerX = x
				next.playerY = y - 1
			case ">":
				next.playerX = x + 1
				next.playerY = y
			case "v":
				next.playerX = x
				next.playerY = y + 1
			}
			states = append(states, next)
		}
	}
	if len(states) == 0 {
		fmt.Printf("PATH NOT FOUND: counter: %d, time: %d\n", counter, currTime)
	}
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
