package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

type Solver struct {
	m                  [][]rune
	i, j, end_i, end_j int
	dir                rune
}

func (s *Solver) Print() {
	var sb strings.Builder
	for _, row := range s.m {
		for _, c := range row {
			sb.WriteRune(c)
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}
	s.m = [][]rune{}

	i := 0
	for input.Scan() {
		l := input.Text()
		row := []rune{}
		for j, c := range l {
			row = append(row, c)
			if c == 'S' {
				s.i = i
				s.j = j
			}
			if c == 'E' {
				s.end_i = i
				s.end_j = j
			}
		}
		s.m = append(s.m, row)
		i++
	}

	s.dir = '>'

	return s
}

type P struct {
	i, j  int
	dir   rune
	score int
	path  map[Pair]bool
}

func (p P) Clone() P {
	pp := P{}
	pp.i = p.i
	pp.j = p.j
	pp.dir = p.dir
	pp.score = p.score
	pp.path = make(map[Pair]bool)
	for k, v := range p.path {
		pp.path[k] = v
	}
	return pp
}

func (p P) LRS() (P, P, P) {
	l := p.Clone()
	l.score += 1000
	r := p.Clone()
	r.score += 1000
	f := p.Clone()
	f.score += 1
	switch p.dir {
	case '>':
		l.dir = '^'
		r.dir = 'v'
		f.j++
		f.path[Pair{f.i, f.j}] = true
	case '<':
		l.dir = 'v'
		r.dir = '^'
		f.j--
		f.path[Pair{f.i, f.j}] = true
	case '^':
		l.dir = '<'
		r.dir = '>'
		f.i--
		f.path[Pair{f.i, f.j}] = true
	case 'v':
		l.dir = '>'
		r.dir = '<'
		f.i++
		f.path[Pair{f.i, f.j}] = true
	}
	return l, r, f
}

func (p P) Key() string {
	return fmt.Sprintf("%d,%d,%c", p.i, p.j, p.dir)
}

func AddBestPaths(existing_paths, new_paths map[Pair]bool) map[Pair]bool {
	for k := range new_paths {
		existing_paths[k] = true
	}
	return existing_paths
}

func (s *Solver) Calculate(isB bool) int {
	best := make(map[string]int)
	best_score := math.Inf(1)
	best_paths := make(map[Pair]bool)
	start := P{s.i, s.j, s.dir, 0, make(map[Pair]bool)}
	start.path[Pair{s.i, s.j}] = true
	candidates := []P{start}
	best[start.Key()] = 0
	for len(candidates) > 0 {
		curr := candidates[0]
		candidates = candidates[1:]
		l, r, f := curr.LRS()
		if s.m[f.i][f.j] == 'E' {
			if f.score < int(best_score) {
				best_paths = make(map[Pair]bool)
				best_paths = AddBestPaths(best_paths, f.path)
				best_score = float64(f.score)
			} else if f.score == int(best_score) {
				best_paths = AddBestPaths(best_paths, f.path)
			}
		} else {
			if f.score <= int(best_score) {
				if s.m[f.i][f.j] != '#' {
					if prev, ok := best[f.Key()]; ok {
						if f.score <= prev {
							best[f.Key()] = f.score
							candidates = append(candidates, f)
						}
					} else {
						best[f.Key()] = f.score
						candidates = append(candidates, f)
					}
				}
			}
			if l.score <= int(best_score) {
				if prev, ok := best[l.Key()]; ok {
					if l.score <= prev {
						best[l.Key()] = l.score
						candidates = append(candidates, l)
					}
				} else {
					best[l.Key()] = l.score
					candidates = append(candidates, l)
				}
			}
			if r.score <= int(best_score) {
				if prev, ok := best[r.Key()]; ok {
					if r.score <= prev {
						best[r.Key()] = r.score
						candidates = append(candidates, r)
					}
				} else {
					best[r.Key()] = r.score
					candidates = append(candidates, r)
				}
			}
		}
	}

	if isB {
		var sb strings.Builder
		for i, row := range s.m {
			for j, c := range row {
				if best_paths[Pair{i, j}] {
					sb.WriteRune('O')
				} else {
					sb.WriteRune(c)
				}
			}
			sb.WriteRune('\n')
		}
		fmt.Println(sb.String())
		return len(best_paths)
	} else {
		return int(best_score)
	}
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
	s.Print()
	fmt.Println(s.Calculate(false))
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input, true)
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
