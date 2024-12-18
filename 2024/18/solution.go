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

type Solver struct {
	m          [][]bool
	size, fall int
	blocks     []Pair
}

func (s *Solver) Print() {
	var sb strings.Builder
	for _, row := range s.m {
		for _, b := range row {
			if b {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func MakeSolver(input *bufio.Scanner, isB bool, real bool) *Solver {
	s := &Solver{}

	var size int
	var fall int
	if real {
		size = 71
		fall = 1024
	} else {
		size = 7
		fall = 12
	}
	s.size = size
	s.fall = fall
	s.m = make([][]bool, size)
	for i := 0; i < size; i++ {
		s.m[i] = make([]bool, size)
	}

	for input.Scan() {
		l := input.Text()
		x := Atoi(strings.Split(l, ",")[0])
		y := Atoi(strings.Split(l, ",")[1])
		if fall > 0 {
			s.m[y][x] = true
			fall--
		}
		s.blocks = append(s.blocks, Pair{x, y})
	}

	return s
}

type Pair struct {
	x, y int
}

type Visit struct {
	x, y, val int
}

func (s *Solver) Next(c Visit, visits [][]int) []Visit {
	next := []Visit{}
	candidates := []Visit{
		{c.x - 1, c.y, c.val + 1},
		{c.x + 1, c.y, c.val + 1},
		{c.x, c.y - 1, c.val + 1},
		{c.x, c.y + 1, c.val + 1},
	}
	for _, candidate := range candidates {
		x := candidate.x
		y := candidate.y
		if x < 0 || y < 0 || x >= len(visits[0]) || y >= len(visits) {
			continue
		}
		if s.m[y][x] {
			continue
		}
		if visits[y][x] == 0 {
			next = append(next, candidate)
		}
	}
	return next
}

func (s *Solver) Calculate(isB bool) string {
	if isB {
		for {
			blocked := true
			next_block := s.blocks[s.fall]
			s.m[next_block.y][next_block.x] = true
			s.fall++
			//fmt.Printf("Fall: %d, block: (%d,%d)\n", s.fall, next_block.x, next_block.y)
			//s.Print()
			visits := [][]int{}
			for _, row := range s.m {
				visits = append(visits, make([]int, len(row)))
			}
			start := Visit{0, 0, 1}
			to_visit := []Visit{start}
			for len(to_visit) > 0 {
				//fmt.Println(len(to_visit))
				curr := to_visit[0]
				//fmt.Println(curr)
				to_visit = to_visit[1:]
				if visits[curr.y][curr.x] > 0 {
					continue
				}
				visits[curr.y][curr.x] = curr.val
				if curr.y == len(visits)-1 && curr.x == len(visits[0])-1 {
					blocked = false
					break
				}
				to_visit = append(to_visit, s.Next(curr, visits)...)
			}
			if blocked {
				return fmt.Sprintf("%d,%d", next_block.x, next_block.y)
			}
		}
	} else {
		visits := [][]int{}
		for _, row := range s.m {
			visits = append(visits, make([]int, len(row)))
		}
		start := Visit{0, 0, 1}
		to_visit := []Visit{start}
		for len(to_visit) > 0 {
			//fmt.Println(len(to_visit))
			curr := to_visit[0]
			//fmt.Println(curr)
			to_visit = to_visit[1:]
			if visits[curr.y][curr.x] > 0 {
				continue
			}
			visits[curr.y][curr.x] = curr.val
			if curr.y == len(visits)-1 && curr.x == len(visits[0])-1 {
				//fmt.Println(visits)
				return Itoa(curr.val - 1)
			}
			to_visit = append(to_visit, s.Next(curr, visits)...)
		}
		return ""
	}
}

func a(input *bufio.Scanner, real bool) {
	s := MakeSolver(input, false, real)
	fmt.Println(s.Calculate(false))
}

func b(input *bufio.Scanner, real bool) {
	s := MakeSolver(input, true, real)
	s.Print()
	fmt.Println(s.Calculate(true))
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
		a(input, *real)
	}
	if *runB {
		b(input, *real)
	}
}

func ReadInput(fileName string) *bufio.Scanner {
	input, err := os.Open(fileName)
	if err != nil {
		panic("could not open")
	}

	return bufio.NewScanner(input)
}
