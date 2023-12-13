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

func (s *Springs) String() string {
	damaged := ""
	for i, d := range s.damaged {
		damaged += fmt.Sprintf("%d", d)
		if i < len(s.damaged)-1 {
			damaged += ","
		}
	}
	return fmt.Sprintf("%s %s", s.conditions, damaged)
}

func MakeSprings(l string, additionalFolds int) *Springs {
	conditions := strings.Split(l, " ")[0]
	damaged := strings.Split(l, " ")[1]
	conditionsCopy := conditions
	damagedCopy := damaged
	for i := 0; i < additionalFolds; i++ {
		conditions += "?" + conditionsCopy
		damaged += "," + damagedCopy
	}
	s := &Springs{}
	s.conditions = conditions
	for _, n := range strings.Split(damaged, ",") {
		s.damaged = append(s.damaged, Atoi(n))
	}
	return s
}

func (s *Springs) GetPossibilities() int {
	memo := make(map[string]int)
	ret := CalculatePossibilities(s.conditions, s.damaged, 0, memo)

	return ret
}

func MinRequired(damaged []int, currentTally int) int {
	minRequired := len(damaged) - 1
	for _, d := range damaged {
		minRequired += d
	}
	minRequired -= currentTally
	return minRequired
}

func GetMemoKey(remaining string, currentTally int, damaged []int) string {
	return fmt.Sprintf("%s-%d-%v", remaining, currentTally, damaged)
}

func CalculatePossibilities(remaining string, damaged []int, currentTally int, memo map[string]int) int {
	memoKey := GetMemoKey(remaining, currentTally, damaged)
	if v, ok := memo[memoKey]; ok {
		return v
	}
	if len(damaged) == 0 {
		for _, c := range remaining {
			if c == '#' {
				memo[memoKey] = 0
				return 0
			}
		}
		memo[memoKey] = 1
		return 1
	}
	if len(remaining) == 0 {
		memo[memoKey] = 0
		return 0
	}
	if len(remaining) < MinRequired(damaged, currentTally) {
		memo[memoKey] = 0
		return 0
	}
	if currentTally == 0 {
		c := remaining[0]
		remaining = remaining[1:]
		for c == '.' {
			if len(remaining) == 0 {
				return 0
			}
			c = remaining[0]
			remaining = remaining[1:]
		}
		if c == '#' {
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
						v := CalculatePossibilities(remaining, damaged, 0, memo)
						memo[memoKey] = v
						return v
					}
				}
			} else {
				v := CalculatePossibilities(remaining, damaged, 1, memo)
				memo[memoKey] = v
				return v
			}
		} else if c == '?' {
			ifNotDamaged := CalculatePossibilities(remaining, damaged, 0, memo)
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
						ifDamaged = CalculatePossibilities(remaining, damaged, 0, memo)
					}
				}
			} else {
				ifDamaged = CalculatePossibilities(remaining, damaged, 1, memo)
			}
			v := ifNotDamaged + ifDamaged
			memo[memoKey] = v
			return v
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
					v := CalculatePossibilities(remaining, damaged, 0, memo)
					memo[memoKey] = v
					return v
				}
			}
		} else {
			v := CalculatePossibilities(remaining, damaged, currentTally+1, memo)
			memo[memoKey] = v
			return v
		}
	}
	panic("Reached end of CalculatePossibilities")
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	//????????#??#????????????????#??#???????
	sum := 0

	for _, l := range input {
		s := MakeSprings(l, 4)
		possibilities := s.GetPossibilities()
		sum += possibilities
		fmt.Printf("s: %s, possibilities: %d\n", s, possibilities)
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

func ReadFakeInput2() []string {
	//input := `?.???..??? 2,1`
	input := `???.### 1,1,3
.??..??...?##. 1,1,3
?.???..??? 2,1`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
