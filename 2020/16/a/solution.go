package main

import (
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
	name string
	lowA, highA, lowB, highB int
}

func parseRule(s string) rule {
	parts := strings.Split(s, ": ")
	ranges := strings.Split(parts[1], " or ")
	a := strings.Split(ranges[0], "-")
	b := strings.Split(ranges[1], "-")
	la, _ := strconv.Atoi(a[0])
	ha, _ := strconv.Atoi(a[1])
	lb, _ := strconv.Atoi(b[0])
	hb, _ := strconv.Atoi(b[1])
	return rule{parts[0], la, ha, lb, hb}
}

type ticket struct {
	nums []int
}

func parseTicket(s string) ticket {
	parts := strings.Split(s, ",")
	nums := make([]int, len(parts))
	for i, p := range parts {
		n, _ := strconv.Atoi(p)
		nums[i] = n
	}
	return ticket{nums}
}

func (t ticket) invalidNums(rules []rule) []int {
	invalid := []int{}
	
	for _, n := range t.nums {
		valid := false
		for _, r := range rules {
			if r.isValidFor(n) {
				valid = true
				break
			}
		}
		if !valid {
			invalid = append(invalid, n)
		}
	}
	
	return invalid
}

func (r rule) isValidFor(n int) bool {
	if (n >= r.lowA && n <= r.highA) || (n >= r.lowB && n <= r.highB) {
		return true
	}
	return false
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	rules := []rule{}
	var myTicket ticket
	nearbyTickets := []ticket{}
	mode := 0
	for _, l := range input {
		if l == "" {
			mode += 1
			continue
		}
		if mode == 0 {
			rules = append(rules, parseRule(l))
		}
		if mode == 1 {
			if l == "your ticket:" {
				continue
			}
			myTicket = parseTicket(l)
		}
		if mode == 2 {
			if l == "nearby tickets:" {
				continue
			}
			nearbyTickets = append(nearbyTickets, parseTicket(l))
		}
	}
	
	errRate := 0
	for _, t := range nearbyTickets {
		inv := t.invalidNums(rules)
		for _, i := range inv {
			errRate += i
		}
	}

	//fmt.Printf("rules: %v\n", rules)
	fmt.Printf("myTicket: %v\n", myTicket)
	//fmt.Printf("nearbyTickets: %v\n", nearbyTickets)
	fmt.Printf("errRate: %d\n", errRate)
}


func ReadFakeInput() []string {
	input := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
