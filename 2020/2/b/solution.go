package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pwCheck struct {
	f, s int
	check rune
	pw string
}

func main() {
	input := ReadInput()
	pws := []pwCheck{}
	for _, l := range(input) {
		parts := strings.Split(l, " ")
		rang := strings.Split(parts[0], "-")
		f, err := strconv.Atoi(rang[0])
		if err != nil {
			panic(err)
		}
		s, err := strconv.Atoi(rang[1])
		if err != nil {
			panic(err)
		}
		check := parts[1]
		
		newPW := pwCheck{
			f: f,
			s: s,
			check: []rune(check)[0],
			pw: parts[2],
		}
		pws = append(pws, newPW)
	}
	
	numValid := 0
	for _, pw := range(pws) {
		valid := false
		runes := []rune(pw.pw)
		if runes[pw.f - 1] == pw.check {
			valid = !valid
		}
		if runes[pw.s - 1] == pw.check {
			valid = !valid
		}
		if valid {
			numValid += 1
		}
	}
	fmt.Printf("Num valid: %d", numValid)
}


func log(l string) {
	fmt.Println(l)
}

func ReadFakeInput() []string {
	input := []string{"1-3 a: abcde",
"1-3 b: cdefg",
"2-9 c: ccccccccc"}
	return input
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
