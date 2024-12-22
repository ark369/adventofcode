package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	secrets             []int
	sequence_to_bananas map[string]int
}

func MakeSolver(input *bufio.Scanner, isB bool) *Solver {
	s := &Solver{}
	s.sequence_to_bananas = make(map[string]int)

	for input.Scan() {
		l := input.Text()
		s.secrets = append(s.secrets, Atoi(l))
	}

	return s
}

type Pair struct {
	x, y int
}

func (s *Solver) NextSecret(curr int) int {
	times64 := curr * 64
	curr ^= times64
	curr %= 16777216

	div32 := curr / 32
	curr ^= div32
	curr %= 16777216

	times2048 := curr * 2048
	curr ^= times2048
	curr %= 16777216
	return curr
}

type Buf struct {
	buf [4]int
	idx int
}

func (b *Buf) Add(i int) {
	b.buf[b.idx] = i
	b.idx++
	if b.idx == 4 {
		b.idx = 0
	}
}

func (b *Buf) Key() string {
	first := b.buf[b.idx]
	second := b.buf[(b.idx+1)%4]
	third := b.buf[(b.idx+2)%4]
	fourth := b.buf[(b.idx+3)%4]
	return fmt.Sprintf("%d,%d,%d,%d", first, second, third, fourth)
}

func (s *Solver) Calculate(isB bool) int {
	for _, secret := range s.secrets {
		b := Buf{}
		seen := make(map[string]bool)

		for i := 0; i < 2000; i++ {
			prev := secret
			secret = s.NextSecret(secret)
			delta := (secret % 10) - (prev % 10)
			b.Add(delta)
			//fmt.Printf("%8d: %d (%d)\n", secret, secret%10, delta)
			if i >= 3 {
				k := b.Key()
				if _, ok := seen[k]; ok {
					continue
				}
				s.sequence_to_bananas[k] += (secret % 10)
				seen[k] = true
			}
		}
	}
	best_v := 0
	best_k := ""
	for k, v := range s.sequence_to_bananas {
		if v > best_v {
			best_v = v
			best_k = k
		}
	}
	fmt.Println(best_k)
	return best_v
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
