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

type Report struct {
	levels []int
}

func MakeReport(s string) *Report {
	r := &Report{}
	parts := strings.Split(s, " ")
	for _, v := range parts {
		r.levels = append(r.levels, Atoi(v))
	}
	return r
}

func (r *Report) IsSafe() bool {
	increasing := true
	prev := r.levels[0]
	curr := r.levels[1]
	if curr == prev {
		return false
	} else if curr < prev {
		increasing = false
	}
	for i := 1; i < len(r.levels); i++ {
		curr = r.levels[i]
		if increasing {
			if curr < prev+1 || curr > prev+3 {
				return false
			} else {
				prev = curr
			}
		} else {
			if curr < prev-3 || curr > prev-1 {
				return false
			} else {
				prev = curr
			}
		}
	}
	return true
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	num := 0

	for _, l := range input {
		r := MakeReport(l)
		if r.IsSafe() {
			num++
		}
	}

	fmt.Print(num)
}

func ReadFakeInput() []string {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
