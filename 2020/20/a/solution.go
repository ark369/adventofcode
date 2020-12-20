package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
ABCDEF
G    H
I    J
K    L
M    N
OPQRST

top (0):    ABCDEF
bottom (1): TSRQPO
left (2):   OMKIGA
right (3):  FHJLNT

topR (4):    FEDCBA
bottomR (5): OPQRST
leftR (6):   AGIKMO
rightR (7):  TNLJHF
*/
type tile struct {
	num int
	top, bottom, left, right string
	sides map[string]int
	connections []*tile
}

func makeTile(lines []string) *tile {
	numStr := lines[0][5:len(lines[0]) - 1]
	num, _ := strconv.Atoi(numStr)
	lines = lines[1:]
	var top, bottom, left, right, topR, bottomR, leftR, rightR string
	for i, l := range lines {
		for j, r := range l {
			s := string(r)
			if i == 0 {
				top += s
				topR = s + topR
			}
			if i == len(lines) - 1 {
				bottomR += s
				bottom = s + bottom
			}
			if j == 0 {
				leftR += s
				left = s + left
			}
			if j == len(l) - 1 {
				right += s
				rightR = s + rightR
			}
		}
	}
	sides := map[string]int{
		top: 0,
		bottom: 1,
		left: 2,
		right: 3,
		topR: 4,
		bottomR: 5,
		leftR: 6,
		rightR: 7,
	}
	return &tile{
		num: num,
		top: top,
		bottom: bottom,
		left: left,
		right: right,
		sides: sides,
		connections: []*tile{},
	}
}

func (t *tile) String() string {
	return fmt.Sprintf("num: %d, top: %s, bottom: %s, left: %s, right: %s", t.num, t.top, t.bottom, t.left, t.right)
}

type tiles struct {
	t []*tile
	sideMap map[string][]*tile
}

func (t *tiles) buildSideMap() {
	for _, tt := range t.t {
		for k, _ := range tt.sides {
			list, ok := t.sideMap[k]
			if !ok {
				list = []*tile{}
			}
			list = append(list, tt)
			if len(list) == 2 {
				list[0].connections = append(list[0].connections, list[1])
				list[1].connections = append(list[1].connections, list[0])
			}
			if len(list) > 2 {
				panic("more than 2 matches")
			}
			t.sideMap[k] = list
		}
	}
}

func (t *tiles) getCorners() []int {
	p := []int{}
	for _, tt := range t.t {
		if len(tt.connections) == 4 {
			p = append(p, tt.num)
		}
	}
	return p
}

func (t *tiles) String() string {
	output := ""
	for _, tt := range t.t {
		output += tt.String()
	}
	return output
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	t := tiles{}
	lines := []string{}
	for i, l := range input {
		if l == "" {
			t.t = append(t.t, makeTile(lines))
			lines = []string{}
			continue
		}
		lines = append(lines, l)
		if i == len(input) - 1 {
			t.t = append(t.t, makeTile(lines))
		}
	}
	
	t.sideMap = map[string][]*tile{}
	t.buildSideMap()
	
	p := 1
	for _, c := range t.getCorners() {
		p *= c
		fmt.Printf("C: %d\n", c)
	}
	fmt.Printf("product: %d\n", p)
	
	//fmt.Printf("t: %s\n", t)
	//fmt.Printf("%s\n", p)
}


func ReadFakeInput() []string {
	input := `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
