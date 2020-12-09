package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ring struct {
	len int
	nums []int
	curr int
}

func newRing(l int) *ring {
	return &ring{l, make([]int, l), 0}
}

func (r *ring)insert(n int) {
	r.nums[r.curr] = n
	r.curr += 1
	if r.curr == r.len {
		r.curr = 0
	}
}

func (r *ring)isValid(n int) bool {
	m := map[int]int{}
	for _, rn := range r.nums {
		if _, ok := m[rn]; ok {
			return true
		}
		m[n - rn] = rn
	}
	return false
}

func main() {
	num := 25
	r := newRing(num)
	input := ReadInput()
	for i, l := range(input) {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		if i < num {
			r.nums[i] = n
		} else {
			if r.isValid(n) {
				r.insert(n)
			} else {
				fmt.Printf("Invalid %d", n)
				return
			}
		}
	}
	fmt.Printf("%+v", r)
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
