package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

type Arcade struct {
	a, b, prize Pair
}

type Solver struct {
	arcades []Arcade
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}

	var a Arcade
	state := 'a'
	ra := regexp.MustCompile("Button A: X\\+([0-9]+), Y\\+([0-9]+)")
	rb := regexp.MustCompile("Button B: X\\+([0-9]+), Y\\+([0-9]+)")
	rp := regexp.MustCompile("Prize: X=([0-9]+), Y=([0-9]+)")

	for input.Scan() {
		l := input.Text()
		switch state {
		case 'a':
			a = Arcade{}
			matches := ra.FindSubmatch([]byte(l))
			a.a = Pair{Atoi(string(matches[1])), Atoi(string(matches[2]))}
			state = 'b'
		case 'b':
			matches := rb.FindSubmatch([]byte(l))
			a.b = Pair{Atoi(string(matches[1])), Atoi(string(matches[2]))}
			state = 'p'
		case 'p':
			matches := rp.FindSubmatch([]byte(l))
			if isB {
				a.prize = Pair{Atoi(string(matches[1])) + 10000000000000, Atoi(string(matches[2])) + 10000000000000}
			} else {
				a.prize = Pair{Atoi(string(matches[1])), Atoi(string(matches[2]))}
			}
			s.arcades = append(s.arcades, a)
			state = 'n'
		case 'n':
			state = 'a'
		}
	}

	return s
}

type Pair struct {
	x, y int
}

func (s *Solver) Calculate(isB bool) int {
	tokens := 0

	// 94a + 22b = 8400
	// 34a + 67b = 5400
	//
	// 34(94a + 22b) - 94(34a + 67b) = 34(8400) - 94(5400)
	// 34(22b) - 94(67b) = 34(8400) - 94(5400)
	// b(34x22 - 94x67) = (34x8400 - 94x5400)
	// b = (34x8400 - 94x5400) / (34x22 - 94x67)
	// a = (8400 - 22b) / 94
	//
	// 94 = a.x
	// 22 = b.x
	// 34 = a.y
	// 67 = b.y
	// 8400 = p.x
	// 5400 = p.y
	//
	// b = (a.y * p.x - a.x * p.y) / (a.y * b.x - a.x * b.y)
	// a = (p.x - b.x * b) / a.x

	for _, arcade := range s.arcades {
		a := arcade.a
		b := arcade.b
		p := arcade.prize
		b_press := 1.0 * (a.y*p.x - a.x*p.y) / (a.y*b.x - a.x*b.y)
		a_press := 1.0 * (p.x - b.x*b_press) / a.x
		tok := 3*a_press + b_press
		x_actual := a_press*a.x + b_press*b.x
		y_actual := a_press*a.y + b_press*b.y
		fmt.Printf("prize: %v, a: %v, b: %v\n", p, a, b)
		fmt.Printf("actual: {%d %d}\n", x_actual, y_actual)
		fmt.Printf("a_press %d, b_press %d, tok %d\n", a_press, b_press, tok)
		if !isB {
			if a_press >= 100 || b_press >= 100 {
				fmt.Println("SKIPPED")
				continue
			}
		}
		if a_press < 0 || b_press < 0 || x_actual != p.x || y_actual != p.y {
			fmt.Println("SKIPPED")
			continue
		}
		tokens += tok
	}

	return tokens
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
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
