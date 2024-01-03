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

type Range struct {
	min, max int
}

func (r *Range) Clone() *Range {
	ret := &Range{}
	ret.min = r.min
	ret.max = r.max
	return ret
}

func Split(r *Range, c Condition) (*Range, *Range) {
	var matching *Range
	var nonMatching *Range
	if c.comp == "<" {
		if r.max < c.target {
			matching = r.Clone()
		} else if r.min >= c.target {
			nonMatching = r.Clone()
		} else {
			matching = &Range{r.min, c.target - 1}
			nonMatching = &Range{c.target, r.max}
		}
	} else {
		if r.min > c.target {
			matching = r.Clone()
		} else if r.max <= c.target {
			nonMatching = r.Clone()
		} else {
			matching = &Range{c.target + 1, r.max}
			nonMatching = &Range{r.min, c.target}
		}
	}
	return matching, nonMatching
}

func CloneRanges(m map[string]*Range) map[string]*Range {
	rangeMap := make(map[string]*Range)
	for k, v := range m {
		rangeMap[k] = v.Clone()
	}
	return rangeMap
}

func PossibleRangeValues(m map[string]*Range) int {
	sum := 1
	for _, v := range m {
		sum *= (v.max - v.min) + 1
	}
	return sum
}

func Calculate(rangeMap map[string]*Range, rulesMap map[string][]Rule, ruleKey string) int {
	rules := rulesMap[ruleKey]
	sum := 0
	if ruleKey == "R" {
		return 0
	}
	if ruleKey == "A" {
		return PossibleRangeValues(rangeMap)
	}
	for _, rule := range rules {
		cond := rule.cond
		field := cond.field
		if field == "" {
			sum += Calculate(rangeMap, rulesMap, rule.dest)
			continue
		}
		matching, nonMatching := Split(rangeMap[field], cond)
		if matching != nil {
			newRanges := CloneRanges(rangeMap)
			newRanges[field] = matching
			sum += Calculate(newRanges, rulesMap, rule.dest)
		}
		if nonMatching != nil {
			rangeMap[field] = nonMatching
			continue
		}
	}
	return sum
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	rulesMap := make(map[string][]Rule)

	for _, l := range input {
		//fmt.Printf(l + "\n")
		if l == "" {
			break
		}
		k, rules := ParseRules(l)
		rulesMap[k] = rules
	}

	rangeMap := make(map[string]*Range)
	rangeMap["x"] = &Range{1, 4000}
	rangeMap["m"] = &Range{1, 4000}
	rangeMap["a"] = &Range{1, 4000}
	rangeMap["s"] = &Range{1, 4000}

	sum := Calculate(rangeMap, rulesMap, "in")

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
