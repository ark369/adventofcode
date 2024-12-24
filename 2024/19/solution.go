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
	towels  map[rune][]string
	designs []string
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}
	s.towels = map[rune][]string{}

	towels := true

	for input.Scan() {
		l := input.Text()
		if l == "" {
			continue
		}
		if towels {
			t := strings.Split(l, ", ")
			towels = false
			for _, towel := range t {
				c := rune(towel[0])
				s.towels[c] = append(s.towels[c], towel)
			}
		} else {
			s.designs = append(s.designs, l)
		}
	}

	return s
}

type Pair struct {
	x, y int
}

func (s *Solver) GetNew(start rune) []string {
	if match, ok := s.towels[start]; !ok {
		return []string{}
	} else {
		return match
	}

}

func (s *Solver) Calculate(isB bool) int {
	possible := 0
	for _, d := range s.designs {
		streams := []string{}
		new_allowed := true
		//fmt.Println(d)
		fulfilled := false
		for i, c := range d {
			//fmt.Println(string(c))
			//fmt.Println(streams)
			if len(streams) == 0 && !new_allowed {
				//fmt.Println("K")
				break
			}
			next_streams := []string{}
			if new_allowed {
				new_allowed = false
				//fmt.Println("KK")
				new_streams := s.GetNew(c)
				streams = append(streams, new_streams...)
				//fmt.Println(streams)
			}
			for _, stream := range streams {
				//fmt.Println(stream)
				if rune(stream[0]) != c {
					//fmt.Println("KKK")
					continue
				}
				if len(stream) == 1 {
					//fmt.Println("KKKK")
					if i == len(d)-1 {
						//fmt.Println("KKKK 1")
						fulfilled = true
						break
					}
					//fmt.Println("KKKK 2")
					new_allowed = true
				} else {
					//fmt.Println("KKKK 3")
					next_streams = append(next_streams, stream[1:])
				}
			}
			if fulfilled {
				//fmt.Println("KKKKK")
				break
			}
			streams = next_streams
		}
		if fulfilled {
			//fmt.Println("fulfilled")
			possible++
		}
	}
	return possible
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
	fmt.Println(s)
	fmt.Println(s.Calculate(false))
}

type Stream struct {
	start, curr int
	str         string
}

func NewStream(str string, start int) *Stream {
	s := Stream{}
	s.start = start
	s.str = str
	return &s
}

func (s *Stream) Match(c rune) bool {
	if rune(s.str[s.curr]) == c {
		s.curr++
		return true
	}
	s.curr++
	return false
}

func (s *Stream) Done() bool {
	return s.curr == len(s.str)
}

func (s *Solver) CalculateB() int {
	possible := 0
	for _, d := range s.designs {
		memo := make(map[int]int)
		memo[0] = 1
		streams := []*Stream{}
		new_allowed := true
		for i, c := range d {
			if len(streams) == 0 && !new_allowed {
				break
			}
			next_streams := []*Stream{}
			if new_allowed {
				new_allowed = false
				new_streams := s.GetNew(c)
				for _, new_stream := range new_streams {
					streams = append(streams, NewStream(new_stream, i))
				}
			}
			for _, stream := range streams {
				if !stream.Match(c) {
					continue
				}
				if stream.Done() {
					if _, ok := memo[stream.start]; !ok {
						memo[stream.start] = 0
					}
					if _, ok := memo[i+1]; !ok {
						memo[i+1] = 0
					}
					memo[i+1] += memo[stream.start]
					new_allowed = true
				} else {
					next_streams = append(next_streams, stream)
				}
			}
			streams = next_streams
		}
		if v, ok := memo[len(d)]; ok {
			possible += v
		}
	}
	return possible
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
