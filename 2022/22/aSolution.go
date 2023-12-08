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

func Score(row, col int, facing rune) int {
	score := row*1000 + col*4
	// Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
	switch facing {
	case 's':
		score += 1
	case 'a':
		score += 2
	case 'w':
		score += 3
	}
	return score
}

type Node struct {
	wall       bool
	w, s, a, d *Node
	row, col   int
}

type Map struct {
	curr      *Node
	dir       rune
	bottomRow []*Node
}

func (m *Map) AddRow(l string, rowNum int) {
	if len(m.bottomRow) == 0 {
		m.bottomRow = make([]*Node, len(l))
	}
	newRow := make([]*Node, len(l))
	for i, c := range l {
		if c == ' ' {
			continue
		}
		n := &Node{}
		n.row = rowNum
		n.col = i
		if c == '#' {
			n.wall = true
		}
		if m.curr == nil {
			m.curr = n
			m.dir = 'd'
		}
		newRow[i] = n
		// left & right
		if i == 0 || newRow[i-1] == nil {
			n.a = n
			n.d = n
		} else {
			left := newRow[i-1]
			n.a = left
			n.d = left.d
			left.d.a = n
			left.d = n
		}
		// up & down
		if i >= len(m.bottomRow) || m.bottomRow[i] == nil {
			n.w = n
			n.s = n
		} else {
			up := m.bottomRow[i]
			n.w = up
			n.s = up.s
			up.s.w = n
			up.s = n
		}
	}
	m.bottomRow = newRow
}

func (m *Map) Move(n int) {
	switch m.dir {
	case 'a':
		for i := 0; i < n; i++ {
			if m.curr.a.wall {
				break
			}
			m.curr = m.curr.a
		}
	case 'd':
		for i := 0; i < n; i++ {
			if m.curr.d.wall {
				break
			}
			m.curr = m.curr.d
		}
	case 'w':
		for i := 0; i < n; i++ {
			if m.curr.w.wall {
				break
			}
			m.curr = m.curr.w
		}
	case 's':
		for i := 0; i < n; i++ {
			if m.curr.s.wall {
				break
			}
			m.curr = m.curr.s
		}
	}
}

func (m *Map) Turn(dir rune) {
	switch m.dir {
	case 'a':
		if dir == 'R' {
			m.dir = 'w'
		} else {
			m.dir = 's'
		}
	case 'd':
		if dir == 'R' {
			m.dir = 's'
		} else {
			m.dir = 'w'
		}
	case 'w':
		if dir == 'R' {
			m.dir = 'd'
		} else {
			m.dir = 'a'
		}
	case 's':
		if dir == 'R' {
			m.dir = 'a'
		} else {
			m.dir = 'd'
		}
	}
}

func Move(steps string, m *Map) {
	numSoFar := ""
	for _, c := range steps {
		if c >= '0' && c <= '9' {
			numSoFar += string(c)
		} else {
			num := Atoi(numSoFar)
			m.Move(num)
			m.Turn(c)
			numSoFar = ""
		}
	}
	if numSoFar != "" {
		num := Atoi(numSoFar)
		m.Move(num)
	}
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	m := &Map{}

	buildMap := true

	for i, l := range input {
		if l == "" {
			buildMap = false
			continue
		}
		if buildMap {
			m.AddRow(l, i)
		} else {
			Move(l, m)
		}
	}

	row := m.curr.row + 1
	col := m.curr.col + 1
	facing := m.dir
	fmt.Printf("score: %d\n", Score(row, col, facing))
}

func ReadFakeInput() []string {
	input := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
