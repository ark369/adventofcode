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

func hash(s string) int {
	h := 0
	for _, c := range s {
		iValue := int(c)
		h += iValue
		h *= 17
		h %= 256
	}
	return h
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	sum := 0

	for _, l := range input {
		codes := strings.Split(l, ",")
		for _, code := range codes {
			h := hash(code)
			sum += h
		}
	}

	fmt.Printf("sum: %d", sum)
}

func ReadFakeInput() []string {
	input := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
