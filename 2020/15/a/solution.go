package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasBeenSeen(n, ind int, nums map[int]int) (bool, int) {
	seen := false
	dist := 0
	if m, ok := nums[n]; ok {
		seen = true
		dist = ind - m
	}
	nums[n] = ind
	return seen, dist
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	split := strings.Split(input[0], ",")
	nums := make(map[int]int)
	seen := false
	dist := 0
	for i, sp := range split {
		n, _ := strconv.Atoi(sp)
		seen, dist = hasBeenSeen(n, i, nums)
	}
	n := 0
	for i := len(nums); i < 2020; i++ {
		n = 0
		if seen {
			n = dist
		}
		seen, dist = hasBeenSeen(n, i, nums)
		//fmt.Printf("seen: %v, dist: %d, n: %d, i: %d, nums: %v\n", seen, dist, n, i, nums)
	}
	fmt.Printf("n: %d\n", n)
}


func ReadFakeInput() []string {
	input := `0,3,6`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := `17,1,3,16,19,0`
	return strings.Split(input, "\n")
}
