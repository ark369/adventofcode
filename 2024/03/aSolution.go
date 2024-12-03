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

func main() {
	input := ReadInput()
	//input := ReadFakeInput()

	state := ' '
	sum := 0
	left := 0
	right := 0

	for _, l := range input {
		for _, ch := range l {
			if state == ' ' {
				if ch == 'm' {
					state = 'm'
				}
			} else if state == 'm' {
				if ch == 'u' {
					state = 'u'
				} else {
					state = ' '
				}
			} else if state == 'u' {
				if ch == 'l' {
					state = 'l'
				} else {
					state = ' '
				}
			} else if state == 'l' {
				if ch == '(' {
					state = '('
				} else {
					state = ' '
				}
			} else if state == '(' {
				if ch == '0' {
					left = 0
					state = '3'
				} else {
					if ch >= '1' && ch <= '9' {
						left = int(ch - '0')
						state = '1'
					} else {
						state = ' '
					}
				}
			} else if state == '1' {
				if ch >= '0' && ch <= '9' {
					left = left*10 + int(ch-'0')
					state = '2'
				} else if ch == ',' {
					state = ','
				} else {
					left = 0
					state = ' '
				}
			} else if state == '2' {
				if ch >= '0' && ch <= '9' {
					left = left*10 + int(ch-'0')
					state = '3'
				} else if ch == ',' {
					state = ','
				} else {
					left = 0
					state = ' '
				}
			} else if state == '3' {
				if ch == ',' {
					state = ','
				} else {
					left = 0
					state = ' '
				}
			} else if state == ',' {
				if ch == '0' {
					right = 0
					state = 'c'
				} else {
					if ch >= '1' && ch <= '9' {
						right = int(ch - '0')
						state = 'a'
					} else {
						state = ' '
					}
				}
			} else if state == 'a' {
				if ch >= '0' && ch <= '9' {
					right = right*10 + int(ch-'0')
					state = 'b'
				} else if ch == ')' {
					sum += left * right
					left = 0
					right = 0
					state = ' '
				} else {
					left = 0
					right = 0
					state = ' '
				}
			} else if state == 'b' {
				if ch >= '0' && ch <= '9' {
					right = right*10 + int(ch-'0')
					state = 'c'
				} else if ch == ')' {
					sum += left * right
					left = 0
					right = 0
					state = ' '
				} else {
					left = 0
					right = 0
					state = ' '
				}
			} else if state == 'c' {
				if ch == ')' {
					sum += left * right
					left = 0
					right = 0
					state = ' '
				} else {
					left = 0
					right = 0
					state = ' '
				}
			}
			//fmt.Printf("%c %c %d %d %d\n", ch, state, left, right, sum)
		}
	}

	fmt.Print(sum)
}

func ReadFakeInput() []string {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
