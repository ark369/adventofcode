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

type Galaxy struct {
	id, row, col int
}

func (g *Galaxy) String() string {
	return fmt.Sprintf("%d: (%d, %d)", g.id, g.row, g.col)
}

func Expand(input []string) []string {
	ret := []string{}
	colGalaxySeen := make([]bool, len(input[0]))
	emptyRow := ""
	for i := 0; i < len(input[0]); i++ {
		emptyRow += "."
	}
	for _, l := range input {
		rowGalaxySeen := false
		ret = append(ret, l)
		for j, c := range l {
			if c == '#' {
				rowGalaxySeen = true
				colGalaxySeen[j] = true
			}
		}
		if !rowGalaxySeen {
			ret = append(ret, emptyRow)
		}
	}
	numSeen := 0
	for j, seen := range colGalaxySeen {
		if !seen {
			for i := 0; i < len(ret); i++ {
				insertPoint := j + numSeen
				ret[i] = ret[i][:insertPoint] + "." + ret[i][insertPoint:]
			}
			numSeen++
		}
	}
	return ret
}

func ExpandedDistance(a, b int, galaxiesSeen []bool, expansionSize int) int {
	if b < a {
		tmp := b
		b = a
		a = tmp
	}
	dist := 0
	for i := a; i < b; i++ {
		if galaxiesSeen[i] {
			dist += 1
		} else {
			dist += expansionSize
		}
	}
	return dist
}

func CalculateDistance(g1, g2 *Galaxy, rowGalaxiesSeen, colGalaxiesSeen []bool, expansionSize int) int {
	x := ExpandedDistance(g1.row, g2.row, rowGalaxiesSeen, expansionSize)
	y := ExpandedDistance(g1.col, g2.col, colGalaxiesSeen, expansionSize)
	return x + y
}

func CalculateDistances(galaxies []*Galaxy, rowGalaxiesSeen, colGalaxiesSeen []bool, expansionSize int) int {
	sum := 0

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := CalculateDistance(galaxies[i], galaxies[j], rowGalaxiesSeen, colGalaxiesSeen, expansionSize)
			sum += dist
			//fmt.Printf("Distance between %v and %v: %d\n", galaxies[i], galaxies[j], dist)
		}
	}

	return sum
}

func GetGalaxiesSeen(input []string) ([]bool, []bool) {
	rowGalaxiesSeen := make([]bool, len(input))
	colGalaxiesSeen := make([]bool, len(input[0]))
	for i, l := range input {
		for j, c := range l {
			if c == '#' {
				rowGalaxiesSeen[i] = true
				colGalaxiesSeen[j] = true
			}
		}
	}
	return rowGalaxiesSeen, colGalaxiesSeen
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	expansionSize := 1000000
	rowGalaxiesSeen, colGalaxiesSeen := GetGalaxiesSeen(input)

	g := []*Galaxy{}

	gId := 1

	for row, l := range input {
		for col, c := range l {
			if c == '#' {
				g = append(g, &Galaxy{gId, row, col})
				gId++
			}
		}
	}

	fmt.Printf("CalculateDistances(g): %v\n", CalculateDistances(g, rowGalaxiesSeen, colGalaxiesSeen, expansionSize))
}

func ReadFakeInput() []string {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
