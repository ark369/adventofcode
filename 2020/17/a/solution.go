package main

import (
	"fmt"
	"strings"
)

type dim struct {
	cubes [][]bool
	minX, minY, maxX, maxY, num int
}

func (d *dim) empty() bool {
	for _, r := range d.cubes {
		for _, c := range r {
			if c {
				return false
			}
		}
	}
	return true
}

func (d *dim) count() int {
	sum := 0
	for _, r := range d.cubes {
		for _, c := range r {
			if c {
				sum += 1
			}
		}
	}
	return sum
}

func (d *dim) next(below, above *dim) *dim {
	d.minX -= 1
	d.maxX += 1
	d.minY -= 1
	d.maxY += 1
	lx := d.maxX - d.minX
	ly := d.maxY - d.minY
	cubes := make([][]bool, lx)
	for i := 0; i < lx; i++ {
		cubes[i] = make([]bool, ly)
	}
	for x := -1; x < lx - 1; x++ {
		for y := -1; y < ly - 1; y++ {
			adj := d.getAdj(x, y, below, above)
			if x >= 0 && x < len(d.cubes) && y >= 0 && y < len(d.cubes[0]) && d.cubes[x][y] {
				if adj == 2 || adj == 3 {
					cubes[x + 1][y + 1] = true
				}
			} else if adj == 3 {
				cubes[x + 1][y + 1] = true
			}
		}
	}
	return &dim{
		cubes: cubes,
		minX: d.minX,
		maxX: d.maxX,
		minY: d.minY,
		maxY: d.maxY,
		num: d.num,
	}
}

func (d *dim) getAdj(x, y int, below, above *dim) int {
	adj := 0
	if below != nil {
		adj += below.getAdj2D(x, y)
	}
	if above != nil {
		adj += above.getAdj2D(x, y)
	}
	adj += d.getAdj2D(x, y)
	if x >= 0 && x < len(d.cubes) && y >= 0 && y < len(d.cubes[0]) && d.cubes[x][y] {
		adj -= 1
	}
	return adj
}

func (d *dim) getAdj2D(x, y int) int {
	adj := 0
	for i := x - 1; i <= x + 1; i++ {
		for j := y - 1; j <= y + 1; j++ {
			if i < 0 || i >= len(d.cubes) {
				continue
			}
			if j < 0 || j >= len(d.cubes[i]) {
				continue
			}
			if d.cubes[i][j] {
				adj += 1
			}
		}
	}
	return adj
}

func (d *dim) String() string {
	output := fmt.Sprintf("z=%d, x:(%d->%d), y:(%d->%d)\n", d.num, d.minX, d.maxX, d.minY, d.maxY)
	for _, r := range d.cubes {
		for _, b := range r {
			if b {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

func makeFirstDim(input []string) *dim {
	cubes := [][]bool{}
	for _, l := range input {
		row := []bool{}
		for _, r := range l {
			if r == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		cubes = append(cubes, row)
	}
	return &dim{
		cubes: cubes,
		maxX: len(cubes),
		maxY: len(cubes[0]),
	}
}

func makeDim(minX, maxX, minY, maxY, num int) *dim {
	x := maxX - minX
	cubes := make([][]bool, x)
	for i := 0; i < x; i++ {
		cubes[i] = make([]bool, maxY - minY)
	}
	return &dim{
		cubes: cubes,
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
		num: num,
	}
}

type pocket struct {
	dims map[int]*dim
	minD, maxD, minX, minY, maxX, maxY int
}

func (p *pocket) next() *pocket {
	minD := p.minD
	maxD := p.maxD
	minX := p.minX
	maxX := p.maxX
	minY := p.minY
	maxY := p.maxY
	dims := map[int]*dim{}
	for d := p.minD - 1; d < p.maxD + 1; d++ {
		dim, ok := p.dims[d]
		if !ok {
			dim = makeDim(p.minX, p.maxX, p.minY, p.maxY, d)
		}
		newDim := dim.next(p.dims[d - 1], p.dims[d + 1])
		dims[d] = newDim
	}
	if dims[minD - 1].empty() {
		delete(dims, minD - 1)
	} else {
		minD -= 1
	}
	if dims[maxD].empty() {
		delete(dims, maxD)
	} else {
		maxD += 1
	}
	return &pocket{
		dims: dims,
		minD: minD,
		maxD: maxD,
		minX: minX - 1,
		maxX: maxX + 1,
		minY: minY - 1,
		maxY: maxY + 1,
	}
}

func (p *pocket) count() int {
	sum := 0
	for d := p.minD; d < p.maxD; d++ {
		sum += p.dims[d].count()
	}
	return sum
}

func (p *pocket) String() string {
	output := ""
	for i := p.minD; i < p.maxD; i++ {
		output += p.dims[i].String() + "\n"
	}
	return output
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	d := makeFirstDim(input)
	p := &pocket{
		dims: map[int]*dim{0: d},
		maxD: 1,
		maxX: d.maxX,
		maxY: d.maxY,
	}
	for i := 0; i < 6; i++ {
		p = p.next()
	}
	fmt.Printf("Active: %d\n", p.count())
	//fmt.Printf("%s\n", p)
}


func ReadFakeInput() []string {
	input := `.#.
..#
###`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := `#.#.#.##
.####..#
#####.#.
#####..#
#....###
###...##
...#.#.#
#.##..##`
	return strings.Split(input, "\n")
}
