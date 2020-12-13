package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	input := ReadInput()
	//input := ReadFakeInput()
	id := 0
	startTime, _ := strconv.Atoi(input[0])
	mins := startTime
	busses := strings.Split(input[1], ",")
	ids := []int{}
	for _, bus := range busses {
		if bus != "x" {
			id, _ := strconv.Atoi(bus)
			ids = append(ids, id)
		}
	//fmt.Printf("%s > %d, %d, %d, %d\n", l, b.x, b.y, b.wx, b.wy)
	}
	for _, i := range ids {
		min := (startTime / i + 1) * (i) - startTime
		if min < mins {
			mins = min
			id = i
		}
	}
	fmt.Printf("%d x %d = %d\n", mins, id, mins * id)
}


func ReadFakeInput() []string {
	input := `939
7,13,x,x,59,x,31,19`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := `1005162
19,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,823,x,x,x,x,x,x,x,23,x,x,x,x,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,443,x,x,x,x,x,37,x,x,x,x,x,x,13`
	return strings.Split(input, "\n")
}
