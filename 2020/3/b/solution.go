package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	var r1d1, r3d1, r5d1, r7d1, r1d2, tr1d1, tr3d1, tr5d1, tr7d1, tr1d2 int
	even := true
	tree := "#"[0]
	for _, l := range(input) {
		if l[r1d1 % len(l)] == tree {
			tr1d1 += 1
		}
		r1d1 += 1

		if l[r3d1 % len(l)] == tree {
			tr3d1 += 1
		}
		r3d1 += 3

		if l[r5d1 % len(l)] == tree {
			tr5d1 += 1
		}
		r5d1 += 5

		if l[r7d1 % len(l)] == tree {
			tr7d1 += 1
		}
		r7d1 += 7

		if even && l[r1d2 % len(l)] == tree {
			tr1d2 += 1
		}
		if even {
			r1d2 += 1
		}
		even = !even
	}
	fmt.Printf("%d x %d x %d x %d x %d = %d", tr1d1, tr3d1, tr5d1, tr7d1, tr1d2, tr1d1 * tr3d1 * tr5d1 * tr7d1 * tr1d2)
}

func ReadFakeInput() []string {
	input := 
`..##.......
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
