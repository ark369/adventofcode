package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	val int
	children []*node
}

func (n *node)addChild(c *node) {
	n.children = append(n.children, c)
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	//input := []string{"1","2","3"}
	ads := []int{0}
	for _, l := range(input) {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ads = append(ads, n)
	}
	sort.Ints(ads)
	ads = append(ads, ads[len(ads) - 1] + 3)
	nodes := []*node{}
	for i, a := range ads {
		n := &node{a, []*node{}}
		nodes = append(nodes, n)
		for j := i - 1; j > i - 4; j-- {
			if j < 0 {
				break
			}
			if a - nodes[j].val > 3 {
				break
			}
			nodes[j].addChild(n)
		}
	}
	
	paths := map[int]int{}
	comb := numPaths(nodes[0], paths)
	fmt.Printf("comb: %d\n", comb)
	//fmt.Printf("%+v\n", paths)
}

func numPaths(n *node, paths map[int]int) int {
	//fmt.Printf("numPaths for n.val: %d\n", n.val)
	if p, ok := paths[n.val]; ok {
		return p
	}
	//fmt.Printf("numPaths for len(n.children): %v\n", len(n.children))
	if len(n.children) == 0 {
		paths[n.val] = 1
		return 1
	}
	p := 0
	for _, c := range n.children {
		p += numPaths(c, paths)
	}
	paths[n.val] = p
	return p
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
