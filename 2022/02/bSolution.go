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

func ScoreRound(s string) int {
	opp := strings.Split(s, " ")[0]
	requiredOutcome := strings.Split(s, " ")[1]
	me := ""
	switch requiredOutcome {
	case "X":
		switch opp {
		case "A":
			me = "Z"
		case "B":
			me = "X"
		case "C":
			me = "Y"
		}
	case "Y":
		switch opp {
		case "A":
			me = "X"
		case "B":
			me = "Y"
		case "C":
			me = "Z"
		}
	case "Z":
		switch opp {
		case "A":
			me = "Y"
		case "B":
			me = "Z"
		case "C":
			me = "X"
		}
	}
	var base, outcome int
	if me == "X" {
		base = 1
		if opp == "A" {
			outcome = 3
		} else if opp == "B" {
			outcome = 0
		} else if opp == "C" {
			outcome = 6
		}
	} else if me == "Y" {
		base = 2
		if opp == "A" {
			outcome = 6
		} else if opp == "B" {
			outcome = 3
		} else if opp == "C" {
			outcome = 0
		}
	} else if me == "Z" {
		base = 3
		if opp == "A" {
			outcome = 0
		} else if opp == "B" {
			outcome = 6
		} else if opp == "C" {
			outcome = 3
		}
	}
	return base + outcome
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	score := 0
	for _, l := range input {
		score += ScoreRound(l)
	}

	fmt.Printf("score: %d\n", score)
}

func ReadFakeInput() []string {
	input := `A Y
B X
C Z`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
