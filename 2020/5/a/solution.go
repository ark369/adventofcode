package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	var highest int
	rows := make(map[string]int)
	seats := make(map[string]int)
	for r := 0; r < 128; r++ {
		rows[makeBinaryStr(r, 7, 'F', 'B')] = r
	}
	for s := 0; s < 8; s++ {
		seats[makeBinaryStr(s, 3, 'L', 'R')] = s
		fmt.Println(seats)
	}
	fmt.Printf("rows: %v\nseats: %v\n", rows, seats)
	for _, l := range(input) {
		row := rows[l[:7]]
		seat := seats[l[7:]]
		id := row * 8 + seat
		if id > highest {
			highest = id
		}
	}
	fmt.Printf("highest %d", highest)
}

func makeBinaryStr(b, len int, zero, one byte) string {
	ret := make([]byte, len)
	for i := len-1; i >= 0; i-- {
		if b % 2 == 1 {
			ret[i] = one
		} else {
			ret[i] = zero
		}
		b = b/2
	}
	return string(ret) 
}

func ReadFakeInput() []string {
	input := `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
