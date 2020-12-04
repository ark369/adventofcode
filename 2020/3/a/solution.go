package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	numTrees := 0
	x := 0
	for _, l := range(input) {
		if l[x % len(l)] == "#"[0] {
			numTrees += 1
		}
		x += 3
	}
	fmt.Printf("Num trees: %d", numTrees)
}


func log(l string) {
	fmt.Println(l)
}

func ReadFakeInput() []string {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
