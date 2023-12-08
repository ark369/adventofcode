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

type Node struct {
	id          string
	left, right *Node
}

func (n *Node) String() string {
	// AAA = (BBB, BBB)
	return fmt.Sprintf("%s = (%s, %s)", n.id, n.left.id, n.right.id)
}



func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	m := make(map[string]*Node)

	steps := ""

	for _, l := range input {
		if steps == "" {
			steps = l
			continue
		}
		if l == "" {
			continue
		}
		id := strings.Split(l, " = ")[0]
		m[id] = &Node{id, nil, nil}
	}

	for i, l := range input {
		if i < 2 {
			continue
		}
		id := strings.Split(l, " = ")[0]
		targets := strings.Split(strings.Split(l, " = ")[1], ", ")
		left := targets[0][1:]
		m[id].left = m[left]
		right := targets[1][:3]
		m[id].right = m[right]
	}

	curr := m["AAA"]
	n := 0
	for curr.id != "ZZZ" {
		i := n % len(steps)
		dir := steps[i : i+1]
		if dir == "L" {
			curr = curr.left
		} else {
			curr = curr.right
		}
		//fmt.Printf("n: %d, i: %d, dir: %s, curr.id: %s\n", n, i, dir, curr.id)
		n++
	}

	fmt.Printf("n: %d", n)
}

func ReadFakeInput() []string {
	input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
