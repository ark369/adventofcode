package main

import (
	"fmt"
	"strings"
)

type ferry struct {
	seats [][]int
	rows, cols int
}

func newFerry(input []string) *ferry {
	cols := len(input[0])
	rows := len(input)
	seats := make([][]int, rows + 2)
	for i := range seats {
		seats[i] = make([]int, cols + 2)
	}
	for i, in := range input {
		for j, val := range in {
			if val == 'L' {
				seats[i+1][j+1] = 1
			}
		}
	}
	return &ferry{seats, rows, cols}
}

func (f *ferry) str() string {
	str := []rune{}
	for i := 1; i < f.rows + 1; i++ {
		for j := 1; j < f.cols + 1; j++ {
			if f.seats[i][j] == 0 {
				str = append(str, '.')
			}
			if f.seats[i][j] == 1 {
				str = append(str, 'L')
			}
			if f.seats[i][j] == 2 {
				str = append(str, '#')
			}
		}
		str = append(str, '\n')
	}
	return string(str)
}

func (f *ferry) numOcc() int {
	num := 0
	for i := 1; i < f.rows + 1; i++ {
		for j := 1; j < f.cols + 1; j++ {
			if f.seats[i][j] == 2 {
				num += 1
			}
		}
	}
	return num
}

func (f *ferry) equals(o *ferry) bool {
	for i := 1; i < f.rows + 1; i++ {
		for j := 1; j < f.cols + 1; j++ {
			if f.seats[i][j] != o.seats[i][j] {
				return false
			}
		}
	}
	return true
}

func (f *ferry) adjOcc(i, j int) int {
	occ := 0
	for r := i - 1; r <= i + 1; r++ {
		for c := j - 1; c <= j + 1; c++ {
			if c == j && r == i {
				continue
			}
			if f.seats[r][c] == 2 {
				occ += 1
			}
		}
	}
	return occ
}

func (f *ferry) next() *ferry {
	n := make([][]int, f.rows + 2)
	for i := range f.seats {
		n[i] = make([]int, f.cols + 2)
	}
	for i := 1; i < f.rows + 1; i++ {
		for j := 1; j < f.cols + 1; j++ {
			if f.seats[i][j] == 0 {
				continue
			}
			occ := f.adjOcc(i, j)
			if f.seats[i][j] == 1 && occ == 0 {
				n[i][j] = 2
				continue
			}
			if f.seats[i][j] == 2 && occ >= 4 {
				n[i][j] = 1
				continue
			}
			n[i][j] = f.seats[i][j]
		}
	}
	return &ferry{n, f.rows, f.cols}
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	f := newFerry(input)
	//fmt.Println(f.str())
	prev := f
	f = f.next()
	for !f.equals(prev) {
		//fmt.Println(f.str())
		prev = f
		f = f.next()
	}
	fmt.Printf("%d\n", f.numOcc())
}


func ReadFakeInput() []string {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
