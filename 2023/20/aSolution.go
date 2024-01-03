package main

import (
	"fmt"
	"sort"
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

type ModuleType int

const (
	BROADCAST ModuleType = iota
	FLIPFLOP
	CONJUNCTION
	OUTPUT
)

type Module struct {
	name           string
	dest           []string
	typ            ModuleType
	ffOn           bool
	conjMemoryHigh map[string]bool
}

func MakeFlipFlop(name string, dest []string) *Module {
	m := &Module{}
	m.name = name
	m.dest = dest
	m.typ = FLIPFLOP
	return m
}

func MakeConjunction(name string, dest []string) *Module {
	m := &Module{}
	m.name = name
	m.dest = dest
	m.typ = CONJUNCTION
	m.conjMemoryHigh = make(map[string]bool)
	return m
}

func (m *Module) Signal(high bool, src string) []ToSend {
	next := []ToSend{}
	switch m.typ {
	case BROADCAST:
		for _, d := range m.dest {
			next = append(next, ToSend{m.name, d, high})
		}
	case OUTPUT:
	case FLIPFLOP:
		if !high {
			if m.ffOn {
				m.ffOn = false
				for _, d := range m.dest {
					next = append(next, ToSend{m.name, d, false})
				}
			} else {
				m.ffOn = true
				for _, d := range m.dest {
					next = append(next, ToSend{m.name, d, true})
				}
			}
		}
	case CONJUNCTION:
		m.conjMemoryHigh[src] = high
		allHigh := true
		for _, h := range m.conjMemoryHigh {
			if !h {
				allHigh = false
				break
			}
		}
		for _, d := range m.dest {
			next = append(next, ToSend{m.name, d, !allHigh})
		}
	}
	return next
}

func (m *Module) Hash() string {
	switch m.typ {
	case FLIPFLOP:
		if m.ffOn {
			return "1"
		} else {
			return "0"
		}
	case CONJUNCTION:
		keys := []string{}
		for k, _ := range m.conjMemoryHigh {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		h := ""
		for _, k := range keys {
			h += ","
			h += k
			if m.conjMemoryHigh[k] {
				h += "1"
			} else {
				h += "0"
			}
		}
		return h
	default:
		panic("hashing bad typ")
	}
}

type Machine struct {
	modules map[string]*Module
}

func (m *Machine) Hash() string {
	keys := []string{}
	for k, _ := range m.modules {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	h := ""
	for _, k := range keys {
		if k == "broadcaster" || m.modules[k].typ == OUTPUT {
			continue
		}
		h += k
		h += m.modules[k].Hash()
		h += "/"
	}
	return h
}

func MakeMachine() *Machine {
	m := &Machine{}
	m.modules = make(map[string]*Module)
	return m
}

func (m *Machine) ParseModule(l string) {
	pieces := strings.Split(l, " -> ")
	src := pieces[0]
	dests := strings.Split(pieces[1], ", ")
	if src == "broadcaster" {
		mod := &Module{}
		mod.name = "broadcaster"
		mod.dest = dests
		mod.typ = BROADCAST
		m.modules[src] = mod
	} else {
		if src[0] == '%' {
			m.modules[src[1:]] = MakeFlipFlop(src[1:], dests)
		} else if src[0] == '&' {
			m.modules[src[1:]] = MakeConjunction(src[1:], dests)
		} else {
			panic("Unrecognized src typ")
		}
	}
}

func (m *Machine) MaybeInitConjunction(l string) {
	pieces := strings.Split(l, " -> ")
	src := pieces[0]
	dests := strings.Split(pieces[1], ", ")
	for _, d := range dests {
		if _, ok := m.modules[d]; !ok {
			//fmt.Printf("d not found: %s\n", d)

			mod := &Module{}
			mod.name = d
			mod.typ = OUTPUT
			m.modules[d] = mod
		}
		if m.modules[d].typ == CONJUNCTION {
			s := src
			if src != "broadcaster" {
				s = src[1:]
			}
			m.modules[d].conjMemoryHigh[s] = false
		}
	}
}

type ToSend struct {
	src, dst string
	high     bool
}

func (t ToSend) String() string {
	strength := "low"
	if t.high {
		strength = "high"
	}
	return fmt.Sprintf("%s -%s-> %s", t.src, strength, t.dst)
}

func (m *Machine) PushButton() (int, int) {
	lowPulses := 0
	highPulses := 0
	pending := []ToSend{ToSend{"button", "broadcaster", false}}
	for len(pending) > 0 {
		curr := pending[0]
		//fmt.Printf("%s\n", curr)
		pending = pending[1:]
		next := m.modules[curr.dst].Signal(curr.high, curr.src)
		if curr.high {
			highPulses++
		} else {
			lowPulses++
		}
		pending = append(pending, next...)
	}
	return lowPulses, highPulses
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	m := MakeMachine()

	for _, l := range input {
		m.ParseModule(l)
	}

	for _, l := range input {
		m.MaybeInitConjunction(l)
	}

	lowSum := 0
	highSum := 0

	seen := make(map[string]bool)
	seen[m.Hash()] = true

	i := 1
	for i = 1; i <= 1000; i++ {
		low, high := m.PushButton()
		lowSum += low
		highSum += high
		h := m.Hash()
		if _, ok := seen[h]; ok {
			break
		}
		seen[h] = true
	}
	rem := int(1000 % i)

	if i <= 1000 {
		div := int(1000 / i)
		lowSum *= div
		highSum *= div

		for j := 0; j < rem; j++ {
			low, high := m.PushButton()
			lowSum += low
			highSum += high
		}
	}

	fmt.Printf("i: %d, rem: %d, low: %d, high: %d, low x high: %d\n", i, rem, lowSum, highSum, lowSum*highSum)
}

func ReadFakeInput() []string {
	input := `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

	return strings.Split(input, "\n")
}

func ReadFakeInput2() []string {
	input := `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
