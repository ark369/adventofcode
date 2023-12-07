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

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	sum := 0
	for _, l := range input {
		sum += GetPriority(l)
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
