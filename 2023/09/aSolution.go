package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Oasis struct {
	row []int
}

func MakeOasis(s string) *Oasis {
	nums := strings.Split(s, " ")
	o := &Oasis{}
	for _, n := range nums {
		o.row = append(o.row, Atoi(n))
	}
	return o
}

func (o *Oasis) Calculate() int {
	size := len(o.row)
	var rows [][]int
	rows = append(rows, o.row)
	allZero := false
	for !allZero {
		prev := rows[len(rows)-1]
		rows = append(rows, make([]int, size-1))
		size--
		curr := rows[len(rows)-1]
		allZero = true
		for i, p := range prev {
			if i < size {
				diff := prev[i+1] - p
				curr[i] = diff
				if diff != 0 {
					allZero = false
				}
			}
			//fmt.Printf("%d  ", p)
		}
		//fmt.Printf("\n")
	}
	nextVal := 0
	for i := len(rows) - 1; i >= 0; i-- {
		nextVal += rows[i][len(rows[i])-1]
	}
	return nextVal
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	sum := 0

	for _, l := range input {
		o := MakeOasis(l)
		sum += o.Calculate()
		//fmt.Printf("sum: %d\n", sum)
	}

	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
