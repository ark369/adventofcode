package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput()
	target := 2020
	diffs := make(map[int]bool)
	for _, l := range input {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		if _, ok := diffs[i]; ok {
			fmt.Printf("Found %d x %d = %d", i, target-i, i*(target-i))
			return
		}
		diffs[target-i] = true
	}
	fmt.Printf("NOT FOUND")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
