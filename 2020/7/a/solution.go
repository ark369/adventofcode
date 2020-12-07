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
	revBagMap := map[string][]string{}
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
			if rv, ok := revBagMap[b]; ok {
				revBagMap[b] = append(rv, k)
			} else {
				revBagMap[b] = []string{k}
			}
		}
		bagMap[k] = vals
	}
	check := []string{"shiny gold"}
	possible := map[string]bool{}
	for len(check) > 0 {
		next := check[0]
		check = check[1:]
		if _, ok := possible[next]; ok {
			continue
		}
		if next != "shiny gold" {
			possible[next] = true
		}
		for _, newCheck := range(revBagMap[next]) {
			check = append(check, newCheck)
		}
	}
	fmt.Printf("%d", len(possible))
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
