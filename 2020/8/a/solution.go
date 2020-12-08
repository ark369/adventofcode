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
	input := ReadFakeInput()
	ins := make([]instr, len(input))
	seen := map[int]bool{}
	for i, l := range(input) {
		ins[i] = makeInstr(l)
	}
	acc := 0
	curr := 0
	for {
		if _, ok := seen[curr]; ok {
			break
		}
		seen[curr] = true
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
	fmt.Printf("%d", acc)
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
