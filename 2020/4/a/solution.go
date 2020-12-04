package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ReadInput()
	var total int
	fieldsSeen := make(map[string]int)
	for ind, l := range(input) {
	/*
	byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
*/
		items := strings.Split(l, " ")
		for _, item := range(items) {
			kv := strings.Split(item, ":")
			if kv[0] == "byr" || kv[0] == "iyr" || kv[0] == "eyr" || kv[0] == "hgt" || kv[0] == "hcl" || kv[0] == "ecl" || kv[0] == "pid" {
				fieldsSeen[kv[0]] = 1
			}
		}
		if len(l) == 0 || ind == len(input) - 1 {
			var sum int
			for _, v := range fieldsSeen {
				sum += v
			}
			if sum == 7 {
				total += 1
			}
			fieldsSeen = make(map[string]int)
			continue
		}
	}
	fmt.Printf("Total %d", total)
}

func ReadFakeInput() []string {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
