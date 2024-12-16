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
	stones []int
}

func MakeSolver(input *bufio.Scanner) *Solver {
	s := &Solver{}
	s.stones = []int{}

	for input.Scan() {
		l := input.Text()
		parts := strings.Split(l, " ")
		for _, a := range parts {
			i := Atoi(a)
			s.stones = append(s.stones, i)
		}
	}

	return s
}

func Blink(val, count int) []int {
	//fmt.Printf("Blink %d %d\n", val, count)
	next := []int{}
	if val == 0 {
		//fmt.Println("0 becomes 1")
		next = append(next, 1)
	} else if a := Itoa(val); len(a)%2 == 0 {
		l := len(a) / 2
		left := Atoi(a[:l])
		right := Atoi(a[l:])
		next = append(next, left)
		next = append(next, right)
		//fmt.Printf("%d becomes %d, %d\n", val, left, right)
	} else {
		//fmt.Printf("%d becomes %d\n", val, val*2024)
		next = append(next, val*2024)
	}
	if count == 1 {
		return next
	} else {
		ret := []int{}
		for _, n := range next {
			ret = append(ret, Blink(n, count-1)...)
		}
		return ret
	}
}

func (s *Solver) Calculate() int {
	sum := 0

	str := ""
	for _, i := range s.stones {
		newStones := Blink(i, 25)
		for _, ii := range newStones {
			str += Itoa(ii) + " "
		}
		sum += len(newStones)
	}
	//fmt.Println(str)

	return sum
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input)
	fmt.Println(s.Calculate())
}

func AddToMemo(memo map[int]map[int][]int, ancestors map[int]int, vals []int) {
	for k, v := range ancestors {
		if _, ok := memo[k]; !ok {
			memo[k] = make(map[int][]int)
		}
		// if _, ok := memo[k][v]; ok {
		// 	return
		// }
		for _, val := range vals {
			if _, ok := memo[k][v]; !ok {
				memo[k][v] = []int{val}
			} else {
				memo[k][v] = append(memo[k][v], val)
			}
		}
	}
}

func MemoStr(memo map[int]map[int][]int) string {
	str := ""
	for k, counts := range memo {
		str += fmt.Sprintf("%d:\n", k)
		for c, vals := range counts {
			str += fmt.Sprintf("  %d: %v\n", c, vals)
		}
	}
	return str
}

func NextAncestors(ancestors map[int]int) map[int]int {
	next := make(map[int]int)
	for k, v := range ancestors {
		next[k] = v + 1
	}
	return next
}

func BlinkWithMemo(val, count int, memo map[int]map[int][]int, ancestors map[int]int) []int {
	//fmt.Printf("Blink %d %d\n", val, count)
	if a, ok := memo[val]; ok {
		if b, ok := a[count]; ok {
			AddToMemo(memo, ancestors, b)
			return b
		}
	}
	next := []int{}
	if val == 0 {
		//fmt.Println("0 becomes 1")
		next = append(next, 1)
	} else if a := Itoa(val); len(a)%2 == 0 {
		l := len(a) / 2
		left := Atoi(a[:l])
		right := Atoi(a[l:])
		next = append(next, left)
		next = append(next, right)
		//fmt.Printf("%d becomes %d, %d\n", val, left, right)
	} else {
		//fmt.Printf("%d becomes %d\n", val, val*2024)
		next = append(next, val*2024)
	}
	AddToMemo(memo, ancestors, next)
	if count == 1 {
		fmt.Printf("memo after %d\n%s\n\n", count, MemoStr(memo))
		return next
	} else {
		ret := []int{}
		for _, n := range next {
			nextAncestors := NextAncestors(ancestors)
			nextAncestors[n] = 1
			nnext := BlinkWithMemo(n, count-1, memo, nextAncestors)
			ret = append(ret, nnext...)
		}
		fmt.Printf("memo after %d\n%s\n\n", count, MemoStr(memo))
		return ret
	}
}

func NumStones(stone int, blinks int, memo map[int]map[int]int) int {
	if blinks == 0 {
		return 1
	}
	if v, ok := memo[stone][blinks]; ok {
		return v
	}
	var v int
	if stone == 0 {
		v = NumStones(1, blinks-1, memo)
	} else if a := Itoa(stone); len(a)%2 == 0 {
		l := len(a) / 2
		left := Atoi(a[:l])
		right := Atoi(a[l:])
		v = NumStones(left, blinks-1, memo) + NumStones(right, blinks-1, memo)
	} else {
		v = NumStones(stone*2024, blinks-1, memo)
	}
	if _, ok := memo[stone]; !ok {
		memo[stone] = make(map[int]int)
	}
	memo[stone][blinks] = v
	return v
}

func (s *Solver) CalculateB() int {
	total_count := 75

	sum := 0
	memo := make(map[int]map[int]int)
	for _, stone := range s.stones {
		sum += NumStones(stone, total_count, memo)
	}

	return sum
}

func b(input *bufio.Scanner) {
	s := MakeSolver(input)
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
