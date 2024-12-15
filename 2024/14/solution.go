package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

type Robot struct {
	x, y, vx, vy int
}

func MakeRobot(l string) Robot {
	r := Robot{}
	re := regexp.MustCompile("p=(.+),(.+) v=(.+),(.+)")
	matches := re.FindSubmatch([]byte(l))
	r.x = Atoi(string(matches[1]))
	r.y = Atoi(string(matches[2]))
	r.vx = Atoi(string(matches[3]))
	r.vy = Atoi(string(matches[4]))
	return r
}

type Solver struct {
	robots        []Robot
	width, height int
	m             [][]int
}

func MakeSolver(input *bufio.Scanner, isB bool, width, height int) *Solver {
	s := &Solver{}
	s.width = width
	s.height = height
	s.m = [][]int{}
	for i := 0; i < height; i++ {
		row := make([]int, width)
		s.m = append(s.m, row)
	}

	for input.Scan() {
		l := input.Text()
		s.robots = append(s.robots, MakeRobot(l))
	}

	return s
}

func (s *Solver) CountRobots(x1, x2, y1, y2 int) int {
	r := 0
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			r += s.m[y][x]
		}
	}
	return r
}

func (s *Solver) Print() {
	var sb strings.Builder
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			v := s.m[y][x]
			if v == 0 {
				sb.WriteRune(' ')
			} else {
				sb.WriteString(Itoa(v))
			}
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func (s *Solver) Calculate(isB bool) int {
	if !isB {
		steps := 100
		for _, r := range s.robots {
			x := (r.x + steps*r.vx) % s.width
			if x < 0 {
				x += s.width
			}
			y := (r.y + steps*r.vy) % s.height
			if y < 0 {
				y += s.height
			}
			s.m[y][x]++
		}
		s.Print()

		x_mid := (s.width - 1) / 2
		y_mid := (s.height - 1) / 2

		topleft := s.CountRobots(0, x_mid-1, 0, y_mid-1)
		topright := s.CountRobots(x_mid+1, s.width-1, 0, y_mid-1)
		bottomleft := s.CountRobots(0, x_mid-1, y_mid+1, s.height-1)
		bottomright := s.CountRobots(x_mid+1, s.width-1, y_mid+1, s.height-1)

		fmt.Printf("%d, %d, %d, %d\n", topleft, topright, bottomleft, bottomright)

		return topleft * topright * bottomleft * bottomright
	} else {
		mult := 103
		for steps := 1570; steps < 10000; steps += mult {
			s.m = [][]int{}
			fmt.Println(steps)
			for i := 0; i < s.height; i++ {
				row := make([]int, s.width)
				s.m = append(s.m, row)
				fmt.Printf("*")
			}
			fmt.Println("")
			fmt.Println("")
			for _, r := range s.robots {
				x := (r.x + steps*r.vx) % s.width
				if x < 0 {
					x += s.width
				}
				y := (r.y + steps*r.vy) % s.height
				if y < 0 {
					y += s.height
				}
				s.m[y][x]++
			}
			s.Print()
			fmt.Println("")
			fmt.Println("")
			fmt.Println("")
		}
		return 0
	}
}

func a(input *bufio.Scanner, real bool) {
	width := 11
	height := 7
	if real {
		width = 101
		height = 103
	}
	s := MakeSolver(input, false, width, height)
	fmt.Println(s.Calculate(false))
}

func b(input *bufio.Scanner, real bool) {
	width := 11
	height := 7
	if real {
		width = 101
		height = 103
	}
	s := MakeSolver(input, true, width, height)
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
