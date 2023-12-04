package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseNums(s string) []int {
	ret := []int{}
	nums := strings.Split(s, " ")
	for _, ns := range nums {
		if ns == "" {
			continue
		}
		n, _ := strconv.Atoi(ns)
		ret = append(ret, n)
	}
	return ret
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()
	sum := 0
	cards := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		cards[i] = 1
	}
	for i, l := range input {
		winners := make(map[int]bool)
		num := 0
		for _, n := range parseNums(strings.Split(strings.Split(l, ":")[1], "|")[0]) {
			winners[n] = true
		}
		for _, n := range parseNums(strings.Split(strings.Split(l, ":")[1], "|")[1]) {
			if winners[n] {
				num++
			}
		}
		for j := 0; j < num; j++ {
			cards[i+j+1] += cards[i]
		}
	}
	for i := 0; i < len(input); i++ {
		sum += cards[i]
	}
	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
