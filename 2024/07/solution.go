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

type OpTree struct {
	target   int
	operands []int
}

func MakeOpTree(target int, ops []string) *OpTree {
	o := &OpTree{}
	o.target = target
	o.operands = []int{}
	for _, op_string := range ops {
		op := Atoi(op_string)
		o.operands = append(o.operands, op)
	}
	return o
}

func DFSEvalMultPlus(curr int, target int, remaining []int) bool {
	next := remaining[0]
	remaining = remaining[1:]

	if len(remaining) == 0 {
		if curr*next == target {
			return true
		}
		if curr+next == target {
			return true
		}
		return false
	}

	if curr*next <= target {
		if DFSEvalMultPlus(curr*next, target, remaining) {
			return true
		}
	}
	if curr+next <= target {
		if DFSEvalMultPlus(curr+next, target, remaining) {
			return true
		}
	}
	return false
}

func (o *OpTree) CanReachTargetMultPlus() bool {
	curr := o.operands[0]
	remaining := o.operands[1:]

	return DFSEvalMultPlus(curr, o.target, remaining)
}

func ParseA(l string) int {
	parts := strings.Split(l, ": ")
	target := Atoi(parts[0])
	operands_parts := strings.Split(parts[1], " ")
	tree := MakeOpTree(target, operands_parts)
	if tree.CanReachTargetMultPlus() {
		return target
	}
	return 0
}

func a(input *bufio.Scanner) {
	sum := 0
	for input.Scan() {
		l := input.Text()
		sum += ParseA(l)
	}
	fmt.Println(sum)
}

func Concat(a, b int) int {
	return Atoi(Itoa(a) + Itoa(b))
}

func DFSEvalMultPlusConcat(curr int, target int, remaining []int) bool {
	next := remaining[0]
	remaining = remaining[1:]

	if len(remaining) == 0 {
		if curr*next == target {
			return true
		}
		if curr+next == target {
			return true
		}
		if Concat(curr, next) == target {
			return true
		}
		return false
	}

	if curr*next <= target {
		if DFSEvalMultPlusConcat(curr*next, target, remaining) {
			return true
		}
	}
	if curr+next <= target {
		if DFSEvalMultPlusConcat(curr+next, target, remaining) {
			return true
		}
	}
	if Concat(curr, next) <= target {
		if DFSEvalMultPlusConcat(Concat(curr, next), target, remaining) {
			return true
		}
	}
	return false
}

func (o *OpTree) CanReachTargetMultPlusConcat() bool {
	curr := o.operands[0]
	remaining := o.operands[1:]

	return DFSEvalMultPlusConcat(curr, o.target, remaining)
}

func ParseB(l string) int {
	parts := strings.Split(l, ": ")
	target := Atoi(parts[0])
	operands_parts := strings.Split(parts[1], " ")
	tree := MakeOpTree(target, operands_parts)
	if tree.CanReachTargetMultPlusConcat() {
		return target
	}
	return 0
}

func b(input *bufio.Scanner) {
	sum := 0
	for input.Scan() {
		l := input.Text()
		sum += ParseB(l)
	}
	fmt.Println(sum)
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
