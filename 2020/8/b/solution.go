package main

import (
	"fmt"
	"strconv"
	"strings"
)

type instr struct {
	op int // 0 nop, 1 acc, 2 jmp
	val int
}

func main() {
	input := ReadInput()
	ins := make([]instr, len(input))
	for i, l := range(input) {
		ins[i] = makeInstr(l)
	}
	_, path, _ := isLoop(ins)
	for _, i := range(path) {
		in := ins[i]
		if in.op == 1 {
			continue
		}
		c := make([]instr, len(ins))
		copy(c, ins)
		if in.op == 0 {
			c[i].op = 2
		}
		if in.op == 2 {
			c[i].op = 0
		}
		if loop, _, acc := isLoop(c); !loop {
			fmt.Printf("%d", acc)
			break
		}
	}
}

func isLoop(ins []instr) (bool, []int, int) {
	loop := false
	path := []int{}
	seen := map[int]bool{}
	curr := 0
	acc := 0
	for {
		if curr >= len(ins) {
			break
		}
		if _, ok := seen[curr]; ok {
			loop = true
			break
		}
		seen[curr] = true
		path = append(path, curr)
		in := ins[curr]
		if in.op == 1 {
			acc += in.val
		}
		if in.op == 2 {
			curr += in.val
		} else {
			curr += 1
		}
	}
	return loop, path, acc
}

func makeInstr(l string) instr {
	op := l[:3]
	val, err := strconv.Atoi(l[4:])
	if err != nil {
		panic(err)
	}
	opCode := 0
	if op == "acc" {
		opCode = 1
	}
	if op == "jmp" {
		opCode = 2
	}
	return instr{opCode, val}
}

func ReadFakeInput() []string {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
