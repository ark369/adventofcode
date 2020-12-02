package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pwCheck struct {
	min, max int
	check rune
	pw string
}

func main() {
	input := ReadInput()
	pws := []pwCheck{}
	for _, l := range input {
		parts := strings.Split(l, " ")
		minMax := strings.Split(parts[0], "-")
		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			panic(err)
		}
		check := parts[1]
		
		newPW := pwCheck{
			min: min,
			max: max,
			check: []rune(check)[0],
			pw: parts[2],
		}
		pws = append(pws, newPW)
	}
	
	numValid := 0
	for _, pw := range(pws) {
		valid := true
		numSeen := 0
		for _, c := range([]rune(pw.pw)) {
			if c == pw.check {
				numSeen += 1
			}
			if numSeen > pw.max {
				valid = false
				break
			}
		}
		if valid && numSeen >= pw.min {
			numValid += 1
		}
	}
	fmt.Printf("Num valid: %d", numValid)
}

func ReadFakeInput() []string {
	input := []string{"1-3 a: abcde",
"1-3 b: cdefg",
"2-9 c: ccccccccc"}
	return input
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
