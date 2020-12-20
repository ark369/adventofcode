package main

import (
	"fmt"
	"strings"
)

type dim struct {
	cubes [][]bool
	minX, minY, maxX, maxY, z, w int
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

func (d *dim) next(adjD []*dim) *dim {
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
			adj := d.getAdj(x, y, adjD)
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
		z: d.z,
		w: d.w,
	}
}

func (d *dim) getAdj(x, y int, adjD []*dim) int {
	adj := 0
	for _, ad := range adjD {
		if ad != nil {
			adj += ad.getAdj2D(x, y)
		}
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
	output := fmt.Sprintf("z=%d, w=%d, x:(%d->%d), y:(%d->%d)\n", d.z, d.w, d.minX, d.maxX, d.minY, d.maxY)
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

func makeDim(minX, maxX, minY, maxY, z, w int) *dim {
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
		z: z,
		w: w,
	}
}

type pocket struct {
	dims map[zw]*dim
	minZ, maxZ, minW, maxW, minX, minY, maxX, maxY int
}

func (p *pocket) next() *pocket {
	minZ := p.minZ
	maxZ := p.maxZ
	minW := p.minW
	maxW := p.maxW
	minX := p.minX
	maxX := p.maxX
	minY := p.minY
	maxY := p.maxY
	dims := map[zw]*dim{}
	for z := p.minZ - 1; z < p.maxZ + 1; z++ {
		for w := p.minW - 1; w < p.maxW + 1; w++ {
			d, ok := p.dims[zw{z:z, w:w}]
			if !ok {
				d = makeDim(p.minX, p.maxX, p.minY, p.maxY, z, w)
			}
			adjD := []*dim{
				p.dims[zw{z:z - 1, w:w}],
				p.dims[zw{z:z + 1, w:w}],
				p.dims[zw{z:z - 1, w:w - 1}],
				p.dims[zw{z:z, w:w - 1}],
				p.dims[zw{z:z + 1, w:w - 1}],
				p.dims[zw{z:z - 1, w:w + 1}],
				p.dims[zw{z:z, w:w + 1}],
				p.dims[zw{z:z + 1, w:w + 1}],
			}
			newDim := d.next(adjD)
			dims[zw{z:z, w:w}] = newDim
		}
	}
	return &pocket{
		dims: dims,
		minZ: minZ - 1,
		maxZ: maxZ + 1,
		minW: minW - 1,
		maxW: maxW + 1,
		minX: minX - 1,
		maxX: maxX + 1,
		minY: minY - 1,
		maxY: maxY + 1,
	}
}

func (p *pocket) count() int {
	sum := 0
	for z := p.minZ; z < p.maxZ; z++ {
		for w := p.minW; w < p.maxW; w++ {
			sum += p.dims[zw{z:z, w:w}].count()
		}
	}
	return sum
}

func (p *pocket) String() string {
	output := ""
	for w := p.minW; w < p.maxW; w++ {
		for z := p.minZ; z < p.maxZ; z++ {
			output += p.dims[zw{z:z, w:w}].String() + "\n"
		}
	}
	return output
}

type zw struct {
	z, w int
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	d := makeFirstDim(input)
	p := &pocket{
		dims: map[zw]*dim{zw{}: d},
		maxZ: 1,
		maxW: 1,
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
