package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getAddresses(mask string, loc int) []int {
	adds := []string{""}
	locStr := fmt.Sprintf("%036s", strconv.FormatUint(uint64(loc), 2))
	//fmt.Printf("  mask: %s\nlocStr: %s (loc: %d)\n", mask, locStr, loc)
	for i, r := range mask {
		newA := []string{}
		if r == '0' {
			for _, a := range adds {
				newA = append(newA, a + string(locStr[i]))
			}
		} else if r == '1' {
			for _, a := range adds {
				newA = append(newA, a + "1")
			}
		} else {
			for _, a := range adds {
				newA = append(newA, a + "0")
				newA = append(newA, a + "1")
			}
		}
		adds = newA
		//fmt.Printf("%v\n", adds)
	}
	as := make([]int, len(adds))
	for i, a := range adds {
		aUint, _ := strconv.ParseUint(a, 2, len(mask))
		as[i] = int(aUint)
	}
	return as
}

func appendToAdds(r rune, loc int, adds []string) []string {
	ret := []string{}
	if r == '0' {
		for _, a := range adds {
			ret = append(ret, a + string(r))
		}
	} else {
		for _, a := range adds {
			ret = append(ret, a + "0")
			ret = append(ret, a + "1")
		}
	}
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
		adds := getAddresses(mask, locInt)
		for _, a := range adds {
			mem[a] = val
		}
		//fmt.Printf("mem[%d]: %d (val: %d)\n", locInt, ret, val)
	}
	sum := 0
	for _, v := range mem {
		sum += v
	}
	fmt.Printf("sum: %d\n", sum)
}


func ReadFakeInput() []string {
	input := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
