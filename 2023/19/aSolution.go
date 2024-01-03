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

type Part struct {
	x, m, a, s int
	accepted   bool
}

func (p *Part) Value() int {
	return p.x + p.m + p.a + p.s
}

func (p *Part) PassesCondition(c Condition) bool {
	switch c.field {
	case "":
		return true
	case "x":
		if c.comp == "<" {
			return p.x < c.target
		} else {
			return p.x > c.target
		}
	case "m":
		if c.comp == "<" {
			return p.m < c.target
		} else {
			return p.m > c.target
		}
	case "a":
		if c.comp == "<" {
			return p.a < c.target
		} else {
			return p.a > c.target
		}
	case "s":
		if c.comp == "<" {
			return p.s < c.target
		} else {
			return p.s > c.target
		}
	}
	panic("!!!")
}

func (p *Part) ExecuteRules(rules []Rule) string {
	for _, r := range rules {
		if p.PassesCondition(r.cond) {
			return r.dest
		}
	}
	panic("Did not pass any conditions")
}

func ExecuteRules(p *Part, rulesMap map[string][]Rule) {
	k := "in"
	for {
		r := rulesMap[k]
		result := p.ExecuteRules(r)
		if result == "A" {
			p.accepted = true
			break
		} else if result == "R" {
			break
		} else {
			k = result
		}
	}
}

func ParsePart(l string) *Part {
	p := &Part{}
	//{x=787,m=2655,a=1222,s=2876}
	l = strings.Trim(l, "{}")
	pieces := strings.Split(l, ",")
	for _, piece := range pieces {
		k := strings.Split(piece, "=")[0]
		val := Atoi(strings.Split(piece, "=")[1])
		switch k {
		case "x":
			p.x = val
		case "m":
			p.m = val
		case "a":
			p.a = val
		case "s":
			p.s = val
		}
	}
	return p
}

type Condition struct {
	field  string
	comp   string
	target int
}

type Rule struct {
	cond Condition
	dest string
}

func ParseRule(p string) Rule {
	r := Rule{}
	parts := strings.Split(p, ":")
	if len(parts) == 1 {
		r.dest = p
	} else {
		r.dest = parts[1]
		c := parts[0]
		cond := Condition{}
		if strings.Contains(c, "<") {
			cond.field = strings.Split(c, "<")[0]
			cond.target = Atoi(strings.Split(c, "<")[1])
			cond.comp = "<"
		} else {
			cond.field = strings.Split(c, ">")[0]
			cond.target = Atoi(strings.Split(c, ">")[1])
			cond.comp = ">"
		}
		r.cond = cond
	}
	return r
}

func ParseRules(l string) (string, []Rule) {
	k := strings.Split(l, "{")[0]
	rules := []Rule{}

	pieces := strings.Split(strings.Split(strings.Split(l, "{")[1], "}")[0], ",")
	for _, p := range pieces {
		rules = append(rules, ParseRule(p))
	}

	return k, rules
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	readingRules := true

	rulesMap := make(map[string][]Rule)
	parts := []*Part{}

	for _, l := range input {
		//fmt.Printf(l + "\n")
		if l == "" {
			readingRules = false
			continue
		}
		if readingRules {
			k, rules := ParseRules(l)
			rulesMap[k] = rules
		} else {
			parts = append(parts, ParsePart(l))
		}
	}

	for _, p := range parts {
		ExecuteRules(p, rulesMap)
	}

	sum := 0
	for _, p := range parts {
		if p.accepted {
			sum += p.Value()
		}
	}

	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
