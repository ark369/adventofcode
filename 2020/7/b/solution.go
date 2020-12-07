package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bagAndCount struct {
	bag string
	count int
}

func main() {
	input := ReadInput()
	bagMap := map[string][]bagAndCount{}
	for _, l := range(input) {
		pars := strings.Split(l, " bags contain ")
		k := pars[0]
		vs := strings.Split(pars[1], ", ")
		vals := []bagAndCount{}
		for _, v := range(vs) {
			bac := strings.Split(v, " ")
			c, err := strconv.Atoi(bac[0])
			if err != nil {
				break
			}
			b := fmt.Sprintf("%s %s", bac[1], bac[2])
			vals = append(vals, bagAndCount{b, c})
		}
		bagMap[k] = vals
	}
	fmt.Printf("%d", contains("shiny gold", &bagMap))
}

func contains(bag string, bagMap *map[string][]bagAndCount) int {
	total := 0
	for _, bac := range((*bagMap)[bag]) {
		total += bac.count * (1 + contains(bac.bag, bagMap))
	}	
	return total
}

func ReadFakeInput() []string {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
