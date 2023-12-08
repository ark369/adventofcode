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

func AllOnZ(nodes []*StartAndCurr) bool {
	for _, n := range nodes {
		if string(n.curr.id[2]) != "Z" {
			return false
		}
	}
	return true
}

type StartAndCurr struct {
	start string
	curr  *Node
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	m := make(map[string]*Node)

	steps := ""

	curr := []*StartAndCurr{}
	idToTimesZSeen := make(map[string][]int)

	for _, l := range input {
		if steps == "" {
			steps = l
			continue
		}
		if l == "" {
			continue
		}
		id := strings.Split(l, " = ")[0]
		n := &Node{id, nil, nil}
		m[id] = n
		if string(id[2]) == "A" {
			curr = append(curr, &StartAndCurr{n.id, n})
			idToTimesZSeen[n.id] = []int{}
		}
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

	time := 0
	for time < 100000 {
		i := time % len(steps)
		dir := steps[i : i+1]
		time++
		if dir == "L" {
			for _, n := range curr {
				n.curr = n.curr.left
				if string(n.curr.id[2]) == "Z" {
					idToTimesZSeen[n.start] = append(idToTimesZSeen[n.start], time)
				}
			}
		} else {
			for _, n := range curr {
				n.curr = n.curr.right
				if string(n.curr.id[2]) == "Z" {
					idToTimesZSeen[n.start] = append(idToTimesZSeen[n.start], time)
				}
			}
		}
		//fmt.Printf("n: %d, i: %d, dir: %s, curr.id: %s\n", n, i, dir, curr.id)
	}

	fmt.Printf("%v\n", idToTimesZSeen)

	firsts := []int{}
	for _, v := range idToTimesZSeen {
		firsts = append(firsts, v[0])
	}

	lcm := LCM(firsts...)

	fmt.Printf("lcm: %d\n", lcm)
	fmt.Printf("time: %d", time)
}

func ReadFakeInput() []string {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
