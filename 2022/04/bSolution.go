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

type TwoJobs struct {
	a1, a2, b1, b2 int
}

func MakeTwoJobs(s string) *TwoJobs {
	parts := strings.Split(s, ",")
	a := strings.Split(parts[0], "-")
	b := strings.Split(parts[1], "-")
	return &TwoJobs{Atoi(a[0]), Atoi(a[1]), Atoi(b[0]), Atoi(b[1])}
}

func (t *TwoJobs) FullyContained() bool {
	if t.a1 >= t.b1 && t.a2 <= t.b2 {
		return true
	}
	if t.a1 <= t.b1 && t.a2 >= t.b2 {
		return true
	}
	return false
}

func (t *TwoJobs) Overlap() bool {
	if t.a1 <= t.b2 && t.a1 >= t.b1 {
		return true
	}
	if t.b1 <= t.a2 && t.b1 >= t.a1 {
		return true
	}
	return false
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	sum := 0
	for _, l := range input {
		t := MakeTwoJobs(l)
		if t.Overlap() {
			sum++
		}
	}

	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
