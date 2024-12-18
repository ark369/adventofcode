package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
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
	A, B, C int
	program []int
	adr     int
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}

	input.Scan()
	s.A = Atoi(strings.Split(input.Text(), ": ")[1])
	input.Scan()
	s.B = Atoi(strings.Split(input.Text(), ": ")[1])
	input.Scan()
	s.C = Atoi(strings.Split(input.Text(), ": ")[1])
	input.Scan()
	input.Scan()
	nums := strings.Split(strings.Split(input.Text(), ": ")[1], ",")
	for _, n := range nums {
		s.program = append(s.program, Atoi(n))
	}
	return s
}

type Pair struct {
	x, y int
}

func (s *Solver) Next() (int, int) {
	return s.program[s.adr], s.program[s.adr+1]
}

func (s *Solver) Combo(operand int) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return s.A
	case 5:
		return s.B
	case 6:
		return s.C
	case 7:
		panic("NOT VALID")
	}
	panic("NOT VALID operand")
}

func (s *Solver) dv(operand int) int {
	combo := s.Combo(operand)
	denom := int(math.Exp2(float64(combo)))
	return s.A / denom
}

func (s *Solver) Calculate(isB bool) string {
	outs := []string{}
	for s.adr < len(s.program) {
		opcode, operand := s.Next()
		switch opcode {
		case 0:
			val := s.dv(operand)
			s.A = val
			s.adr += 2
		case 1:
			val := s.B ^ operand
			s.B = val
			s.adr += 2
		case 2:
			combo := s.Combo(operand)
			val := combo % 8
			s.B = val
			s.adr += 2
		case 3:
			if s.A == 0 {
				s.adr += 2
			} else {
				s.adr = operand
			}
		case 4:
			val := s.B ^ s.C
			s.B = val
			s.adr += 2
		case 5:
			combo := s.Combo(operand)
			val := combo % 8
			outs = append(outs, Itoa(val))
			s.adr += 2
		case 6:
			val := s.dv(operand)
			s.B = val
			s.adr += 2
		case 7:
			val := s.dv(operand)
			s.C = val
			s.adr += 2
		}
	}
	return strings.Join(outs, ",")
}

func (s *Solver) State() string {
	return fmt.Sprintf("%d,%d,%d", s.A, s.B, s.C)
}

func (s *Solver) CalculateToTarget(final_target []int) int {
	var outs []int
	var outa []string
	curr := 0
	for idx := 0; idx < len(final_target); idx++ {
		target := final_target[len(final_target)-idx-1:]
		for i := 0; i < 800; i++ {
			start := curr*8 + i
			s.A = start
			s.B = 0
			s.C = 0
			s.adr = 0
			outs = []int{}
			outa = []string{}
			valid := true
			for s.adr < len(s.program) && valid {
				opcode, operand := s.Next()
				//fmt.Printf("lowest: %d, state: %s, outs: %s, ops (%d, %d)\n", lowest, s.State(), strings.Join(outa, ","), opcode, operand)
				switch opcode {
				case 0:
					val := s.dv(operand)
					s.A = val
					s.adr += 2
				case 1:
					val := s.B ^ operand
					s.B = val
					s.adr += 2
				case 2:
					combo := s.Combo(operand)
					val := combo % 8
					s.B = val
					s.adr += 2
				case 3:
					if s.A == 0 {
						s.adr += 2
					} else {
						s.adr = operand
					}
				case 4:
					val := s.B ^ s.C
					s.B = val
					s.adr += 2
				case 5:
					combo := s.Combo(operand)
					val := combo % 8
					outs = append(outs, val)
					outa = append(outa, Itoa(val))
					if len(outs) > len(target) {
						valid = false
						break
					}
					if !slices.Equal(outs, target[:len(outs)]) {
						valid = false
						break
					}
					s.adr += 2
				case 6:
					val := s.dv(operand)
					s.B = val
					s.adr += 2
				case 7:
					val := s.dv(operand)
					s.C = val
					s.adr += 2
				}
			}
			if slices.Equal(outs, target) {
				curr = start
				fmt.Printf("%d: %s\n", curr, strings.Join(outa, ","))
			}
		}
	}
	// 136902133483675 TOO HIGH
	return curr
}

func (s *Solver) CalculateB(isB bool) int {
	var outs []int
	var outa []string
	//35184000000000
	lowest := 35184372088831
	lowest += 1000000
	//        136902150000000
	i := 0
	for i < 1000000 {
		i++
		//seen := make(map[string]bool)
		s.A = lowest
		s.B = 0
		s.C = 0
		s.adr = 0
		outs = []int{}
		outa = []string{}
		valid := true
		for s.adr < len(s.program) && valid {
			opcode, operand := s.Next()
			//fmt.Printf("lowest: %d, state: %s, outs: %s, ops (%d, %d)\n", lowest, s.State(), strings.Join(outa, ","), opcode, operand)
			switch opcode {
			case 0:
				val := s.dv(operand)
				s.A = val
				s.adr += 2
			case 1:
				val := s.B ^ operand
				s.B = val
				s.adr += 2
			case 2:
				combo := s.Combo(operand)
				val := combo % 8
				s.B = val
				s.adr += 2
			case 3:
				if s.A == 0 {
					s.adr += 2
				} else {
					// if _, ok := seen[s.State()]; !ok {
					// 	//fmt.Printf("state: %s, lowest: %d\n", s.State(), lowest)
					// 	seen[s.State()] = true
					// 	s.adr = operand
					// } else {
					// 	valid = false
					// 	break
					// }
					s.adr = operand
				}
			case 4:
				val := s.B ^ s.C
				s.B = val
				s.adr += 2
			case 5:
				combo := s.Combo(operand)
				val := combo % 8
				outs = append(outs, val)
				outa = append(outa, Itoa(val))
				if len(outs) > len(s.program) {
					valid = false
					break
				}
				if !slices.Equal(outs, s.program[:len(outs)]) {
					valid = false
					break
				}
				s.adr += 2
			case 6:
				val := s.dv(operand)
				s.B = val
				s.adr += 2
			case 7:
				val := s.dv(operand)
				s.C = val
				s.adr += 2
			}
		}
		fmt.Printf("%d: %s\n", lowest, strings.Join(outa, ","))
		if slices.Equal(outs, s.program) {
			break
		}
		lowest += 1
	}
	// 136902133483675 TOO HIGH
	return lowest
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
	fmt.Println(s)
	fmt.Println(s.Calculate(false))
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input, true)
	//fmt.Println(s.CalculateB(true))
	fmt.Println(s.CalculateToTarget(s.program))
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
