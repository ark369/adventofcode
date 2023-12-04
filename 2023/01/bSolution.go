package main

import (
	"fmt"
	"strings"
)

type NumTracker struct {
	target string
	val    int
	soFar  string
}

func NewNumTracker(target string, val int) *NumTracker {
	return &NumTracker{target, val, ""}
}

func (t *NumTracker) Add(c rune) int {
	next := t.soFar + string(c)
	if next == t.target {
		t.soFar = ""
		return t.val
	} else if strings.HasPrefix(t.target, next) {
		t.soFar = next
	} else if next == "nini" && t.val == 9 {
		t.soFar = "ni"
	} else if strings.HasPrefix(t.target, string(c)) {
		t.soFar = string(c)
	} else {
		t.soFar = ""
	}
	return -1
}

func (t *NumTracker) String() string {
	return fmt.Sprintf("target: %s, val: %d, soFar: %s", t.target, t.val, t.soFar)
}

type Tracker struct {
	trackers []*NumTracker
}

func NewTracker() *Tracker {
	t := &Tracker{}
	t.trackers = append(t.trackers, NewNumTracker("one", 1))
	t.trackers = append(t.trackers, NewNumTracker("two", 2))
	t.trackers = append(t.trackers, NewNumTracker("three", 3))
	t.trackers = append(t.trackers, NewNumTracker("four", 4))
	t.trackers = append(t.trackers, NewNumTracker("five", 5))
	t.trackers = append(t.trackers, NewNumTracker("six", 6))
	t.trackers = append(t.trackers, NewNumTracker("seven", 7))
	t.trackers = append(t.trackers, NewNumTracker("eight", 8))
	t.trackers = append(t.trackers, NewNumTracker("nine", 9))
	return t
}

func (t *Tracker) Add(c rune) int {
	ret := -1
	curr := -1
	for _, nt := range t.trackers {
		curr = nt.Add(c)
		if curr != -1 {
			ret = curr
		}
	}
	return ret
}

func (t *Tracker) String() string {
	str := ""
	for _, nt := range t.trackers {
		str += fmt.Sprintf("%s\n", nt)
	}
	return str
}

type Values struct {
	first int
	last  int
}

func NewValues() *Values {
	return &Values{-1, -1}
}

func (v *Values) set(val int) {
	if v.first == -1 {
		v.first = val
	}
	v.last = val
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	sum := 0
	for _, l := range input {
		v := NewValues()
		t := NewTracker()
		//fmt.Printf("l: %s\n", l)
		for _, c := range l {
			//fmt.Printf("c: %c\n", c)
			val := t.Add(c)
			//fmt.Printf("t: %s\n", t)
			//fmt.Printf("val: %d\n", val)
			if val != -1 {
				v.set(val)
			} else if c >= '0' && c <= '9' {
				val := int(c - '0')
				v.set(val)
			}
		}
		//fmt.Printf("first: %d, last: %d\n", v.first, v.last)
		sum += v.first*10 + v.last
	}
	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
