package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	houses := make(map[string]bool)
	var x, y int
	houses[fmt.Sprintf("%d,%d", x, y)] = true
	for _, c := range(input[0]) {
		if c == '^' {
			y += 1
		}
		if c == '<' {
			x -= 1
		}
		if c == '>' {
			x += 1
		}
		if c == 'v' {
			y -= 1
		}
		houses[fmt.Sprintf("%d,%d", x, y)] = true
		//fmt.Printf("%v\n", houses)
	}
	fmt.Printf("Total %d", len(houses))
}

func ReadFakeInput() []string {
	input := `^v^v^v^v^v`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
