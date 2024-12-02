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

func IsSafe(levels []int) bool {
	increasing := true
	prev := levels[0]
	curr := levels[1]
	if curr == prev {
		return false
	} else if curr < prev {
		increasing = false
	}
	for i := 1; i < len(levels); i++ {
		curr = levels[i]
		if increasing {
			if curr < prev+1 || curr > prev+3 {
				return false
			} else {
				prev = curr
			}
		} else {
			if curr < prev-3 || curr > prev-1 {
				return false
			} else {
				prev = curr
			}
		}
	}
	return true
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	num := 0

	for n, l := range input {
		safe := false
		parts := strings.Split(l, " ")
		levels := []int{}
		for _, v := range parts {
			levels = append(levels, Atoi(v))
		}
		//fmt.Printf("n=%d, levels=%v\n", n, levels)
		if IsSafe(levels) {
			num++
			safe = true
			//fmt.Println("Safe 1")
		} else {
			if IsSafe(levels[1:]) {
				num++
				safe = true
				//fmt.Println("Safe 2")
			} else if IsSafe(levels[:len(levels)-1]) {
				num++
				safe = true
				//fmt.Println("Safe 3")
			} else {
				for i := 1; i < len(levels)-1; i++ {
					newLevels := []int{}
					newLevels = append(newLevels, levels[0:i]...)
					//fmt.Printf("newLevels i=%d, levels=%v\n", i, newLevels)
					newLevels = append(newLevels, levels[i+1:]...)
					//fmt.Printf("newLevels i=%d, levels=%v\n", i, newLevels)
					if IsSafe(newLevels) {
						num++
						safe = true
						//fmt.Printf("Safe 4:%d\n", i)
						break
					}
				}
			}
		}
		if !safe {
			fmt.Printf("UNSAFE n=%d, levels=%v\n", n, levels)
		}
	}

	fmt.Print(num)
}

func ReadFakeInput() []string {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
