package main

import (
	"fmt"
	"math"
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

func Reverse(s string) string {
	ret := ""
	for _, c := range s {
		ret = string(c) + ret
	}
	return ret
}

func ToSnafu(n int) string {
	s := ""
	curr := n + 2
	for curr > 5 {
		rem := curr % 5
		switch rem {
		case 0:
			s = "=" + s
		case 1:
			s = "-" + s
		case 2:
			s = "0" + s
		case 3:
			s = "1" + s
		case 4:
			s = "2" + s
		}
		curr = (curr / 5) + 2
	}
	rem := curr % 5
	switch rem {
	case 0:
		s = "=" + s
	case 1:
		s = "-" + s
	case 2:
		s = "0" + s
	case 3:
		s = "1" + s
	case 4:
		s = "2" + s
	}
	return s
}

func FromSnafu(s string) int {
	ret := 0
	flip := Reverse(s)
	for i, c := range flip {
		base := int(math.Pow(5, float64(i)))
		var mult int
		switch c {
		case '=':
			mult = -2
		case '-':
			mult = -1
		case '0':
			mult = 0
		case '1':
			mult = 1
		case '2':
			mult = 2
		}
		ret += base * mult
	}
	return ret
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	sum := 0
	for _, l := range input {
		n := FromSnafu(l)
		sum += n
		//fmt.Printf("l: %s, %d\n", l, n)
	}
	fmt.Printf("sum: %d, s: %s", sum, ToSnafu(sum))
}

func ReadFakeInput() []string {
	input := `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
