package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput()
	var total int
	for _, l := range(input) {
		dims := strings.Split(l, "x")
		w, err := strconv.Atoi(dims[0])
		if err != nil {
			panic(err)
		}
		h, err := strconv.Atoi(dims[1])
		if err != nil {
			panic(err)
		}
		l, err := strconv.Atoi(dims[2])
		if err != nil {
			panic(err)
		}
		a := w * h
		b := w * l
		c := h * l
		var slack int
		if a < b && a < c {
			slack = a
		} else if b < c {
			slack = b
		} else {
			slack = c
		}
		total += 2 * (a + b + c) + slack
	}
	fmt.Printf("Total %d", total)
}

func ReadFakeInput() []string {
	input := `2x3x4`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
