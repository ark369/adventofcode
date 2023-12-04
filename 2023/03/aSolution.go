package main

import (
	"fmt"
	"strings"
)

func isSymbol(c rune) bool {
	if c >= '0' && c <= '9' {
		return false
	}
	if c == '.' {
		return false
	}
	return true
}

func AdjacentToParts(l string, start, end int) bool {
	i := start - 1
	if i < 0 {
		i = 0
	}
	j := end + 1
	if j > len(l)-1 {
		j = len(l) - 1
	}
	for k := i; k <= j; k++ {
		if isSymbol(rune(l[k])) {
			return true
		}
	}
	return false
}

func AdjacentToPartsCurrLine(l string, start, end int) bool {
	if start > 0 && isSymbol(rune(l[start-1])) {
		return true
	}
	if end < len(l)-1 && isSymbol(rune(l[end+1])) {
		return true
	}
	return false
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	numLines := len(input)
	sum := 0
	for i, l := range input {
		currIsNumber := false
		currNumber := 0
		currNumberLen := 0
		for j, c := range l {
			if c >= '0' && c <= '9' {
				n := int(c - '0')
				currNumber = currNumber*10 + n
				currNumberLen++
				currIsNumber = true
				if j == len(l)-1 {
					include := false
					//fmt.Printf("currNumber: %d\n", currNumber)
					if i != 0 {
						if AdjacentToParts(input[i-1], j-currNumberLen+1, j) {
							include = true
						}
					}
					if i != numLines-1 {
						if AdjacentToParts(input[i+1], j-currNumberLen+1, j) {
							include = true
						}
					}
					if AdjacentToPartsCurrLine(l, j-currNumberLen+1, j) {
						include = true
					}
					if include {
						//fmt.Printf("included\n")
						sum += currNumber
					}
				}
			} else if currIsNumber && currNumber > 0 {
				include := false
				//fmt.Printf("currNumber: %d\n", currNumber)
				if i != 0 {
					if AdjacentToParts(input[i-1], j-currNumberLen, j-1) {
						include = true
					}
				}
				if i != numLines-1 {
					if AdjacentToParts(input[i+1], j-currNumberLen, j-1) {
						include = true
					}
				}
				if AdjacentToPartsCurrLine(l, j-currNumberLen, j-1) {
					include = true
				}
				if include {
					//fmt.Printf("included\n")
					sum += currNumber
				}
				currNumber = 0
				currNumberLen = 0
			}
		}
	}
	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
