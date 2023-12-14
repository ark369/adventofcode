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

func (m *Map) RollWest() {
	for row := 0; row < len(m.cells); row++ {
		leftEmptyCol := 0
		for col := 0; col < len(m.cells[0]); col++ {
			c := m.cells[row][col]
			switch c {
			case 'O':
				m.cells[row][col] = '.'
				m.cells[row][leftEmptyCol] = 'O'
				leftEmptyCol++
			case '#':
				leftEmptyCol = col + 1
			}
		}
	}
}

func (m *Map) RollSouth() {
	numRows := len(m.cells)
	for col := 0; col < len(m.cells[0]); col++ {
		bottomEmptyRow := numRows - 1
		for row := 0; row < numRows; row++ {
			currRow := numRows - 1 - row
			c := m.cells[currRow][col]
			switch c {
			case 'O':
				m.cells[currRow][col] = '.'
				m.cells[bottomEmptyRow][col] = 'O'
				bottomEmptyRow--
			case '#':
				bottomEmptyRow = currRow - 1
			}
		}
	}
}

func (m *Map) RollEast() {
	numCols := len(m.cells[0])
	for row := 0; row < len(m.cells); row++ {
		rightEmptyCol := numCols - 1
		for col := 0; col < numCols; col++ {
			currCol := numCols - 1 - col
			c := m.cells[row][currCol]
			switch c {
			case 'O':
				m.cells[row][currCol] = '.'
				m.cells[row][rightEmptyCol] = 'O'
				rightEmptyCol--
			case '#':
				rightEmptyCol = currCol - 1
			}
		}
	}
}

func (m *Map) Cycle() {
	m.RollNorth()
	//fmt.Printf("%s\n", m)
	m.RollWest()
	//fmt.Printf("%s\n", m)
	m.RollSouth()
	//fmt.Printf("%s\n", m)
	m.RollEast()
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

func (m *Map) Sparse() string {
	str := ""
	for _, row := range m.cells {
		var last rune
		currTotal := 0
		for i, c := range row {
			if i == 0 {
				if c == 'O' {
					last = 'x'
				} else {
					last = '.'
				}
				currTotal = 1
			} else if c == 'O' {
				if last == 'x' {
					currTotal++
				} else if last == '.' {
					str += fmt.Sprintf("%s%d", string(last), currTotal)
					last = 'x'
					currTotal = 1
				}
			} else {
				if last == 'x' {
					str += fmt.Sprintf("%s%d", string(last), currTotal)
					last = '.'
					currTotal = 1
				} else {
					currTotal++
				}
			}
		}
		str += "\n"
	}
	return str
}

func BinaryAToInt64(a string) int64 {
	ret, _ := strconv.ParseInt(a, 2, 64)
	return ret
}

func (m *Map) SparseBinary() [][]int64 {
	str := ""
	ret := [][]int64{}
	for _, row := range m.cells {
		retRow := []int64{}
		for _, c := range row {
			if c == 'O' {
				str += "1"
			} else {
				str += "0"
			}
		}
		i := 0
		for i = 0; (i+1)*64 < len(str); i++ {
			retRow = append(retRow, BinaryAToInt64(str[i*64:(i+1)*64]))
		}

		retRow = append(retRow, BinaryAToInt64(str[(i)*64:]))
		//fmt.Printf("str: %s, retRow: %v\n", str, retRow)
		ret = append(ret, retRow)
		str = ""
	}
	return ret
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	m := &Map{}

	hsh := make(map[string]int)

	for _, l := range input {
		m.AddRow(l)
	}
	//fmt.Printf("%s\n", m)
	var i int
	cycleRepeat := 1
	cycleStart := 1
	totalCycles := 1000000000
	//totalCycles := 100
	for i = 1; i <= totalCycles; i++ {
		m.Cycle()
		key := fmt.Sprintf("%v", m.SparseBinary())
		//fmt.Printf("%s\n", key)
		if prev, ok := hsh[key]; ok {
			cycleRepeat = i - prev
			cycleStart = prev
			break

			//fmt.Printf("Repeat found: prev: %d, i: %d\n", prev, i)
			//hsh[key] = i
		} else {
			hsh[key] = i
		}
	}

	remainingCycles := (totalCycles - cycleStart) % cycleRepeat
	for i = 0; i < remainingCycles; i++ {
		m.Cycle()
	}
	fmt.Printf("totalCycles: %d, cycleStart: %d, cycleRepeat: %d, remainingCycles: %d\n", totalCycles, cycleStart, cycleRepeat, remainingCycles)

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
