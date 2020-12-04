package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	var floor int
	for ind, c := range(input[0]) {
		if c == rune("("[0]) {
			floor += 1
		} else {
			floor -= 1
		}
		if floor < 0 {
			fmt.Printf("Pos %d is lower than 0", ind + 1)
			return
		}
	}
	fmt.Printf("Floor %d", floor)
}

func ReadFakeInput() []string {
	input := ``
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
