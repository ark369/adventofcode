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
	m     [][]rune
	i, j  int
	moves []rune
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}

	grid := true
	i := 0

	for input.Scan() {
		l := input.Text()
		if l == "" {
			grid = false
			continue
		}

		if grid {
			row := []rune{}
			for j, c := range l {
				if isB {
					switch c {
					case '#':
						row = append(row, []rune{'#', '#'}...)
					case 'O':
						row = append(row, []rune{'[', ']'}...)
					case '.':
						row = append(row, []rune{'.', '.'}...)
					case '@':
						row = append(row, []rune{'@', '.'}...)
						s.i = i
						s.j = j * 2
					}
				} else {
					row = append(row, c)
					if c == '@' {
						s.i = i
						s.j = j
					}
				}
			}
			s.m = append(s.m, row)
			i++
		} else {
			s.moves = append(s.moves, []rune(l)...)
		}
	}

	return s
}

func (s *Solver) Print() {
	var sb strings.Builder
	for i := 0; i < len(s.m); i++ {
		for j := 0; j < len(s.m[0]); j++ {
			sb.WriteRune(s.m[i][j])
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

type Pair struct {
	x, y int
}

func (s *Solver) Move(dir rune) {
	switch dir {
	case '^':
		j := s.j
		for i := s.i - 1; i > 0; i-- {
			if s.m[i][j] == '.' {
				s.m[s.i][j] = '.'
				if i != s.i-1 {
					s.m[i][j] = 'O'
				}
				s.m[s.i-1][j] = '@'
				s.i--
				return
			}
			if s.m[i][j] == 'O' {
				continue
			}
			if s.m[i][j] == '#' {
				return
			}
		}
	case 'v':
		j := s.j
		for i := s.i + 1; i < len(s.m); i++ {
			if s.m[i][j] == '.' {
				s.m[s.i][j] = '.'
				if i != s.i+1 {
					s.m[i][j] = 'O'
				}
				s.m[s.i+1][j] = '@'
				s.i++
				return
			}
			if s.m[i][j] == 'O' {
				continue
			}
			if s.m[i][j] == '#' {
				return
			}
		}
	case '<':
		i := s.i
		for j := s.j - 1; j > 0; j-- {
			if s.m[i][j] == '.' {
				s.m[i][s.j] = '.'
				if j != s.j-1 {
					s.m[i][j] = 'O'
				}
				s.m[i][s.j-1] = '@'
				s.j--
				return
			}
			if s.m[i][j] == 'O' {
				continue
			}
			if s.m[i][j] == '#' {
				return
			}
		}
	case '>':
		i := s.i
		for j := s.j + 1; j < len(s.m[0]); j++ {
			if s.m[i][j] == '.' {
				s.m[i][s.j] = '.'
				if j != s.j+1 {
					s.m[i][j] = 'O'
				}
				s.m[i][s.j+1] = '@'
				s.j++
				return
			}
			if s.m[i][j] == 'O' {
				continue
			}
			if s.m[i][j] == '#' {
				return
			}
		}
	}
}

func (s *Solver) MoveB(dir rune) {
	m := [][]rune{}
	for i := 0; i < len(s.m); i++ {
		row := []rune{}
		for j := 0; j < len(s.m[0]); j++ {
			row = append(row, s.m[i][j])
		}
		m = append(m, row)
	}
	switch dir {
	case '^':
		pushing := make(map[int]rune)
		pushing[s.j] = '@'
		m[s.i][s.j] = '.'
		i := s.i - 1
		for len(pushing) > 0 {
			new_pushing := make(map[int]rune)
			for j, v := range pushing {
				switch m[i][j] {
				case '#':
					return
				case '[':
					new_pushing[j] = '['
					new_pushing[j+1] = ']'
					if _, ok := pushing[j+1]; !ok {
						m[i][j+1] = '.'
					}
				case ']':
					new_pushing[j] = ']'
					new_pushing[j-1] = '['
					if _, ok := pushing[j-1]; !ok {
						m[i][j-1] = '.'
					}
				}
				m[i][j] = v
			}
			pushing = new_pushing
			i--
		}
		s.i--
	case 'v':
		pushing := make(map[int]rune)
		pushing[s.j] = '@'
		m[s.i][s.j] = '.'
		i := s.i + 1
		for len(pushing) > 0 {
			new_pushing := make(map[int]rune)
			for j, v := range pushing {
				switch m[i][j] {
				case '#':
					return
				case '[':
					new_pushing[j] = '['
					new_pushing[j+1] = ']'
					if _, ok := pushing[j+1]; !ok {
						m[i][j+1] = '.'
					}
				case ']':
					new_pushing[j] = ']'
					new_pushing[j-1] = '['
					if _, ok := pushing[j-1]; !ok {
						m[i][j-1] = '.'
					}
				}
				m[i][j] = v
			}
			pushing = new_pushing
			i++
		}
		s.i++
	case '<':
		i := s.i
		for j := s.j - 1; j > 0; j-- {
			if s.m[i][j] == '.' {
				for new_j := j; new_j < s.j; new_j++ {
					s.m[i][new_j] = s.m[i][new_j+1]
				}
				s.m[i][s.j] = '.'
				s.j--
				return
			}
			if s.m[i][j] == '[' {
				continue
			}
			if s.m[i][j] == ']' {
				continue
			}
			if s.m[i][j] == '#' {
				return
			}
		}
	case '>':
		i := s.i
		for j := s.j + 1; j < len(s.m[0]); j++ {
			if s.m[i][j] == '.' {
				for new_j := j; new_j > s.j; new_j-- {
					s.m[i][new_j] = s.m[i][new_j-1]
				}
				s.m[i][s.j] = '.'
				s.j++
				return
			}
			if s.m[i][j] == '[' {
				continue
			}
			if s.m[i][j] == ']' {
				continue
			}
			if s.m[i][j] == '#' {
				return
			}
		}
	}
	s.m = m
}

func (s *Solver) Calculate(isB bool) int {
	if isB {
		for _, dir := range s.moves {
			s.MoveB(dir)
			//fmt.Printf("%c\n", dir)
			//s.Print()
		}
		sum := 0
		for i, row := range s.m {
			for j, c := range row {
				if c == '[' {
					fmt.Printf("100x%d + %d = %d\n", i, j, 100*i+j)
					sum += (i*100 + j)
				}
			}
		}
		return sum
	} else {
		for _, dir := range s.moves {
			s.Move(dir)
			//s.Print()
		}
		sum := 0
		for i, row := range s.m {
			for j, c := range row {
				if c == 'O' {
					//fmt.Printf("100x%d + %d = %d\n", i, j, 100*i+j)
					sum += (i*100 + j)
				}
			}
		}
		return sum
	}
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
	//s.Print()
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
