package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

type Solver struct {
	m [][]int
}

func MakeSolver(input *bufio.Scanner) *Solver {
	s := &Solver{}
	s.m = [][]int{}

	for input.Scan() {
		l := input.Text()
		row := []int{}
		for _, c := range l {
			i := Atoi(string(c))
			row = append(row, i)
		}
		s.m = append(s.m, row)
	}

	return s
}

type Pair struct {
	x, y int
}

func (s *Solver) CalcTrailheadScore(x, y, currHeight int) []Pair {
	if x < 0 || y < 0 || x >= len(s.m) || y >= len(s.m[0]) {
		return []Pair{}
	}
	if s.m[x][y] != currHeight {
		return []Pair{}
	}
	if currHeight == 9 {
		return []Pair{{x, y}}
	}

	p := []Pair{}

	p = append(p, s.CalcTrailheadScore(x-1, y, currHeight+1)...)
	p = append(p, s.CalcTrailheadScore(x+1, y, currHeight+1)...)
	p = append(p, s.CalcTrailheadScore(x, y-1, currHeight+1)...)
	p = append(p, s.CalcTrailheadScore(x, y+1, currHeight+1)...)
	return p
}

func (s *Solver) Calculate() int {
	sum := 0

	for x, row := range s.m {
		for y, c := range row {
			if c == 0 {
				pairs := s.CalcTrailheadScore(x, y, 0)
				m := make(map[Pair]bool)
				for _, p := range pairs {
					m[p] = true
				}
				score := len(m)
				sum += score
				fmt.Printf("Trailhead (%d, %d) has score %d\n", x, y, score)
			}
		}
	}

	return sum
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input)
	fmt.Println(s.Calculate())
}

func (s *Solver) CalcTrailheadRating(x, y, currHeight int) int {
	if x < 0 || y < 0 || x >= len(s.m) || y >= len(s.m[0]) {
		return 0
	}
	if s.m[x][y] != currHeight {
		return 0
	}
	if currHeight == 9 {
		return 1
	}

	r := 0

	r += s.CalcTrailheadRating(x-1, y, currHeight+1)
	r += s.CalcTrailheadRating(x+1, y, currHeight+1)
	r += s.CalcTrailheadRating(x, y-1, currHeight+1)
	r += s.CalcTrailheadRating(x, y+1, currHeight+1)
	return r
}

func (s *Solver) CalculateB() int {
	sum := 0

	for x, row := range s.m {
		for y, c := range row {
			if c == 0 {
				rating := s.CalcTrailheadRating(x, y, 0)
				sum += rating
				fmt.Printf("Trailhead (%d, %d) has rating %d\n", x, y, rating)
			}
		}
	}

	return sum
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input)
	fmt.Println(s.CalculateB())
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
