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
	for i := maxElf; i > 1; i-- {
		for j := 1; j < maxElf; j++ {
			if j * i > maxElf {
				break
			}
			houses[j * i] += i * 10
		}
	}
	for i:= 1; i <= maxElf; i++ {
		houses[i] += 10
		if houses[i] >= target {
			houseNum = i
			maxPresents = houses[i]
			break
		}
	}
	fmt.Printf("Reached %d at houseNum %d", maxPresents, houseNum)
}
