package main

import (
	"fmt"
	"strings"
)

func AdjacentGears(l string, start, end int) []int {
	gears := []int{}
	i := start - 1
	if i < 0 {
		i = 0
	}
	j := end + 1
	if j > len(l)-1 {
		j = len(l) - 1
	}
	for k := i; k <= j; k++ {
		if rune(l[k]) == '*' {
			gears = append(gears, k)
		}
	}
	return gears
}

func AdjacentGearsCurrLine(l string, start, end int) []int {
	gears := []int{}
	if start > 0 && rune(l[start-1]) == '*' {
		gears = append(gears, start-1)
	}
	if end < len(l)-1 && rune(l[end+1]) == '*' {
		gears = append(gears, end+1)
	}
	return gears
}

func AddToGears(m map[string][]int, row int, gears []int, value int) {
	for _, i := range gears {
		gearKey := fmt.Sprintf("%d,%d", row, i)
		m[gearKey] = append(m[gearKey], value)
	}
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	numLines := len(input)
	sum := 0
	gears := make(map[string][]int)
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
					//fmt.Printf("currNumber: %d\n", currNumber)
					if i != 0 {
						AddToGears(gears, i-1, AdjacentGears(input[i-1], j-currNumberLen+1, j), currNumber)
					}
					if i != numLines-1 {
						AddToGears(gears, i+1, AdjacentGears(input[i+1], j-currNumberLen+1, j), currNumber)
					}
					AddToGears(gears, i, AdjacentGearsCurrLine(l, j-currNumberLen+1, j), currNumber)
				}
			} else if currIsNumber && currNumber > 0 {
				//fmt.Printf("currNumber: %d\n", currNumber)
				if i != 0 {
					AddToGears(gears, i-1, AdjacentGears(input[i-1], j-currNumberLen, j-1), currNumber)
				}
				if i != numLines-1 {
					AddToGears(gears, i+1, AdjacentGears(input[i+1], j-currNumberLen, j-1), currNumber)
				}
				AddToGears(gears, i, AdjacentGearsCurrLine(l, j-currNumberLen, j-1), currNumber)
				currNumber = 0
				currNumberLen = 0
			}
		}
	}
	for _, v := range gears {
		//fmt.Printf("k: %s\n", k)
		//for _, val := range v {
		//	fmt.Printf("val: %d\n", val)
		//}
		if len(v) == 2 {
			sum += v[0] * v[1]
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
