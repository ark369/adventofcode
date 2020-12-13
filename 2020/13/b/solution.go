package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// https://brilliant.org/wiki/chinese-remainder-theorem/
	input := ReadInput()
	//input := ReadFakeInput()
	busses := strings.Split(input[1], ",")
	ms := map[int]int{}
	ids := []int{}
	for i, bus := range busses {
		if bus != "x" {
			id, _ := strconv.Atoi(bus)
			ids = append(ids, id)
			ms[id] = i
		}
	//fmt.Printf("%s > %d, %d, %d, %d\n", l, b.x, b.y, b.wx, b.wy)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	for _, id := range ids {
		fmt.Printf("%d: %d\n", id, ms[id])
	}

	a := ids[0]
	b := calcB(ms[a], a)
	for i := 1; i < len(ids); i++ {
		m: = ids[i]
		k: = calcB(ms[m], m)
		fmt.Printf("a: %d, b: %d, k: %d, m: %d\n", a, b, k, m)
		t: = solve(a, b, k, m)
		b += a * t
		a *= m
	}
	fmt.Printf("%d", b)
}

// given
//   x + b2 === 0 mod m
// return b where
//   x === b mod m
func calcB(b2, m int) int {
	b := b2 * -1
	for ; b < 0; b += m {}
	return b
}

// given
//   ax + b === k mod m
// return t where
//   x === t mod m
func solve(a, b, k, m int) int {
	for i := 0; i < m; i++ {
		if (i*a + b) % m == k {
			return i
		}
	}
	panic(fmt.Sprintf("no solution found for solve(%d, %d, %d, %d)\n", a, b, k, m))
}

func ReadFakeInput() []string {
	input := `939
7,13,x,x,59,x,31,19`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := `1005162
19,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,823,x,x,x,x,x,x,x,23,x,x,x,x,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,443,x,x,x,x,x,37,x,x,x,x,x,x,13`
	return strings.Split(input, "\n")
}
