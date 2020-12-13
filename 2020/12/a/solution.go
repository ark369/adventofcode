package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type boat struct {
	dir int // N:0, E:1, S:2, W:3
	x, y int
}

func newBoat() *boat {
	return &boat{1, 0, 0}
}

func (b *boat) travel(instr string) {
	dir := instr[0]
	val, err := strconv.Atoi(instr[1:])
	if err != nil {
		panic(err)
	}
	switch dir {
	case 'N':
		b.y += val
	case 'S':
		b.y -= val
	case 'E':
		b.x += val
	case 'W':
		b.x -= val
	case 'F':
		switch b.dir {
		case 0:
			b.y += val
		case 2:
			b.y -= val
		case 1:
			b.x += val
		case 3:
			b.x -= val
		}
	case 'R':
		num := val / 90
		b.dir = (b.dir + num) % 4
	case 'L':
		num := val / 90
		b.dir = (b.dir - num)
		if b.dir < 0 {
			b.dir += 4
		}
	}
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	b := newBoat()
	for _, l := range input {
		b.travel(l)
	//fmt.Printf("%s > %d, %d\n", l, b.x, b.y)
	}
	fmt.Printf("%d + %d = %v\n", b.x, b.y, math.Abs(float64(b.x)) + math.Abs(float64(b.y)))
}


func ReadFakeInput() []string {
	input := `F10
N3
F7
R90
F11`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
