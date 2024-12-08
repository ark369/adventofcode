package main

import (
	"bufio"
	"flag"
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

func Itoa(i int) string {
	return strconv.Itoa(i)
}

type Pair struct {
	x, y int
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Solver struct {
	m             [][]rune
	antenna       map[rune][]Pair
	width, height int
}

func MakeSolver(input *bufio.Scanner) *Solver {
	s := &Solver{}

	s.m = [][]rune{}
	s.antenna = make(map[rune][]Pair)

	for x := 0; input.Scan(); x++ {
		row := []rune{}
		l := input.Text()
		for y, c := range l {
			row = append(row, c)
			if c != '.' {
				if _, ok := s.antenna[c]; !ok {
					s.antenna[c] = []Pair{}
				}
				s.antenna[c] = append(s.antenna[c], Pair{x, y})
			}
		}
		s.m = append(s.m, row)
	}
	s.height = len(s.m)
	s.width = len(s.m[0])

	return s
}

func (s *Solver) String() string {
	str := ""
	for _, row := range s.m {
		str += string(row) + "\n"
	}
	for k, v := range s.antenna {
		pairs := []string{}
		for _, p := range v {
			pairs = append(pairs, p.String())
		}
		str += fmt.Sprintf("%c: %s\n", k, strings.Join(pairs, ", "))
	}
	return str
}

func Valid(x, y, width, height int) bool {
	if x < 0 || y < 0 || x >= height || y >= width {
		return false
	}
	return true
}

func CalculateForTwoLocs(a, b Pair, width, height int) []Pair {
	pairs := []Pair{}
	cx := 2*a.x - b.x
	cy := 2*a.y - b.y
	if Valid(cx, cy, width, height) {
		pairs = append(pairs, Pair{cx, cy})
	}
	dx := 2*b.x - a.x
	dy := 2*b.y - a.y
	if Valid(dx, dy, width, height) {
		pairs = append(pairs, Pair{dx, dy})
	}
	return pairs
}

func CalculateAntinodes(locs []Pair, width int, height int) []Pair {
	antinodes := []Pair{}
	for i := 0; i < len(locs)-1; i++ {
		for j := i + 1; j < len(locs); j++ {
			li := locs[i]
			lj := locs[j]
			antinodes = append(antinodes, CalculateForTwoLocs(li, lj, width, height)...)
		}
	}
	return antinodes
}

func (s *Solver) FindAntinodes() int {
	antinodes := make(map[Pair]bool)
	for _, pairs := range s.antenna {
		candidates := CalculateAntinodes(pairs, s.width, s.height)
		for _, c := range candidates {
			antinodes[c] = true
		}
	}
	return len(antinodes)
}

func CalculateForTwoLocsWithHarmonics(a, b Pair, width, height int) []Pair {
	pairs := []Pair{}
	// (3, 4), (5, 5) -> (1, 3), (7, 6)
	// +2, +1 OR -2, -1
	dx := a.x - b.x
	dy := a.y - b.y
	// Positive increase
	currX := a.x + dx
	currY := a.y + dy
	for Valid(currX, currY, width, height) {
		pairs = append(pairs, Pair{currX, currY})
		currX += dx
		currY += dy
	}
	// Negative increase
	currX = a.x - dx
	currY = a.y - dy
	for Valid(currX, currY, width, height) {
		pairs = append(pairs, Pair{currX, currY})
		currX -= dx
		currY -= dy
	}
	// Add the sentinel
	pairs = append(pairs, a)

	return pairs
}

func CalculateAntinodesWithHarmonics(locs []Pair, width int, height int) []Pair {
	antinodes := []Pair{}
	for i := 0; i < len(locs)-1; i++ {
		for j := i + 1; j < len(locs); j++ {
			li := locs[i]
			lj := locs[j]
			antinodes = append(antinodes, CalculateForTwoLocsWithHarmonics(li, lj, width, height)...)
		}
	}
	return antinodes
}

func (s *Solver) FindAntinodesWithHarmonics() int {
	antinodes := make(map[Pair]bool)
	for k, pairs := range s.antenna {
		candidates := CalculateAntinodesWithHarmonics(pairs, s.width, s.height)
		for _, c := range candidates {
			antinodes[c] = true
		}
		s.PrintHarmonics(k, candidates)
	}
	return len(antinodes)
}

func (s *Solver) PrintHarmonics(k rune, antinodes []Pair) {
	str := ""
	for x, row := range s.m {
		for y := range row {
			p := Pair{x, y}
			printed := false
			for _, v := range s.antenna[k] {
				if v == p {
					str += string(k)
					printed = true
					break
				}
			}
			if !printed {
				for _, v := range antinodes {
					if v == p {
						str += "#"
						printed = true
						break
					}
				}
			}
			if !printed {
				str += "."
			}
		}
		str += "\n"
	}
	fmt.Println(str)
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input)
	fmt.Println(s)
	fmt.Println(s.FindAntinodes())
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input)
	fmt.Println(s)
	fmt.Println(s.FindAntinodesWithHarmonics())
}

func main() {
	real := flag.Bool("real", false, "Whether to use the real input")
	runA := flag.Bool("a", false, "Run program a")
	runB := flag.Bool("b", false, "Run program b")

	flag.Parse()

	fileName := "sample.txt"
	if *real {
		fileName = "input.txt"
	}

	input := ReadInput(fileName)

	if !*runA && !*runB {
		panic("Did not specify a or b")
	}

	if *runA && *runB {
		panic("Specified both a and b")
	}

	if *runA {
		a(input)
	}
	if *runB {
		b(input)
	}
}

func ReadInput(fileName string) *bufio.Scanner {
	input, err := os.Open(fileName)
	if err != nil {
		panic("could not open")
	}

	return bufio.NewScanner(input)
}
