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

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	most := 0

	currTotal := 0
	for _, l := range input {
		if l == "" {
			if currTotal > most {
				most = currTotal
			}
			currTotal = 0
			continue
		}
		currTotal += Atoi(l)
	}

	fmt.Printf("most: %d\n", most)
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
