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

type Hand struct {
	cards []rune
	typ   int
	bid   int
	next  *Hand
}

func MakeHand(s string) *Hand {
	h := &Hand{}
	parts := strings.Split(s, " ")
	h.bid = Atoi(parts[1])
	m := make(map[rune]int)
	for _, c := range parts[0] {
		h.cards = append(h.cards, c)
		m[c]++
	}
	num := len(m)
	if num == 1 {
		// 5 of a kind
		h.typ = 6
	} else if num == 2 {
		for _, v := range m {
			if v == 1 || v == 4 {
				// 4 of a kind
				h.typ = 5
			} else {
				// full house
				h.typ = 4
			}
			break
		}
	} else if num == 3 {
		for _, v := range m {
			if v == 3 {
				// 3 of a kind
				h.typ = 3
				break
			} else if v == 2 {
				// 2 pair
				h.typ = 2
				break
			}
		}
	} else if num == 4 {
		// pair
		h.typ = 1
	} else {
		// singles
		h.typ = 0
	}
	return h
}

func ToRank(c rune) int {
	switch c {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	panic("unrecorgnized rank")
}

func IsLower(c1, c2 rune) bool {
	return ToRank(c1) < ToRank(c2)
}

func (h1 *Hand) IsLowerThan(h2 *Hand) bool {
	for i := 0; i < len(h1.cards); i++ {
		if h1.cards[i] == h2.cards[i] {
			continue
		}
		if IsLower(h1.cards[i], h2.cards[i]) {
			return true
		} else {
			return false
		}
	}
	return false
}

func (h *Hand) String() string {
	s := fmt.Sprintf("typ: %d, ", h.typ)
	s += "cards: "
	for _, c := range h.cards {
		s += fmt.Sprintf("%s", string(c))
	}
	s += fmt.Sprintf(" (%d)", h.bid)
	return s
}

type Hands struct {
	ordered [7]*Hand
}

func MakeHands() *Hands {
	h := &Hands{}
	h.ordered = [7]*Hand{}
	return h
}

func (hands *Hands) AddHand(h *Hand) {
	head := hands.ordered[h.typ]
	if head == nil {
		hands.ordered[h.typ] = h
	} else {
		var prev *Hand
		curr := head
		inserted := false
		for curr != nil {
			if h.IsLowerThan(curr) {
				if prev == nil {
					hands.ordered[h.typ] = h
				} else {
					prev.next = h
				}
				h.next = curr
				inserted = true
				break
			}
			prev = curr
			curr = curr.next
		}
		if !inserted {
			prev.next = h
		}
	}
}

func (hands *Hands) String() string {
	s := ""
	for i := 0; i < 7; i++ {
		curr := hands.ordered[i]
		hStr := ""
		for curr != nil {
			hStr += fmt.Sprintf("%s, ", curr)
			curr = curr.next
		}
		s += fmt.Sprintf("%d: %s\n", i, hStr)
	}
	return s
}

func (hands *Hands) Score() int {
	rank := 1
	score := 0
	for i := 0; i < 7; i++ {
		curr := hands.ordered[i]
		for curr != nil {
			score += rank * curr.bid
			rank++
			curr = curr.next
		}
	}
	return score
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	hands := MakeHands()
	for _, l := range input {
		h := MakeHand(l)
		hands.AddHand(h)
	}
	//fmt.Printf("%s\n", hands)
	fmt.Printf("score: %d\n", hands.Score())
}

func ReadFakeInput() []string {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
