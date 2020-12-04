package main

import (
	"fmt"
)

func main() {
	//target := 121
	target := 34000000
	maxElf := target / 10
	var houseNum, maxPresents int
	houses := make([]int, maxElf + 1)
	for i := maxElf; i > 0; i-- {
		for j := 1; j < maxElf; j++ {
			if j * i > maxElf || j > 50 {
				break
			}
			houses[j * i] += i * 11
		}
	}
	for ind, h := range(houses) {
		if h >= target {
			houseNum = ind
			maxPresents = h
			break
		}
	}
	fmt.Printf("Reached %d at houseNum %d", maxPresents, houseNum)
}
