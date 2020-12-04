package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	var floor int
	for _, c := range(input[0]) {
		if c == rune("("[0]) {
			floor += 1
		} else {
			floor -= 1
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
