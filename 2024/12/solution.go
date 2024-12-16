package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
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
	garden  [][]rune
	visited [][]bool
}

func MakeSolver(input *bufio.Scanner) *Solver {
	s := &Solver{}
	s.garden = [][]rune{}
	s.visited = [][]bool{}

	for input.Scan() {
		l := input.Text()
		s.garden = append(s.garden, []rune(l))
		s.visited = append(s.visited, make([]bool, len(l)))
	}

	return s
}

type Pair struct {
	x, y int
}

type Fence struct {
	idx, other_idx    int
	horz, vert        bool
	outside_downright bool
}

func (s *Solver) GetCandidatesAndFencing(p Pair, target rune) ([]Pair, []Fence) {
	candidates := []Pair{}
	ret := []Pair{}
	fencing := []Fence{}
	if s.visited[p.x][p.y] {
		return ret, fencing
	}
	s.visited[p.x][p.y] = true
	candidates = append(candidates, Pair{p.x - 1, p.y})
	candidates = append(candidates, Pair{p.x + 1, p.y})
	candidates = append(candidates, Pair{p.x, p.y - 1})
	candidates = append(candidates, Pair{p.x, p.y + 1})
	fmt.Printf("GetCandidatesAndFencing x: %d, y: %d, target: %c\n", p.x, p.y, target)
	for _, c := range candidates {
		x := c.x
		y := c.y
		if x < 0 || y < 0 || x >= len(s.garden) || y >= len(s.garden[0]) {
			f := Fence{}
			if x == p.x {
				// Vertical fencing
				f.vert = true
				if y == p.y-1 {
					f.idx = y + 1
				} else {
					f.idx = y
					f.outside_downright = true
				}
				f.other_idx = x
			} else {
				// Horizontal fencing
				f.horz = true
				if x == p.x-1 {
					f.idx = x + 1
				} else {
					f.idx = x
					f.outside_downright = true
				}
				f.other_idx = y
			}
			fencing = append(fencing, f)
			continue
		}
		if s.garden[x][y] != target {
			f := Fence{}
			if x == p.x {
				// Vertical fencing
				f.vert = true
				if y == p.y-1 {
					f.idx = y + 1
				} else {
					f.idx = y
					f.outside_downright = true
				}
				f.other_idx = x
			} else {
				// Horizontal fencing
				f.horz = true
				if x == p.x-1 {
					f.idx = x + 1
				} else {
					f.idx = x
					f.outside_downright = true
				}
				f.other_idx = y
			}
			fencing = append(fencing, f)
			continue
		}
		ret = append(ret, c)
	}
	return ret, fencing
}

type FenceSide struct {
	x, y              int
	outside_downright bool
}

func AddSideToRanges(side int, outside_downright bool, ranges []FenceSide) []FenceSide {
	for i, r := range ranges {
		if side >= r.x && side <= r.y {
			// Already inside
			return ranges
		}
		if side < r.x {
			if side == r.x-1 && outside_downright == r.outside_downright {
				ranges[i] = FenceSide{r.x - 1, r.y, outside_downright}
				return ranges
			}
			ranges = slices.Insert(ranges, i, FenceSide{side, side, outside_downright})
			return ranges
		}
		if side == r.y+1 {
			if i == len(ranges)-1 {
				if outside_downright == r.outside_downright {
					ranges[i] = FenceSide{r.x, r.y + 1, outside_downright}
					return ranges
				} else {
					ranges = append(ranges, FenceSide{side, side, outside_downright})
					return ranges
				}
			}
			if next := ranges[i+1]; next.x == side+1 {
				if outside_downright == r.outside_downright {
					if outside_downright == next.outside_downright {
						// Combine all 3
						newRange := FenceSide{r.x, next.y, outside_downright}
						newRanges := append(ranges[:i], ranges[i+1:]...)
						newRanges[i] = newRange
						return newRanges
					} else {
						// Combine prev + curr
						ranges[i] = FenceSide{r.x, r.y + 1, outside_downright}
						return ranges
					}
				} else {
					if outside_downright == next.outside_downright {
						// Combine curr + next
						ranges[i+1] = FenceSide{r.x - 1, r.y, outside_downright}
						return ranges
					} else {
						// All 3 different
						ranges = slices.Insert(ranges, i+1, FenceSide{side, side, outside_downright})
						return ranges
					}
				}
			} else {
				if outside_downright == r.outside_downright {
					ranges[i] = FenceSide{r.x, r.y + 1, outside_downright}
					return ranges
				} else {
					ranges = slices.Insert(ranges, i+1, FenceSide{side, side, outside_downright})
					return ranges
				}
			}
		}
	}
	// Add to end
	ranges = append(ranges, FenceSide{side, side, outside_downright})
	return ranges
}

func (s *Solver) CalculateGarden(x, y int, isB bool) int {
	c := s.garden[x][y]
	candidates := []Pair{{x, y}}
	actual := []Pair{}
	perimeter := 0
	horzSidesMap := make(map[int][]Fence)
	vertSidesMap := make(map[int][]Fence)
	for len(candidates) > 0 {
		curr := candidates[0]
		candidates = candidates[1:]
		if s.visited[curr.x][curr.y] {
			continue
		}
		next, fencing := s.GetCandidatesAndFencing(curr, c)
		if !isB {
			perimeter += len(fencing)
		} else {
			for _, f := range fencing {
				if f.horz {
					horzSidesMap[f.idx] = append(horzSidesMap[f.idx], f)
				} else {
					vertSidesMap[f.idx] = append(vertSidesMap[f.idx], f)
				}
			}
		}
		candidates = append(candidates, next...)
		actual = append(actual, curr)
	}
	area := len(actual)
	if isB {
		fmt.Printf("target: %c\n", c)
		// WRONG Internal adjacent corners is not handled properly (see sample)
		for k, fences := range horzSidesMap {
			fmt.Printf("horz idx %d\n", k)
			sides := []FenceSide{}
			for _, f := range fences {
				sides = AddSideToRanges(f.other_idx, f.outside_downright, sides)
			}
			fmt.Println(sides)
			perimeter += len(sides)
		}
		for k, fences := range vertSidesMap {
			fmt.Printf("vert idx %d\n", k)
			sides := []FenceSide{}
			for _, f := range fences {
				sides = AddSideToRanges(f.other_idx, f.outside_downright, sides)
			}
			fmt.Println(sides)
			perimeter += len(sides)
		}
	}
	fmt.Printf("Target: %c, Area: %d, Perimeter: %d\n", c, area, perimeter)
	return area * perimeter
}

func (s *Solver) Calculate(isB bool) int {
	price := 0
	for x, row := range s.visited {
		for y, v := range row {
			if v {
				continue
			}
			price += s.CalculateGarden(x, y, isB)
		}
	}
	return price
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input)
	fmt.Println(s.Calculate(false))
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input)
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
