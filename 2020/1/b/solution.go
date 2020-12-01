package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput()
	ints := make([]int, len(input))
	target := 2020
	targets := make(map[int]int)
	for ind, l := range input {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ints[ind] = i
	}
	sort.Ints(ints)
	for iInd, i := range ints[:len(ints)] {
		if j, ok := targets[i]; ok {
			k := target - i - j
			fmt.Printf("Found %d x %d x %d = %d", i, j, k, i*j*k)
			return
		}
		if iInd < len(ints) && i+ints[iInd+1] > target-1 {
			continue
		}
		for _, j := range ints[iInd+1:] {
			k := target - i - j
			if k < j-1 {
				break
			}
			targets[k] = i
		}
	}
	fmt.Printf("NOT FOUND")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
