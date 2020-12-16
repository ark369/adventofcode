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

type possibilities struct {
	posToRule map[int][]bool
	posToRuleCount map[int]int
}

func (p *possibilities) eliminate(pos, rule int) {
	if p.posToRuleCount[pos] == 1 {
		return
	}
	p.posToRule[pos][rule] = true
	p.posToRuleCount[pos] = p.posToRuleCount[pos] - 1
	if p.posToRuleCount[pos] == 1 {
		remaining := p.getRemaining(pos)
		for i := 0; i < len(p.posToRule[pos]); i++ {
			if i == pos {
				continue
			}
			p.eliminate(i, remaining)
		}
	}
}

func (p *possibilities) getRemaining(pos int) int {
	for i, b := range p.posToRule[pos] {
		if !b {
			return i
		}
	}
	panic(fmt.Sprintf("Should be at least one remaining for p: %+v\n", p))
}

func (p *possibilities) done() bool {
	for _, v := range p.posToRuleCount {
		if v > 1 {
			return false
		}
	}
	return true
}

func update(p *possibilities, finalRuleToPos, finalPosToRule map[int]int) {
	for pos, num := range p.posToRuleCount {
		if num == 1 {
			rule := p.getRemaining(pos)
			finalRuleToPos[rule] = pos
			finalPosToRule[pos] = rule
		}
	}
}

func getMyDepartureProduct(myTicket ticket, rules []rule, finalRuleToPos map[int]int) int {
	ret := 1
	for i, rule := range rules {
		if !strings.HasPrefix(rule.name, "departure") {
			continue
		}
		ret *= myTicket.nums[finalRuleToPos[i]]
	}
	return ret
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
	
	validTickets := []ticket{myTicket}
	for _, t := range nearbyTickets {
		if inv := t.invalidNums(rules); len(inv) == 0 {
			validTickets = append(validTickets, t)
		}
	}
	
	p := &possibilities{
		posToRule: make(map[int][]bool),
		posToRuleCount: make(map[int]int),
	}
	l := len(myTicket.nums)
	for i := 0; i < l; i++ {
		p.posToRule[i] = make([]bool, l)
		p.posToRuleCount[i] = l
	}
	
	finalRuleToPos := make(map[int]int)
	finalPosToRule := make(map[int]int)
	for !p.done() {
		for _, t := range validTickets {
			for pos, val := range t.nums {
				if _, ok := finalPosToRule[pos]; ok {
					continue
				}
				for r, rule := range rules {
					if _, ok := finalRuleToPos[r]; ok {
						continue
					}
					if !rule.isValidFor(val) {
						p.eliminate(pos, r)
						if p.posToRuleCount[pos] == 1 {
							update(p, finalRuleToPos, finalPosToRule)
							break
						}
					}
				}
			}
		}
	}
	
	v := getMyDepartureProduct(myTicket, rules, finalRuleToPos)

	fmt.Printf("v: %d\n", v)
	fmt.Printf("rules: %v\n", rules)
	fmt.Printf("myTicket: %v\n", myTicket)
	fmt.Printf("validTickets: %v\n", validTickets)
	fmt.Printf("p: %+v\n", p)
	fmt.Printf("finalRuleToPos: %v\n", finalRuleToPos)
	fmt.Printf("finalPosToRule: %v\n", finalPosToRule)
}


func ReadFakeInput() []string {
	input := `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
