package main

import (
	"fmt"
	"strconv"
	"strings"
)

type contiguous struct {
	sum, highest, lowest, start, end int
}

func main() {
	target := 552655238
	input := ReadInput()
	//target := 127
	//input := ReadFakeInput()
	cs := []contiguous{}
	for i, l := range(input) {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("[%d]%d: %v\n", i, n, cs)
		if found, c := check(&cs, n, target, i); found {
			l := c.lowest
			h := c.highest
			fmt.Printf("%d + %d = %d\n", l, h, l + h)
			fmt.Printf("sum: %d, %v", c.sum, input[c.start:c.end])
			return
		}
		//if (n > 500) { return }
	}
}

func check(cs *[]contiguous, n, chec, ind int) (bool, contiguous) {
	newCs := []contiguous{}
	for _, c := range *cs {
		c.sum += n
		c.end += 1
		if n < c.lowest {
			c.lowest = n
		}
		if n > c.highest {
			c.highest = n
		}
		if c.sum == chec {
			return true, c
		}
		if c.sum < chec {
			newCs = append(newCs, c)
		}
	}
	newCs = append(newCs, contiguous{n, n, n, ind, ind})
	*cs = newCs
	return false, contiguous{}
}

func ReadFakeInput() []string {
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
