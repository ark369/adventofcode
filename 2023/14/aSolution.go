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

type Map struct {
	cells [][]rune
}

func (m *Map) AddRow(l string) {
	newRow := make([]rune, len(l))
	m.cells = append(m.cells, newRow)
	for i, c := range l {
		newRow[i] = c
	}
}

func (m *Map) RollNorth() {
	for col := 0; col < len(m.cells[0]); col++ {
		topEmptyRow := 0
		for row := 0; row < len(m.cells); row++ {
			c := m.cells[row][col]
			switch c {
			case 'O':
				m.cells[row][col] = '.'
				m.cells[topEmptyRow][col] = 'O'
				topEmptyRow++
			case '#':
				topEmptyRow = row + 1
			}
		}
	}
}

func (m *Map) GetTotalLoad() int {
	sum := 0
	for row := 0; row < len(m.cells); row++ {
		for col := 0; col < len(m.cells[0]); col++ {
			c := m.cells[row][col]
			if c == 'O' {
				sum += len(m.cells) - row
			}
		}
	}
	return sum
}

func (m *Map) String() string {
	str := ""
	for _, row := range m.cells {
		for _, c := range row {
			str += string(c)
		}
		str += "\n"
	}
	return str
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	m := &Map{}

	for _, l := range input {
		m.AddRow(l)
	}
	//fmt.Printf("%s\n", m)
	m.RollNorth()
	//fmt.Printf("%s\n", m)

	fmt.Printf("m.GetTotalLoad(): %d\n", m.GetTotalLoad())
}

func ReadFakeInput() []string {
	input := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
