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

type Top3 struct {
	first, second, third int
}

func (t *Top3) Record(n int) {
	if n > t.first {
		t.third = t.second
		t.second = t.first
		t.first = n
	} else if n > t.second {
		t.third = t.second
		t.second = n
	} else if n > t.third {
		t.third = n
	}
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	top3 := &Top3{}
	currTotal := 0
	for _, l := range input {
		if l == "" {
			top3.Record(currTotal)
			currTotal = 0
			continue
		}
		currTotal += Atoi(l)
	}

	fmt.Printf("sum: %d\n", top3.first+top3.second+top3.third)
}

func ReadFakeInput() []string {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
