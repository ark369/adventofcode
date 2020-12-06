package main

import (
	"fmt"
	"strings"
)

var (
	bug = '#'
)

type coord struct {
	dim, i, j int
}

func main() {
	input := ReadInput()
	tiles := map[int][][]rune{0: makeTiles()}
	for i, l := range(input) {
		for j, r := range(l) {
			tiles[0][i][j] = r
		}
	}
	lowerDim := -1
	upperDim := 1
	for i := 0; i < 200; i++ {
		next := map[int][][]rune{}
		for dim := lowerDim; dim <= upperDim; dim++ {
			if _, ok := tiles[dim]; !ok {
				tiles[dim] = makeTiles()
			}
			next[dim] = makeTiles()
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if i == 2 && j == 2 {
						continue
					}
					adj := adj(coord{dim, i, j})
					numBugs := 0
					for _, c := range(adj) {
						if isBug(tiles, c) {
							numBugs += 1
						}
					}
					if tiles[dim][i][j] == bug {
						if numBugs == 1 {
							next[dim][i][j] = bug
						}
					} else {
						if numBugs == 1 || numBugs == 2{
							next[dim][i][j] = bug
						}
					}
				}
			}
		}
		tiles = next
		if numBugsInDim(tiles[lowerDim]) > 0 {
			lowerDim -= 1
		}
		if numBugsInDim(tiles[upperDim]) > 0 {
			upperDim += 1
		}
	}
	
	numBugs := 0
	for _, t := range(tiles) {
		numBugs += numBugsInDim(t)
	}
	
	fmt.Println(numBugs)
	
	for i := lowerDim; i <= upperDim; i++ {
		fmt.Printf("Depth %d\n", i)
		for _, j := range(tiles[i]) {
			str := []rune{}
			for _, r := range(j) {
				if r == bug {
					str = append(str, bug)
				} else {
					str = append(str, '.')
				}
			}
			fmt.Println(string(str))
		}
	}
}

func numBugsInDim(t [][]rune) int {
	numBugs := 0
	for _, j := range(t) {
		for _, r := range(j) {
			if r == bug {
				numBugs += 1
			}
		}
	}
	return numBugs
}

func makeTiles() [][]rune {
	tiles := make([][]rune, 5)
	for i := 0; i < 5; i++ {
		tiles[i] = make([]rune, 5)
	}
	return tiles
}

/*
     |     |         |     |     
  1  |  2  |    3    |  4  |  5  
     |     |         |     |     
-----+-----+---------+-----+-----
     |     |         |     |     
  6  |  7  |    8    |  9  |  10 
     |     |         |     |     
-----+-----+---------+-----+-----
     |     |A|B|C|D|E|     |     
     |     |-+-+-+-+-|     |     
     |     |F|G|H|I|J|     |     
     |     |-+-+-+-+-|     |     
 11  | 12  |K|L|?|N|O|  14 |  15 
     |     |-+-+-+-+-|     |     
     |     |P|Q|R|S|T|     |     
     |     |-+-+-+-+-|     |     
     |     |U|V|W|X|Y|     |     
-----+-----+---------+-----+-----
     |     |         |     |     
 16  | 17  |    18   |  19 |  20 
     |     |         |     |     
-----+-----+---------+-----+-----
     |     |         |     |     
 21  | 22  |    23   |  24 |  25 
     |     |         |     |     
*/

func adj(c coord)[]coord {
	ret := []coord{}
	
	// above
	if c.i == 0 { // A, B, C, D, E
		ret = append(ret, coord{c.dim - 1, 1, 2})
	} else if c.i == 3 && c.j == 2 { // 18
		ret = append(ret, 
			coord{c.dim + 1, 4, 0},
			coord{c.dim + 1, 4, 1},
			coord{c.dim + 1, 4, 2},
			coord{c.dim + 1, 4, 3},
			coord{c.dim + 1, 4, 4},
		)
	} else {
		ret = append(ret, coord{c.dim, c.i - 1, c.j})
	}
	
	// below
	if c.i == 4 { // U, V, W, X, Y
		ret = append(ret, coord{c.dim - 1, 3, 2})
	} else if c.i == 1 && c.j == 2 { // 8
		ret = append(ret, 
			coord{c.dim + 1, 0, 0},
			coord{c.dim + 1, 0, 1},
			coord{c.dim + 1, 0, 2},
			coord{c.dim + 1, 0, 3},
			coord{c.dim + 1, 0, 4},
		)
	} else {
		ret = append(ret, coord{c.dim, c.i + 1, c.j})
	}
	
	// left
	if c.j == 0 { // E, J, O, T, Y
		ret = append(ret, coord{c.dim - 1, 2, 1})
	} else if c.i == 2 && c.j == 3 { // 14
		ret = append(ret, 
			coord{c.dim + 1, 0, 4},
			coord{c.dim + 1, 1, 4},
			coord{c.dim + 1, 2, 4},
			coord{c.dim + 1, 3, 4},
			coord{c.dim + 1, 4, 4},
		)
	} else {
		ret = append(ret, coord{c.dim, c.i, c.j - 1})
	}
	
	// right
	if c.j == 4 { // A, F, K, P, U
		ret = append(ret, coord{c.dim - 1, 2, 3})
	} else if c.i == 2 && c.j == 1 { // 12
		ret = append(ret, 
			coord{c.dim + 1, 0, 0},
			coord{c.dim + 1, 1, 0},
			coord{c.dim + 1, 2, 0},
			coord{c.dim + 1, 3, 0},
			coord{c.dim + 1, 4, 0},
		)
	} else {
		ret = append(ret, coord{c.dim, c.i, c.j + 1})
	}	
	
	return ret
}

func isBug(tiles map[int][][]rune, c coord) bool {
	t, ok := tiles[c.dim]
	if !ok {
		t = makeTiles()
		tiles[c.dim] = t
	}
	return t[c.i][c.j] == bug
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
