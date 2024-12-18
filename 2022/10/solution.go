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
	X       int
	changes []int
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}

	s.X = 1

	for input.Scan() {
		l := input.Text()
		parts := strings.Split(l, " ")
		if parts[0] == "noop" {
			s.changes = append(s.changes, 0)
		} else {
			s.changes = append(s.changes, 0)
			s.changes = append(s.changes, Atoi(parts[1]))
		}
	}

	return s
}

type Pair struct {
	x, y int
}

func (s *Solver) Calculate(isB bool) int {
	if !isB {
		sum := 0
		for i := 0; i < 220; i++ {
			cycle := i + 1
			if (cycle-20)%40 == 0 {
				str := s.X * cycle
				fmt.Printf("cycle %d, x: %d, str: %d\n", cycle, s.X, str)
				sum += str
			}
			s.X += s.changes[i]
			//fmt.Printf("i: %d, change: %d\n", i, s.changes[i])
		}
		return sum
	} else {
		var sb strings.Builder
		for i := 0; i < 240; i++ {
			pos := i % 40
			fmt.Printf("i: %d, pos %d, s.X: %d\n", i, pos, s.X)
			if pos >= s.X-1 && pos <= s.X+1 {
				sb.WriteRune('#')
			} else {
				sb.WriteRune(' ')
			}
			if (i+1)%40 == 0 {
				sb.WriteRune('\n')
			}
			s.X += s.changes[i]
			//fmt.Printf("i: %d, change: %d\n", i, s.changes[i])
		}
		fmt.Println(sb.String())
		return 0
	}
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
	fmt.Println(s)
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
