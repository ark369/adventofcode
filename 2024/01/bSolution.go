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

type Counter struct {
	locations []int
	curr, max int
}

func MakeCounter() *Counter {
	c := &Counter{}
	c.max = 100000
	c.locations = make([]int, c.max)
	return c
}

func (c *Counter) Insert(i int) {
	c.locations[i]++
}

func (c *Counter) Next() int {
	for c.locations[c.curr] == 0 {
		c.curr++
		if c.curr == c.max {
			return -1
		}
	}
	c.locations[c.curr]--
	return c.curr
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	c1 := MakeCounter()
	c2 := MakeCounter()

	num := 0

	for _, l := range input {
		parsed := strings.Split(l, "   ")
		c1.Insert(Atoi(parsed[0]))
		c2.Insert(Atoi(parsed[1]))
		num++
	}

	sum := 0
	for i := 0; i < num; i++ {
		c1next := c1.Next()
		c2val := c2.locations[c1next]
		sim := c1next * c2val
		sum += sim
		//fmt.Printf("%d %d: %d\n", c1next, c2val, sim)
	}
	fmt.Print(sum)
}

func ReadFakeInput() []string {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
