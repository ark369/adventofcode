package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	var total int
	numInGroup := 0
	currGroup := map[rune]int{}
	for ind, l := range(input) {
		if len(l) == 0 {
			for _, v := range(currGroup) {
				if v == numInGroup {
					total += 1
				}
			}
			numInGroup = 0
			currGroup = map[rune]int{}
			continue
		}
		for _, r := range(l) {
			currGroup[r] += 1
		}
		numInGroup += 1
		if ind == len(input) - 1 {
			for _, v := range(currGroup) {
				if v == numInGroup {
					total += 1
				}
			}
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
