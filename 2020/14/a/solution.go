package main

import (
	"fmt"
	"strconv"
	"strings"
)

func apply(mask string, val int) int {
	l := len(mask)
	orMask := make([]rune, l)
	andMask := make([]rune, l)
	for i, r := range mask {
		if r == 'X' {
			orMask[i] = '0'
			andMask[i] = '1'
		} else {
			orMask[i] = r
			andMask[i] = '0'
		}
	}
	aMask, _ := strconv.ParseUint(string(andMask), 2, l)
	oMask, _ := strconv.ParseUint(string(orMask), 2, l)
	orig := uint64(val)
	ret := int((orig & aMask) | oMask)
	//fmt.Printf("aMask: %0*b\noMask: %0*b\n orig: %0*b\n. ret: %0*b\n", l, aMask, l, oMask, l, orig, l, ret)
	return ret
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	mem := map[int]int{}
	mask := ""
	for _, l := range input {
		if l[:4] == "mask" {
			mask = l[7:]
			//fmt.Printf(" mask: %s\n", mask)
			continue
		}
		parts := strings.Split(l, " = ")
		loc := parts[0][4:]
		loc = loc[:len(loc)-1]
		locInt, _ := strconv.Atoi(loc)
		val, _ := strconv.Atoi(parts[1])
		ret := apply(mask, val)
		mem[locInt] = ret
		//fmt.Printf("mem[%d]: %d (val: %d)\n", locInt, ret, val)
	}
	sum := 0
	for _, v := range mem {
		sum += v
	}
	fmt.Printf("sum: %d\n", sum)
}


func ReadFakeInput() []string {
	input := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
