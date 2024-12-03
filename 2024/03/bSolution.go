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

	state := ""
	sum := 0
	left := 0
	right := 0

	enabled := true

	for _, l := range input {
		for _, ch := range l {
			print := false
			if state == "" {
				if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				}
			} else if state == "d" {
				if ch == 'o' {
					state = "do"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "do" {
				if ch == 'n' {
					state = "don"
				} else if ch == '(' {
					state = "do("
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "do(" {
				if ch == ')' {
					state = ""
					enabled = true
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "don" {
				if ch == '\'' {
					state = "don'"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "don'" {
				if ch == 't' {
					state = "don't"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "don't" {
				if ch == '(' {
					state = "don't("
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "don't(" {
				if ch == ')' {
					state = ""
					enabled = false
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "m" {
				if ch == 'u' {
					state = "mu"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "mu" {
				if ch == 'l' {
					state = "mul"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "mul" {
				if ch == '(' {
					state = "mul("
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					state = ""
				}
			} else if state == "mul(" {
				if ch == '0' {
					left = 0
					state = "mul(XXX"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					if ch >= '1' && ch <= '9' {
						left = int(ch - '0')
						state = "mul(X"
					} else {
						state = ""
					}
				}
			} else if state == "mul(X" {
				if ch >= '0' && ch <= '9' {
					left = left*10 + int(ch-'0')
					state = "mul(XX"
				} else if ch == ',' {
					state = "mul(XXX,"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					left = 0
					state = ""
				}
			} else if state == "mul(XX" {
				if ch >= '0' && ch <= '9' {
					left = left*10 + int(ch-'0')
					state = "mul(XXX"
				} else if ch == ',' {
					state = "mul(XXX,"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					left = 0
					state = ""
				}
			} else if state == "mul(XXX" {
				if ch == ',' {
					state = "mul(XXX,"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					left = 0
					state = ""
				}
			} else if state == "mul(XXX," {
				if ch == '0' {
					right = 0
					state = "mul(XXX,XXX"
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					if ch >= '1' && ch <= '9' {
						right = int(ch - '0')
						state = "mul(XXX,X"
					} else {
						state = ""
					}
				}
			} else if state == "mul(XXX,X" {
				if ch >= '0' && ch <= '9' {
					right = right*10 + int(ch-'0')
					state = "mul(XXX,XX"
				} else if ch == ')' {
					if enabled {
						sum += left * right
					}
					left = 0
					right = 0
					state = ""
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					left = 0
					right = 0
					state = ""
				}
			} else if state == "mul(XXX,XX" {
				if ch >= '0' && ch <= '9' {
					right = right*10 + int(ch-'0')
					state = "mul(XXX,XXX"
				} else if ch == ')' {
					if enabled {
						sum += left * right
					}
					left = 0
					right = 0
					state = ""
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					left = 0
					right = 0
					state = ""
				}
			} else if state == "mul(XXX,XXX" {
				if ch == ')' {
					if enabled {
						sum += left * right
					}
					left = 0
					right = 0
					state = ""
				} else if ch == 'm' {
					state = "m"
				} else if ch == 'd' {
					state = "d"
				} else {
					left = 0
					right = 0
					state = ""
				}
			}
			if print {
				fmt.Printf("%c %v %s %d %d %d\n", ch, enabled, state, left, right, sum)
			}
		}
	}

	fmt.Print(sum)
}

func ReadFakeInput() []string {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
