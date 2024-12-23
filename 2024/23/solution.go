package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	connections map[string][]string
	pairs       map[string]bool
	raw_pairs   map[string]bool
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}
	s.connections = make(map[string][]string)
	s.pairs = make(map[string]bool)
	s.raw_pairs = make(map[string]bool)

	for input.Scan() {
		l := input.Text()
		parts := strings.Split(l, "-")
		a := parts[0]
		b := parts[1]
		s.connections[a] = append(s.connections[a], b)
		s.connections[b] = append(s.connections[b], a)
		s.pairs[a+"-"+b] = true
		s.pairs[b+"-"+a] = true
		if a < b {
			s.raw_pairs[a+"-"+b] = true
		} else {
			s.raw_pairs[b+"-"+a] = true
		}
	}

	return s
}

type Pair struct {
	x, y int
}

func Key(a, b, c string) string {
	var x, y, z string
	if a < b {
		if b < c {
			x = a
			y = b
			z = c
		} else if a < c {
			x = a
			y = c
			z = b
		} else {
			x = c
			y = a
			z = b
		}
	} else if a < c {
		x = b
		y = a
		z = c
	} else if b < c {
		x = b
		y = c
		z = a
	} else {
		x = c
		y = b
		z = a
	}
	return fmt.Sprintf("%s,%s,%s", x, y, z)
}

func (s *Solver) Calculate(isB bool) int {
	sets := 0
	checked := make(map[string]bool)
	for a, v := range s.connections {
		if len(v) < 2 {
			continue
		}
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				b := v[i]
				c := v[j]
				k := Key(a, b, c)
				if _, ok := checked[k]; ok {
					continue
				}
				checked[k] = true
				if _, ok := s.pairs[b+"-"+c]; ok {
					if a[0] == 't' || b[0] == 't' || c[0] == 't' {
						fmt.Println(k)
						sets++
					}
				}
			}
		}
	}
	return sets
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input, false)
	fmt.Println(s)
	fmt.Println(s.Calculate(false))
}

func (s *Solver) CalculateB() string {
	sorted_k := []string{}
	for k := range s.raw_pairs {
		sorted_k = append(sorted_k, k)
	}
	sort.Strings(sorted_k)

	networks := make(map[string][][]string)

	for _, pair := range sorted_k {
		fmt.Println(pair)
		parts := strings.Split(pair, "-")
		a := parts[0]
		b := parts[1]
		if _, ok := networks[a]; !ok {
			networks[a] = append(networks[a], []string{b})
			continue
		}
		inserted := false
		for i, network := range networks[a] {
			all_connected := true
			for _, c := range network {
				if _, ok := s.pairs[b+"-"+c]; !ok {
					all_connected = false
					break
				}
			}
			if all_connected {
				networks[a][i] = append(networks[a][i], b)
				inserted = true
				break
			}
		}
		if !inserted {
			networks[a] = append(networks[a], []string{b})
		}
		fmt.Println(networks)
	}
	largest := 0
	largest_str := ""
	for k, network := range networks {
		for _, v := range network {
			if len(v) > largest {
				largest = len(v)
				largest_str = fmt.Sprintf("%s,%s", k, strings.Join(v, ","))
			}
		}
	}
	return largest_str
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
