package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func ToPartition(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return -1
}

func GetPriority(s string) int {
	b := [53]int{}
	n := len(s) / 2
	for i, c := range s {
		//fmt.Printf("c: %v, a: %v, A: %v, index: %d\n", c, 'a', 'A', int(c-'a'))
		if i < n {
			b[ToPartition(c)]++
		} else {
			if b[ToPartition(c)] > 0 {
				return ToPartition(c)
			}
		}
	}
	return -1
}

func GetBadgePriority(bag1, bag2, bag3 string) int {
	b := [53]int{}
	for _, c := range bag1 {
		b[ToPartition(c)] = 1
	}
	for _, c := range bag2 {
		if b[ToPartition(c)] == 1 {
			b[ToPartition(c)] = 2
		}
	}
	for _, c := range bag3 {
		if b[ToPartition(c)] == 2 {
			return ToPartition(c)
		}
	}
	return -1
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	sum := 0
	numGroups := len(input) / 3
	for i := 0; i < numGroups; i++ {
		j := i * 3
		sum += GetBadgePriority(input[j], input[j+1], input[j+2])
	}

	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
