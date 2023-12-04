package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	maxR, maxG, maxB int
	r, g, b          int
}

func NewGame(maxR, maxG, maxB int) *Game {
	return &Game{maxR, maxG, maxB, 0, 0, 0}
}

func (g *Game) valid() bool {
	return g.r <= g.maxR && g.g <= g.maxG && g.b <= g.maxB
}

func (g *Game) parseColor(s string) {
	parts := strings.Split(s, " ")
	i, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	switch parts[1] {
	case "red":
		g.r = i
	case "green":
		g.g = i
	case "blue":
		g.b = i
	default:
		panic(fmt.Sprintf("Received unknown color: %s\n", parts[1]))
	}
}

func parseId(s string) int {
	idStr := strings.Split(s, " ")[1]
	i, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	sum := 0
	for _, l := range input {
		gameId := parseId(strings.Split(l, ":")[0])
		gamesStr := strings.Split(strings.Split(l, ":")[1], ";")
		//fmt.Printf("gameId: %d\n", gameId)
		valid := true
		for _, g := range gamesStr {
			game := NewGame(12, 13, 14)
			//fmt.Printf("game: %s\n", g)
			col := strings.Split(g, ",")
			for _, c := range col {
				game.parseColor(strings.TrimSpace(c))
				if !game.valid() {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			sum += gameId
		}
	}
	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
