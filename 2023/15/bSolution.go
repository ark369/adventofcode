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

func hash(s string) int {
	h := 0
	for _, c := range s {
		iValue := int(c)
		h += iValue
		h *= 17
		h %= 256
	}
	return h
}

type Lens struct {
	id          string
	focalLength int
	prev, next  *Lens
}

func MakeLens(id string, focalLength int) *Lens {
	l := &Lens{}
	l.id = id
	l.focalLength = focalLength
	return l
}

func (l *Lens) SpliceSelf() {
	prev := l.prev
	next := l.next
	l.prev = nil
	l.next = nil
	if prev == nil {
		if next == nil {
			return
		} else {
			next.prev = nil
		}
	} else {
		if next == nil {
			prev.next = nil
		} else {
			prev.next = next
			next.prev = prev
		}
	}
}

type Box struct {
	boxNum      int
	first, last *Lens
}

func (b *Box) AddLens(l *Lens) {
	if b.first == nil {
		b.first = l
		b.last = l
		return
	}
	b.last.next = l
	l.prev = b.last
	b.last = l
}

func (b *Box) Remove(id string) {
	curr := b.first
	for curr != nil {
		if curr.id == id {
			if curr == b.first {
				b.first = curr.next
			}
			if curr == b.last {
				b.last = curr.prev
			}
			curr.SpliceSelf()
			return
		}
		curr = curr.next
	}
}

func (b *Box) Power() int {
	curr := b.first
	slot := 1
	power := 0
	for curr != nil {
		currPower := b.boxNum + 1
		currPower *= slot
		currPower *= curr.focalLength
		power += currPower
		curr = curr.next
		slot++
	}
	return power
}

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	sum := 0

	boxes := make([]*Box, 256)
	for i := 0; i < 256; i++ {
		b := &Box{}
		b.boxNum = i
		boxes[i] = b
	}

	lensMap := make(map[string]*Lens)

	for _, l := range input {
		codes := strings.Split(l, ",")
		for _, code := range codes {
			if strings.Contains(code, "=") {
				id := strings.Split(code, "=")[0]
				boxNum := hash(id)
				focalLength := Atoi(strings.Split(code, "=")[1])
				if l, ok := lensMap[id]; ok {
					l.focalLength = focalLength
				} else {
					l := MakeLens(id, focalLength)
					lensMap[id] = l
					boxes[boxNum].AddLens(l)
				}
			} else {
				id := strings.Split(code, "-")[0]
				boxNum := hash(id)
				if _, ok := lensMap[id]; ok {
					boxes[boxNum].Remove(id)
					delete(lensMap, id)
				}
			}
		}
	}

	for _, box := range boxes {
		sum += box.Power()
	}

	fmt.Printf("sum: %d", sum)
}

func ReadFakeInput() []string {
	input := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
