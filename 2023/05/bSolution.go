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

type RangeMap struct {
	d, s, r int
	next    *RangeMap
}

func (rm *RangeMap) String() string {
	return fmt.Sprintf("[%d,%d (%d)]", rm.s, rm.s+rm.r-1, rm.d)
}

func (rm *RangeMap) get(i int) int {
	if i >= rm.s && i <= rm.s+rm.r-1 {
		return i - rm.s + rm.d
	}
	return -1
}

type RangeMaps struct {
	head *RangeMap
}

func (rms *RangeMaps) String() string {
	ret := ""
	curr := rms.head
	for curr != nil {
		ret += fmt.Sprintf("%s, ", curr)
		curr = curr.next
	}
	return ret
}

func (rms *RangeMaps) AddRangeMap(s string) {
	parts := strings.Split(s, " ")
	rm := &RangeMap{}
	rm.d = Atoi(parts[0])
	rm.s = Atoi(parts[1])
	rm.r = Atoi(parts[2])

	var prev *RangeMap
	curr := rms.head
	for curr != nil {
		if rm.s < curr.s {
			if prev == nil {
				rms.head = rm
			} else {
				prev.next = rm
			}
			rm.next = curr
			return
		}
		prev = curr
		curr = curr.next
	}
	if prev == nil {
		rms.head = rm
	} else {
		prev.next = rm
	}
}

func (rms *RangeMaps) MapThrough(r *Range) []*Range {
	//fmt.Printf("MapThrough received %s\n", r)
	newRanges := []*Range{}
	curr := rms.head
	for true {
		if curr == nil {
			//fmt.Printf("biggest\n")
			newRanges = append(newRanges, &Range{r.i, r.j, nil})
			break
		}
		if r.i < curr.s {
			if r.j < curr.s {
				//fmt.Printf("smaller\n")
				newRanges = append(newRanges, &Range{r.i, r.j, nil})
				break
			} else {
				//fmt.Printf("subset smaller\n")
				newRanges = append(newRanges, &Range{r.i, curr.s - 1, nil})
				r.i = curr.s
			}
		} else if r.i < curr.s+curr.r {
			if r.j < curr.s+curr.r {
				//fmt.Printf("contained\n")
				newRanges = append(newRanges, &Range{r.i - curr.s + curr.d, r.j - curr.s + curr.d, nil})
				break
			} else {
				//fmt.Printf("subset contained\n")
				newRanges = append(newRanges, &Range{r.i - curr.s + curr.d, -curr.s + curr.r - 1 + curr.d, nil})
				r.i = curr.s + curr.r
				curr = curr.next
			}
		} else {
			curr = curr.next
		}
	}
	//fmt.Printf("MapThrough returning %v\n", newRanges)
	return newRanges
}

type Range struct {
	i, j int
	next *Range
}

func (r *Range) String() string {
	return fmt.Sprintf("[%d,%d]", r.i, r.j)
}

type Ranges struct {
	head *Range
}

func (rs *Ranges) ParseStartSeeds(s string) {
	parts := strings.Split(s, " ")
	newRanges := []*Range{}
	for i := 1; i < len(parts); i = i + 2 {
		r := &Range{}
		r.i = Atoi(parts[i])
		r.j = r.i + Atoi(parts[i+1]) - 1
		newRanges = append(newRanges, r)
	}
	rs.head = SortRanges(newRanges)
}

func SortRanges(ranges []*Range) *Range {
	//fmt.Printf("SortRanges received %v\n", ranges)
	var head *Range
	for _, r := range ranges {
		var prev *Range
		curr := head
		inserted := false
		for curr != nil {
			if r.i < curr.i {
				if prev == nil {
					head = r
				} else {
					prev.next = r
				}
				r.next = curr
				inserted = true
				break
			}
			prev = curr
			curr = curr.next
		}
		if !inserted {
			if prev == nil {
				head = r
			} else {
				prev.next = r
			}
		}
	}
	//fmt.Printf("SortRanges returning %s\n", head)
	return head
}

func (rs *Ranges) MapThrough(rms *RangeMaps) {
	newRanges := []*Range{}
	curr := rs.head
	for curr != nil {
		newRanges = append(newRanges, rms.MapThrough(&Range{curr.i, curr.j, nil})...)
		curr = curr.next
	}
	rs.head = SortRanges(newRanges)
}

func (rs *Ranges) String() string {
	ret := ""
	curr := rs.head
	for curr != nil {
		ret += fmt.Sprintf("%s, ", curr)
		curr = curr.next
	}
	return ret
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()
	rs := &Ranges{}
	rs.ParseStartSeeds(input[0])

	seed2soil := &RangeMaps{}
	soil2fert := &RangeMaps{}
	fert2water := &RangeMaps{}
	water2light := &RangeMaps{}
	light2temp := &RangeMaps{}
	temp2humid := &RangeMaps{}
	humid2loc := &RangeMaps{}

	var currRM *RangeMaps
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
			currRM.AddRangeMap(l)
			//fmt.Printf("currRM: %s\n", currRM)
		}
	}

	//fmt.Printf("%s\n", rs)
	//fmt.Printf("seed2soil: %s\n", seed2soil)
	rs.MapThrough(seed2soil)
	//fmt.Printf("%s\n", rs)
	//fmt.Printf("soil2fert: %s\n", soil2fert)
	rs.MapThrough(soil2fert)
	//fmt.Printf("%s\n", rs)
	//fmt.Printf("fert2water: %s\n", fert2water)
	rs.MapThrough(fert2water)
	//fmt.Printf("%s\n", rs)
	rs.MapThrough(water2light)
	//fmt.Printf("%s\n", rs)
	rs.MapThrough(light2temp)
	//fmt.Printf("%s\n", rs)
	rs.MapThrough(temp2humid)
	//fmt.Printf("%s\n", rs)
	rs.MapThrough(humid2loc)
	//fmt.Printf("%s\n", rs)

	lowest := rs.head.i

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
