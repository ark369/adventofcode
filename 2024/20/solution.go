package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
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
	m              [][]rune
	best           [][]float64
	is, js, ie, je int
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}

	for i := 0; input.Scan(); i++ {
		l := input.Text()
		row := []rune{}
		for j, c := range l {
			row = append(row, c)
			if c == 'S' {
				s.is = i
				s.js = j
			} else if c == 'E' {
				s.ie = i
				s.je = j
			}
		}
		s.m = append(s.m, row)
	}

	s.best = make([][]float64, len(s.m))
	for i := range s.best {
		s.best[i] = slices.Repeat([]float64{math.Inf(1)}, len(s.m[0]))
	}

	return s
}

type Node struct {
	i, j int
	val  float64
}

type Pair struct {
	i, j int
}

func (s *Solver) GetNext(curr Node) []Node {
	ret := []Node{}
	next := []Pair{
		{curr.i - 1, curr.j},
		{curr.i + 1, curr.j},
		{curr.i, curr.j - 1},
		{curr.i, curr.j + 1},
	}
	for _, n := range next {
		if s.m[n.i][n.j] == '#' {
			continue
		}
		if s.best[n.i][n.j] <= curr.val {
			continue
		}
		ret = append(ret, Node{n.i, n.j, curr.val + 1})
		s.best[n.i][n.j] = curr.val + 1
	}
	return ret
}

func (s *Solver) AdjacentTracks(i, j int) []Pair {
	ret := []Pair{}
	adj := []Pair{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	}
	for _, n := range adj {
		if n.i < 0 || n.j < 0 || n.i >= len(s.m) || n.j >= len(s.m[0]) {
			continue
		}
		if s.m[n.i][n.j] == '#' {
			continue
		}
		ret = append(ret, n)
	}
	return ret
}

func (s *Solver) AdjacentWalls(i, j int) []Pair {
	ret := []Pair{}
	adj := []Pair{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	}
	for _, n := range adj {
		if n.i < 0 || n.j < 0 || n.i >= len(s.m) || n.j >= len(s.m[0]) {
			continue
		}
		if s.m[n.i][n.j] != '#' {
			continue
		}
		ret = append(ret, n)
	}
	return ret
}

func (s *Solver) Adjacent(i, j int) []Pair {
	ret := []Pair{}
	adj := []Pair{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	}
	for _, n := range adj {
		if n.i < 0 || n.j < 0 || n.i >= len(s.m) || n.j >= len(s.m[0]) {
			continue
		}
		ret = append(ret, n)
	}
	return ret
}

func (s *Solver) Calculate(isB bool) int {
	nodes := []Node{{s.ie, s.je, 0}}
	s.best[s.ie][s.je] = 0
	for len(nodes) > 0 {
		curr := nodes[0]
		nodes = nodes[1:]
		next := s.GetNext(curr)
		nodes = append(nodes, next...)
	}
	cheats := 0
	for i := 1; i < len(s.m)-1; i++ {
		for j := 1; j < len(s.m[0])-1; j++ {
			if s.m[i][j] != '#' {
				continue
			}
			adj := s.AdjacentTracks(i, j)
			if len(adj) < 2 {
				continue
			}
			highest := math.Inf(-1)
			lowest := math.Inf(1)
			for _, a := range adj {
				if s.best[a.i][a.j] > highest {
					highest = s.best[a.i][a.j]
				}
				if s.best[a.i][a.j] < lowest {
					lowest = s.best[a.i][a.j]
				}
			}
			if highest-lowest-2 >= 100 {
				cheats++
			}
		}
	}
	return cheats
}

func (s *Solver) Highest(arr []Pair) float64 {
	highest := math.Inf(-1)
	for _, a := range arr {
		if s.best[a.i][a.j] > highest {
			highest = s.best[a.i][a.j]
		}
	}
	return highest
}

func (s *Solver) Lowest(arr []Pair) float64 {
	lowest := math.Inf(1)
	for _, a := range arr {
		if s.best[a.i][a.j] < lowest {
			lowest = s.best[a.i][a.j]
		}
	}
	return lowest
}

func (s *Solver) IsWall(i, j int) bool {
	return s.m[i][j] == '#'
}

func (s *Solver) CalculateB() int {
	nodes := []Node{{s.ie, s.je, 0}}
	s.best[s.ie][s.je] = 0
	for len(nodes) > 0 {
		curr := nodes[0]
		nodes = nodes[1:]
		next := s.GetNext(curr)
		nodes = append(nodes, next...)
	}
	//s.Print()
	cheats := make(map[int]int)
	for i := 0; i < len(s.m); i++ {
		for j := 0; j < len(s.m[0]); j++ {
			if s.IsWall(i, j) {
				continue
			}
			start := s.best[i][j]
			cheat_dist := make([][]int, len(s.m))
			for cdi := range cheat_dist {
				cheat_dist[cdi] = append(cheat_dist[cdi], slices.Repeat([]int{-1}, len(s.m[0]))...)
			}
			cheat_dist[i][j] = 0
			adj := s.Adjacent(i, j)
			for _, a := range adj {
				cheat_dist[a.i][a.j] = 1
			}
			for len(adj) > 0 {
				curr := adj[0]
				adj = adj[1:]
				curr_dist := cheat_dist[curr.i][curr.j]
				if curr_dist > 20 {
					continue
				}
				if !s.IsWall(curr.i, curr.j) {
					savings := int(start-s.best[curr.i][curr.j]) - curr_dist
					if savings >= 100 {
						//fmt.Printf("Start (%d, %d), End (%d, %d), Savings: %d\n", i, j, curr.i, curr.j, savings)
						cheats[savings]++
					}
				}
				next := s.Adjacent(curr.i, curr.j)
				for _, n := range next {
					if cheat_dist[n.i][n.j] >= 0 {
						continue
					}
					cheat_dist[n.i][n.j] = curr_dist + 1
					adj = append(adj, n)
				}
			}
		}
	}
	sum := 0
	keys := []int{}
	for k, _ := range cheats {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		v := cheats[k]
		fmt.Printf("There are %d cheats that save %d picoseconds.\n", v, k)
		sum += v
	}
	return sum
}

func (s *Solver) Print() {
	var sb strings.Builder
	for i := range s.m {
		for j := range s.m[0] {
			if s.m[i][j] == '#' {
				sb.WriteString("####")
			} else {
				sb.WriteString(fmt.Sprintf(" %02d ", int(s.best[i][j])))
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
	fmt.Println(s.Calculate(false))
	//s.Print()
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input, true)
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
