package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	ads := []int{}
	for _, l := range(input) {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ads = append(ads, n)
	}
	sort.Ints(ads)
	curr := 0
	oneDiffs := 0
	threeDiffs := 1
	for _, a := range ads {
		if a - curr == 1 {
			oneDiffs += 1
		}
		if a - curr == 3 {
			threeDiffs += 1
		}
		curr = a
	}
	fmt.Printf("ones: %d, threes: %d, product: %d", oneDiffs, threeDiffs, oneDiffs * threeDiffs)
}

func ReadFakeInput() []string {
	input := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
