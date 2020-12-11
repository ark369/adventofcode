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
	seats := make([][]int, rows)
	for i := range seats {
		seats[i] = make([]int, cols)
	}
	for i, in := range input {
		for j, val := range in {
			if val == 'L' {
				seats[i][j] = 1
			}
		}
	}
	return &ferry{seats, rows, cols}
}

func (f *ferry) str() string {
	str := []rune{}
	for i := 0; i < f.rows; i++ {
		for j := 0; j < f.cols; j++ {
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
	for i := 0; i < f.rows; i++ {
		for j := 0; j < f.cols; j++ {
			if f.seats[i][j] == 2 {
				num += 1
			}
		}
	}
	return num
}

func (f *ferry) equals(o *ferry) bool {
	for i := 0; i < f.rows; i++ {
		for j := 0; j < f.cols; j++ {
			if f.seats[i][j] != o.seats[i][j] {
				return false
			}
		}
	}
	return true
}

func (f *ferry) adjOcc(i, j int) int {	
	occ := 0
	
	// A
	for n := 1; ; n++ {
		r := i - n
		c := j - n
		if r < 0 || c < 0 {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	//B
	for n := 1; ; n++ {
		r := i - n
		c := j
		if r < 0 {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	// C
	for n := 1; ; n++ {
		r := i - n
		c := j + n
		if r < 0 || c >= f.cols {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	// D
	for n := 1; ; n++ {
		r := i
		c := j - n
		if c < 0 {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	// E
	for n := 1; ; n++ {
		r := i
		c := j + n
		if c >= f.cols {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	// F
	for n := 1; ; n++ {
		r := i + n
		c := j - n
		if r >= f.rows || c < 0 {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	// G
	for n := 1; ; n++ {
		r := i + n
		c := j
		if r >= f.rows {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	// H
	for n := 1; ; n++ {
		r := i + n
		c := j + n
		if r >= f.rows || c >= f.cols {
			break
		}
		if f.seats[r][c] == 1 {
			break
		}
		if f.seats[r][c] == 2 {
			occ += 1
			break
		}
	}
	
	// ABC
	// D E
	// FGH
	return occ
}

func (f *ferry) next() *ferry {
	n := make([][]int, f.rows)
	for i := range f.seats {
		n[i] = make([]int, f.cols)
	}
	for i := 0; i < f.rows; i++ {
		for j := 0; j < f.cols; j++ {
			if f.seats[i][j] == 0 {
				continue
			}
			occ := f.adjOcc(i, j)
			if f.seats[i][j] == 1 && occ == 0 {
				n[i][j] = 2
				continue
			}
			if f.seats[i][j] == 2 && occ >= 5 {
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
