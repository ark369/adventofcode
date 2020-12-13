package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type boat struct {
	wx, wy, x, y int
}

func newBoat() *boat {
	return &boat{10, 1, 0, 0}
}

func (b *boat) travel(instr string) {
	dir := instr[0]
	val, err := strconv.Atoi(instr[1:])
	if err != nil {
		panic(err)
	}
	switch dir {
	case 'N':
		b.wy += val
	case 'S':
		b.wy -= val
	case 'E':
		b.wx += val
	case 'W':
		b.wx -= val
	case 'F':
		b.x += val * b.wx
		b.y += val * b.wy
	case 'R':
		num := val / 90
		if num == 1 {
			prevY := b.wy
			b.wy = b.wx * -1
			b.wx = prevY
		}
		if num == 2 {
			b.wx *= -1
			b.wy *= -1
		}
		if num == 3 {
			prevX := b.wx
			b.wx = b.wy * -1
			b.wy = prevX
		}
	case 'L':
		num := val / 90
		if num == 3 {
			prevY := b.wy
			b.wy = b.wx * -1
			b.wx = prevY
		}
		if num == 2 {
			b.wx *= -1
			b.wy *= -1
		}
		if num == 1 {
			prevX := b.wx
			b.wx = b.wy * -1
			b.wy = prevX
		}
	}
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	b := newBoat()
	for _, l := range input {
		b.travel(l)
	//fmt.Printf("%s > %d, %d, %d, %d\n", l, b.x, b.y, b.wx, b.wy)
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
