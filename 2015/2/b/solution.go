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
		bow := w * h * l
		var biggest int
		if w > h && w > l {
			biggest = w
		} else if h > l {
			biggest = h
		} else {
			biggest = l
		}
		total += 2 * (w + h + l - biggest) + bow
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
