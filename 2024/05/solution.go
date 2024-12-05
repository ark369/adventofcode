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

type Solver struct {
	disallowed []int
	rules      map[int][]int
}

func MakeSolver() *Solver {
	s := &Solver{}
	s.rules = make(map[int][]int)
	return s
}

func (s *Solver) AddRule(l string) {
	parts := strings.Split(l, "|")
	a := Atoi(parts[0])
	b := Atoi(parts[1])
	if _, ok := s.rules[b]; !ok {
		s.rules[b] = []int{}
	}
	s.rules[b] = append(s.rules[b], a)
}

func (s *Solver) IsCorrect(l string) bool {
	s.disallowed = make([]int, 100)
	parts := strings.Split(l, ",")
	valid := true
	for _, part := range parts {
		val := Atoi(part)
		if s.disallowed[val] > 0 {
			valid = false
			break
		}
		// Allowed, so update all the disallowed
		for _, newDisallow := range s.rules[val] {
			s.disallowed[newDisallow] = val
		}
	}
	return valid
}

func (s *Solver) IsCorrectInts(ints []int) bool {
	s.disallowed = make([]int, 100)
	valid := true
	for _, val := range ints {
		if s.disallowed[val] > 0 {
			valid = false
			break
		}
		// Allowed, so update all the disallowed
		for _, newDisallow := range s.rules[val] {
			s.disallowed[newDisallow] = val
		}
	}
	return valid
}

func (s *Solver) Solve(l string) int {
	parts := strings.Split(l, ",")

	valid := s.IsCorrect(l)
	if !valid {
		return 0
	} else {
		mid := (len(parts) - 1) / 2
		return Atoi(parts[mid])
	}
}

func (s *Solver) PartialFix(in []int) ([]int, bool) {
	s.disallowed = make([]int, 100)
	correct := true
	for i, val := range in {
		if d := s.disallowed[val]; d > 0 {
			for ii := 0; ii < i; ii++ {
				if in[ii] == d {
					in[ii] = val
					break
				}
			}
			in[i] = d
			correct = false
			break
		}
		// Allowed, so update all the disallowed
		for _, newDisallow := range s.rules[val] {
			s.disallowed[newDisallow] = val
		}
	}
	return in, correct
}

func (s *Solver) Fix(l string) int {
	//fmt.Printf("Update before fix:\n> %s\n\n", l)
	parts := strings.Split(l, ",")
	ints := make([]int, len(parts))
	for i, part := range parts {
		val := Atoi(part)
		ints[i] = val
	}

	for ok := false; !ok; {
		ints, ok = s.PartialFix(ints)
	}

	//fmt.Printf("Update after fix:\n> %v\n\n", ints)
	kv := ""
	for k, v := range s.rules {
		kv += fmt.Sprintf("%d: %v\n", k, v)
	}
	//fmt.Printf("rules:\n> %s\n\n", kv)
	mid := (len(parts) - 1) / 2
	return ints[mid]
}

func a(input *bufio.Scanner) {
	var n int
	s := MakeSolver()

	rules := true
	sum := 0

	for n = 0; input.Scan(); n++ {
		l := input.Text()
		if len(l) == 0 {
			rules = false
		} else if rules {
			s.AddRule(l)
		} else {
			sum += s.Solve(l)
		}
	}

	fmt.Println(sum)
}

func b(input *bufio.Scanner) {
	var n int
	s := MakeSolver()

	rules := true
	sum := 0

	for n = 0; input.Scan(); n++ {
		l := input.Text()
		if len(l) == 0 {
			rules = false
		} else if rules {
			s.AddRule(l)
		} else {
			if s.IsCorrect(l) {
				continue
			}

			sum += s.Fix(l)
		}
	}

	fmt.Println(sum)
}

func main() {
	real := flag.Bool("real", false, "Whether to use the real input")
	runA := flag.Bool("a", false, "Run program a")
	runB := flag.Bool("b", false, "Run program a")

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
