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

func ParseStartSeeds(s string) []int {
	parts := strings.Split(s, " ")
	ret := []int{}
	for i := 1; i < len(parts); i++ {
		ret = append(ret, Atoi(parts[i]))
	}
	return ret
}

type Range struct {
	d, s, r int
}

func (r *Range) get(i int) int {
	if i >= r.s && i <= r.s+r.r-1 {
		return i - r.s + r.d
	}
	return -1
}

type RangeMap struct {
	mappings []*Range
}

func (rm *RangeMap) AddRange(s string) {
	parts := strings.Split(s, " ")
	r := &Range{Atoi(parts[0]), Atoi(parts[1]), Atoi(parts[2])}
	rm.mappings = append(rm.mappings, r)
}

func (rm *RangeMap) get(i int) int {
	for _, r := range rm.mappings {
		n := r.get(i)
		if n != -1 {
			return n
		}
	}
	return i
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()
	seeds := ParseStartSeeds(input[0])
	//fmt.Printf("seeds: %v\n", seeds)

	seed2soil := &RangeMap{}
	soil2fert := &RangeMap{}
	fert2water := &RangeMap{}
	water2light := &RangeMap{}
	light2temp := &RangeMap{}
	temp2humid := &RangeMap{}
	humid2loc := &RangeMap{}

	var currRM *RangeMap
	for i := 2; i < len(input); i++ {
		l := input[i]
		if l == "" {
			continue
		} else if l == "seed-to-soil map:" {
			currRM = seed2soil
		} else if l == "soil-to-fertilizer map:" {
			currRM = soil2fert
		} else if l == "fertilizer-to-water map:" {
			currRM = fert2water
		} else if l == "water-to-light map:" {
			currRM = water2light
		} else if l == "light-to-temperature map:" {
			currRM = light2temp
		} else if l == "temperature-to-humidity map:" {
			currRM = temp2humid
		} else if l == "humidity-to-location map:" {
			currRM = humid2loc
		} else {
			currRM.AddRange(l)
		}
	}

	lowest := -1
	for _, s := range seeds {
		soil := seed2soil.get(s)
		f := soil2fert.get(soil)
		w := fert2water.get(f)
		l := water2light.get(w)
		t := light2temp.get(l)
		h := temp2humid.get(t)
		loc := humid2loc.get(h)

		if lowest == -1 {
			lowest = loc
		} else if loc < lowest {
			lowest = loc
		}
	}

	fmt.Printf("lowest: %d\n", lowest)
}

func ReadFakeInput() []string {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
