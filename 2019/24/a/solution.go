package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	tiles := make([][]rune, 7)
	seen := map[int]bool{}
	for i := 0; i < 7; i++ {
		tiles[i] = make([]rune, 7)
	}
	for i, l := range(input) {
		for j, r := range(l) {
			tiles[i+1][j+1] = r
		}
	}
	for {
		bio := calcBio(tiles)
		if seen[bio] {
			fmt.Println(bio)
			return
		}
		seen[bio] = true
		tiles = next(tiles)
	}
}

func next(tiles [][]rune) [][]rune {
	n := make([][]rune, 7)
	for i := 0; i < 7; i++ {
		n[i] = make([]rune, 7)
	}
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {
			bugs := numBugs(tiles[i-1][j], tiles[i+1][j], tiles[i][j-1], tiles[i][j+1])
			if tiles[i][j] == '#' {
				if bugs == 1 {
					n[i][j] = '#'
				}
			} else {
				if bugs == 1 || bugs == 2 {
					n[i][j] = '#'
				}
			}
		}
	}
	return n
}

func numBugs(u, d, l, r rune) int {
	tot := 0
	if u == '#' {
		tot += 1
	}
	if d == '#' {
		tot += 1
	}
	if l == '#' {
		tot += 1
	}
	if r == '#' {
		tot += 1
	}
	return tot
}

func calcBio(tiles [][]rune) int {
	bio := 0
	val := 1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if tiles[i+1][j+1] == '#' {
				bio += val
			}
			val *= 2
		}
	}
	return bio
}

func ReadFakeInput() []string {
	input := `....#
#..#.
#..##
..#..
#....`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := `#####
...##
#..#.
#....
#...#`
	return strings.Split(input, "\n")
}
