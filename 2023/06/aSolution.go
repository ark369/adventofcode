package main

import (
	"fmt"
	"math"
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

func ParseLine(s string) []int {
	nums := strings.Split(strings.Split(s, ":")[1], " ")
	ret := []int{}
	for _, n := range nums {
		if n != "" {
			ret = append(ret, Atoi(n))
		}
	}
	return ret
}

func waysToWin(t, d int) int {
	ways := 0
	if t%2 == 0 {
		if (t/2)*(t/2) > d {
			ways++
		}
	}
	for i := int(math.Ceil(float64(t)/2)) - 1; i > 0; i-- {
		if i*(t-i) > d {
			ways += 2
		} else {
			break
		}
	}
	return ways
}

func main() {
	//input := ReadInput()
	input := ReadFakeInput()

	times := ParseLine(input[0])
	dists := ParseLine(input[1])

	fmt.Printf("times: %v\n", times)
	fmt.Printf("dists: %v\n", dists)

	val := 1

	for i := 0; i < len(times); i++ {
		numWaysToWin := waysToWin(times[i], dists[i])
		fmt.Printf("times[i]: %d, dists[i]: %d, numWaysToWin: %d\n", times[i], dists[i], numWaysToWin)
		val *= numWaysToWin
	}

	fmt.Printf("val: %d\n", val)
}

func ReadFakeInput() []string {
	input := `Time:      7  15   30
Distance:  9  40  200`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
