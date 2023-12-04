package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	sum := 0
	for _, l := range input {
		first := -1
		last := -1
		for _, c := range l {
			if c >= '0' && c <= '9' {
				n := int(c - '0')
				if first == -1 {
					first = n
				}
				last = n
			}
		}
		//fmt.Printf("first: %d, last: %d\n", first, last)
		sum += first*10 + last
	}
	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
