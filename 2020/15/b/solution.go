package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasBeenSeen(n, ind uint32, nums map[uint32]uint32, lowNums []uint32) (bool, uint32) {
	seen := false
	var dist uint32
	if n < uint32(len(lowNums)) {
		if lowNums[n] > 0 {
			seen = true
			dist = ind - lowNums[n]
		}
		lowNums[n] = ind
		return seen, dist
	}
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
	nums := make(map[uint32]uint32)
	lowNums := make([]uint32, 100000)
	seen := false
	var dist uint32
	for i, sp := range split {
		n, _ := strconv.Atoi(sp)
		seen, dist = hasBeenSeen(uint32(n), uint32(i+1), nums, lowNums)
	}
	var n uint32
	var i uint32
	for i = uint32(len(split))+1; i <= 30000000; i++ {
		n = 0
		if seen {
			n = dist
		}
		seen, dist = hasBeenSeen(n, i, nums, lowNums)
		//fmt.Printf("seen: %v, dist: %d, n: %d, i: %d, nums: %v, lowNums: %v\n", seen, dist, n, i, nums, lowNums)
	}
	fmt.Printf("n: %d\n", n)
	fmt.Printf("len)nums: %d\n", len(nums))
}


func ReadFakeInput() []string {
	input := `0,3,6`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := `17,1,3,16,19,0`
	return strings.Split(input, "\n")
}
