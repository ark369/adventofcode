package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type sortedInts struct {
	ints        []int
	first, last int
}

func (s *sortedInts) insertIndexFor(i int) int {
	f := 0
	l := len(s.ints)
	log(fmt.Sprintf("    insertIndexFor %d, f=%d, l=%d", i, f, l))
	for l-f > 0 {
		mid := ((l - f) / 2) + f
		log(fmt.Sprintf("    insertIndexFor %d, mid=%d", i, mid))
		if i == s.ints[mid] {
			l = mid
			f = mid
		} else if i < s.ints[mid] {
			l = mid
		} else {
			f = mid + 1
		}
	}
	return l
}

func (s *sortedInts) insert(i int) {
	ind := s.insertIndexFor(i)
	if ind == len(s.ints) {
		s.last = i
	}
	if ind == 0 {
		s.first = i
	}
	log(fmt.Sprintf("  Insert %d at %d", i, ind))
	post := make([]int, len(s.ints[ind:]))
	copy(post, s.ints[ind:])
	s.ints = append(s.ints[:ind], i)
	s.ints = append(s.ints, post...)
}

func (s *sortedInts) below(target int) []int {
	ind := s.insertIndexFor(target)
	return s.ints[:ind]
}

func main() {
	input := ReadInput()
	ints := make([]int, len(input))
	target := 2020
	targets := make(map[int]int)
	for ind, l := range input {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ints[ind] = i
	}
	sort.Ints(ints)
	for iInd, i := range(ints[:len(ints)]) {
		log(fmt.Sprintf("i=%d", i))
		if j, ok := targets[i]; ok {
			k := target - i - j
			fmt.Printf("Found %d x %d x %d = %d", i, j, k, i*j*k)
			return
		}
		if iInd < len(ints) && i + ints[iInd + 1] > target - 1 {
			continue
		}
		for _, j := range(ints[iInd + 1:]) {
			k := target - i - j
			if k < j - 1 {
				break
			}
			log(fmt.Sprintf("  j=%d", j))
			targets[k] = i
		}
		log(fmt.Sprintf("    targets: %v", targets))
	}
	fmt.Printf("NOT FOUND")
}

func main2() {
	input := ReadInput()
	target := 2020
	seenList := &sortedInts{}
	diffsTwo := make(map[int]int)
	for _, l := range input {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		log(fmt.Sprintf("See %d", i))
		if j, ok := diffsTwo[i]; ok {
			k := target - i - j
			fmt.Printf("Found %d x %d x %d = %d", i, j, k, i*j*k)
			return
		}
		for _, k := range seenList.below(target - i) {
			log(fmt.Sprintf("  Checking diff of %d", k))
			if target-i-k > 0 {
				log(fmt.Sprintf("    Setting target of %d", target-i-k))
				diffsTwo[i] = target - i - k
			}
		}
		seenList.insert(i)
		log(fmt.Sprintf("  seenList.ints: %v", seenList.ints))
	}
	fmt.Printf("NOT FOUND")
}

func log(l string) {
	fmt.Println(l)
}

func ReadFakeInput() []string {
	input := []string{"979", "366", "675"}
	return input
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
