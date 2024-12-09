package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

type Solver struct {
	files []int
	gaps  []int
}

func MakeSolver(input *bufio.Scanner) *Solver {
	s := &Solver{}

	input.Scan()
	l := input.Text()
	file := true
	for _, c := range l {
		i := Atoi(string(c))
		if file {
			s.files = append(s.files, i)
			file = false
		} else {
			s.gaps = append(s.gaps, i)
			file = true
		}
	}

	return s
}

// func (s *Solver) String() string {
// 	str := ""
// 	return str
// }

func (s *Solver) Defrag() []int {
	gap_idx := 0
	gap_curr := 0
	file_idx := len(s.files) - 1
	file_curr := s.files[file_idx] - 1

	//s.MaybeAdvanceGap(gap_idx, gap_curr)

	ret := []int{}
	for i := 0; i < s.files[0]; i++ {
		ret = append(ret, 0)
	}
	done := false
	for !done {
		// Check if curr_gap is done
		for gap_curr > s.gaps[gap_idx]-1 {
			gap_idx++
			gap_curr = 0
			// Write the file procedding the gap
			for i := 0; i < s.files[gap_idx]; i++ {
				// If we're writing the file matching
				// the one we were defragging, we're
				// done after writing the leftovers
				if gap_idx == file_idx {
					fmt.Printf("Found last file: %d, %d/%d\n", file_idx, i, file_curr)
					if i <= file_curr {
						ret = append(ret, gap_idx)
					}
					if i > file_curr {
						done = true
						break
					}
				} else {
					fmt.Printf("Filling in next file: %d, %d\n", gap_idx, i)
					ret = append(ret, gap_idx)
				}
			}

			if done {
				break
			}
		}

		if done {
			break
		}

		// Defrag once
		fmt.Printf("Defragging %d, %d into %d, %d\n", file_idx, file_curr, gap_idx, gap_curr)
		ret = append(ret, file_idx)
		gap_curr++

		file_curr--
		if file_curr < 0 {
			file_idx--
			file_curr = s.files[file_idx] - 1
		}
	}

	//fmt.Printf("Gap (%d, %d), File (%d, %d)\n", gap_idx, gap_curr, file_idx, file_curr)

	return ret
}

func (s *Solver) TotalFileSize() int {
	sum := 0
	for _, n := range s.files {
		sum += n
	}
	return sum
}

func (s *Solver) Calculate() int {
	defrag := s.Defrag()
	fmt.Println(len(defrag))
	fmt.Println(s.TotalFileSize())
	fmt.Println(defrag)
	sum := 0
	for i, n := range defrag {
		sum += i * n
	}
	return sum
}

func a(input *bufio.Scanner) {
	s := MakeSolver(input)
	fmt.Println(s.Calculate())
}

type Node struct {
	file, gap  bool
	size       int
	file_idx   int
	next, prev *Node
}

func MakeFile(idx, size int) *Node {
	n := &Node{}
	n.file = true
	n.size = size
	n.file_idx = idx
	return n
}

func MakeGap(size int) *Node {
	n := &Node{}
	n.gap = true
	n.size = size
	return n
}

type SolverLL struct {
	head, tail *Node
	files      map[int]*Node
}

func MakeSolverLL(input *bufio.Scanner) *SolverLL {
	s := &SolverLL{}
	s.files = make(map[int]*Node)

	input.Scan()
	l := input.Text()
	fileType := true
	file_idx := 0
	var prev *Node
	for _, c := range l {
		i := Atoi(string(c))
		if fileType {
			file := MakeFile(file_idx, i)
			if prev != nil {
				prev.next = file
				file.prev = prev
			} else {
				s.head = file
			}
			prev = file
			s.files[file_idx] = file
			file_idx++
			fileType = false
		} else {
			gap := MakeGap(i)
			prev.next = gap
			gap.prev = prev
			prev = gap
			fileType = true
		}
	}

	s.tail = prev

	return s
}

func (s *SolverLL) Convert() []int {
	ret := []int{}
	curr := s.head
	for curr != nil {
		if curr.file {
			for i := 0; i < curr.size; i++ {
				ret = append(ret, curr.file_idx)
			}
		}
		if curr.gap {
			for i := 0; i < curr.size; i++ {
				ret = append(ret, 0)
			}
		}
		curr = curr.next
	}
	return ret
}

func (s *SolverLL) FindGap(needed_size, file_idx int) *Node {
	curr := s.head
	for curr != nil {
		if curr.gap {
			if curr.size >= needed_size {
				return curr
			}
		} else {
			if curr.file_idx == file_idx {
				break
			}
		}
		curr = curr.next
	}
	return nil
}

func (s *SolverLL) Defrag() []int {
	i := s.tail.file_idx
	for ; i > 0; i-- {
		//fmt.Printf("%d: %v\n", i, s.Convert())
		to_defrag := s.files[i]
		needed_size := to_defrag.size
		gap := s.FindGap(needed_size, i)
		if gap != nil {
			// Slice out source, and maybe join gaps
			prev := to_defrag.prev
			next := to_defrag.next
			new_gap := MakeGap(to_defrag.size)
			if next == nil {
				if prev.gap {
					prev.size += new_gap.size
					prev.next = nil
				} else {
					prev.next = new_gap
					new_gap.prev = prev
				}
			} else {
				if prev.gap && next.gap {
					prev.size += next.size + new_gap.size
					prev.next = next.next
					if prev.next != nil {
						prev.next.prev = prev
					}
				} else if prev.gap {
					prev.size += new_gap.size
					prev.next = next
					next.prev = prev
				} else if next.gap {
					next.size += new_gap.size
					prev.next = next
					next.prev = prev
				} else {
					prev.next = new_gap
					next.prev = new_gap
					new_gap.prev = prev
					new_gap.next = next
				}
			}
			// Insert into dest, and shrink gap
			gap.prev.next = to_defrag
			to_defrag.prev = gap.prev
			gap.size -= to_defrag.size
			if gap.size == 0 {
				to_defrag.next = gap.next
				if to_defrag.next != nil {
					to_defrag.next.prev = to_defrag
				}
			} else {
				to_defrag.next = gap
				gap.prev = to_defrag
			}
		}
	}
	return s.Convert()
}

func (s *SolverLL) Calculate() int {
	defrag := s.Defrag()
	//fmt.Println(defrag)
	sum := 0
	for i, n := range defrag {
		sum += i * n
	}
	return sum
}

func b(input *bufio.Scanner) {
	s := MakeSolverLL(input)
	//fmt.Println(s)
	fmt.Println(s.Calculate())
}

func main() {
	real := flag.Bool("real", false, "Whether to use the real input")
	runA := flag.Bool("a", false, "Run program a")
	runB := flag.Bool("b", false, "Run program b")

	flag.Parse()

	fileName := "sample.txt"
	if *real {
		fileName = "input.txt"
	}

	input := ReadInput(fileName)

	if !*runA && !*runB {
		panic("Did not specify a or b")
	}

	if *runA && *runB {
		panic("Specified both a and b")
	}

	if *runA {
		a(input)
	}
	if *runB {
		b(input)
	}
}

func ReadInput(fileName string) *bufio.Scanner {
	input, err := os.Open(fileName)
	if err != nil {
		panic("could not open")
	}

	return bufio.NewScanner(input)
}
