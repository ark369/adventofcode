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

type Springs struct {
	conditions string
	damaged    []int
}

func MakeSprings(l string) *Springs {
	s := &Springs{}
	s.conditions = strings.Split(l, " ")[0]
	for _, n := range strings.Split(strings.Split(l, " ")[1], ",") {
		s.damaged = append(s.damaged, Atoi(n))
	}
	return s
}

func (s *Springs) GetPossibilities() int {
	ret := CalculatePossibilities(s.conditions, s.damaged, 0)

	return ret
}

func CalculatePossibilities(remaining string, damaged []int, currentTally int) int {
	if len(damaged) == 0 {
		for _, c := range remaining {
			if c == '#' {
				return 0
			}
		}
		return 1
	}
	if len(remaining) == 0 {
		return 0
	}
	if currentTally == 0 {
		c := remaining[0]
		remaining = remaining[1:]
		if c == '.' {
			return CalculatePossibilities(remaining, damaged, 0)
		} else if c == '#' {
			if damaged[0] == 1 {
				damaged = damaged[1:]
				if len(remaining) == 0 {
					if len(damaged) == 0 {
						return 1
					} else {
						return 0
					}
				} else {
					if remaining[0] == '#' {
						return 0
					} else {
						remaining = remaining[1:]
						return CalculatePossibilities(remaining, damaged, 0)
					}
				}
			} else {
				return CalculatePossibilities(remaining, damaged, 1)
			}
		} else if c == '?' {
			ifNotDamaged := CalculatePossibilities(remaining, damaged, 0)
			ifDamaged := 0
			if damaged[0] == 1 {
				damaged = damaged[1:]
				if len(remaining) == 0 {
					if len(damaged) == 0 {
						ifDamaged = 1
					}
				} else {
					if remaining[0] != '#' {
						remaining = remaining[1:]
						ifDamaged = CalculatePossibilities(remaining, damaged, 0)
					}
				}
			} else {
				ifDamaged = CalculatePossibilities(remaining, damaged, 1)
			}
			return ifNotDamaged + ifDamaged
		}
	} else {
		currDamaged := damaged[0]
		if currDamaged < 2 {
			panic("Assertion failed, currDamaged is < 2")
		}
		c := remaining[0]
		remaining = remaining[1:]
		if c == '.' {
			return 0
		} else if currentTally+1 == currDamaged {
			damaged = damaged[1:]
			if len(remaining) == 0 {
				if len(damaged) == 0 {
					return 1
				} else {
					return 0
				}
			} else {
				if remaining[0] == '#' {
					return 0
				} else {
					remaining = remaining[1:]
					return CalculatePossibilities(remaining, damaged, 0)
				}
			}
		} else {
			return CalculatePossibilities(remaining, damaged, currentTally+1)
		}
	}
	panic("Reached end of CalculatePossibilities")
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	sum := 0

	for _, l := range input {
		s := MakeSprings(l)
		possibilities := s.GetPossibilities()
		sum += possibilities
		//fmt.Printf("l: %s, possibilities: %d\n", l, possibilities)
	}

	fmt.Printf("sum: %d\n", sum)
}

func ReadFakeInput() []string {
	input := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
