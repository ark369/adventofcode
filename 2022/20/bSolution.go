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
	val, index int
	prev, next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%d<%d>%d [%d]", n.prev.index, n.index, n.next.index, n.val)
}

func MakeNode(s string, i int) *Node {
	return &Node{Atoi(s) * 811589153, i, nil, nil}
}

func (n *Node) ShiftRight(k int) {
	for i := 0; i < k; i++ {
		// cleave it out
		left := n.prev
		right := n.next
		left.next = right
		right.prev = left
		// add it back
		left = right
		right = right.next
		left.next = n
		n.prev = left
		right.prev = n
		n.next = right
	}
}

func (n *Node) ShiftLeft(k int) {
	for i := 0; i < k; i++ {
		// cleave it out
		left := n.prev
		right := n.next
		left.next = right
		right.prev = left
		// add it back
		right = left
		left = left.prev
		left.next = n
		n.prev = left
		right.prev = n
		n.next = right
	}
}

type RingWithOriginalOrder struct {
	zero, head, tail *Node
	order            []*Node
}

func (r *RingWithOriginalOrder) Insert(n *Node) {
	if n.val == 0 {
		r.zero = n
	}
	r.order = append(r.order, n)
	if r.head == nil {
		r.head = n
		r.tail = n
		n.next = n
		n.prev = n
	} else {
		n.prev = r.tail
		r.tail.next = n
		n.next = r.head
		r.head.prev = n
		r.tail = n
	}
}

func (r *RingWithOriginalOrder) Decrypt() (int, int, int) {
	for i := 0; i < 10; i++ {
		//fmt.Printf("%s\n", r)
		for _, n := range r.order {
			if n.val > 0 {
				if n == r.head {
					r.head = n.next
				}
				n.ShiftRight(n.val % (len(r.order) - 1))
			} else if n.val < 0 {
				if n == r.head {
					r.head = n.prev
				}
				n.ShiftLeft((-1 * n.val) % (len(r.order) - 1))
			}
			//fmt.Printf("%s\n", r)
		}
	}
	var a, b, c int
	curr := r.zero
	for i := 1; i <= 3000; i++ {
		curr = curr.next
		if i == 1000 {
			a = curr.val
		} else if i == 2000 {
			b = curr.val
		} else if i == 3000 {
			c = curr.val
		}
	}
	return a, b, c
}

func (r *RingWithOriginalOrder) String() string {
	s := ""
	curr := r.head
	for i := 0; i < len(r.order); i++ {
		s += fmt.Sprintf("%v, ", curr)
		curr = curr.next
	}
	return s
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	r := &RingWithOriginalOrder{}

	for i, l := range input {
		r.Insert(MakeNode(l, i))
	}

	a, b, c := r.Decrypt()

	fmt.Printf("a, b, c: %d, %d, %d\nsum := %d\n", a, b, c, a+b+c)
}

func ReadFakeInput() []string {
	input := `1
2
-3
3
-2
0
4`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
