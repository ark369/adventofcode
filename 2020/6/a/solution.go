package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	var total int
	currGroup := make(map[rune]bool)
	for ind, l := range(input) {
		for _, q := range(l) {
			currGroup[q] = true
		}
		if len(l) == 0 || ind == len(input) - 1 {
			for range(currGroup) {
				total += 1
			}
			currGroup = make(map[rune]bool)
		}
	}
	fmt.Printf("Total: %d", total)
}

func ReadFakeInput() []string {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
