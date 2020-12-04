package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput()
	var total int
	fieldsSeen := make(map[string]int)
	for ind, l := range(input) {
	/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
*/
		items := strings.Split(l, " ")
		for _, item := range(items) {
			kv := strings.Split(item, ":")
			if kv[0] == "byr" {
				y, err := strconv.Atoi(kv[1])
				if err == nil {
					if y >= 1920 && y <= 2002 {
						fieldsSeen[kv[0]] = 1
					}
				}
			}
			if kv[0] == "iyr" {
				y, err := strconv.Atoi(kv[1])
				if err == nil {
					if y >= 2010 && y <= 2020 {
						fieldsSeen[kv[0]] = 1
					}
				}
			}
			if kv[0] == "eyr" {
				y, err := strconv.Atoi(kv[1])
				if err == nil {
					if y >= 2020 && y <= 2030 {
						fieldsSeen[kv[0]] = 1
					}
				}
			}
			if kv[0] == "hgt" {
				lMinus2 := len(kv[1]) - 2
				prefix := kv[1][:lMinus2]
				suffix := kv[1][lMinus2:]
				
				h, err := strconv.Atoi(prefix)
				if err == nil {
					if suffix == "cm" && h >= 150 && h <= 193 {
						fieldsSeen[kv[0]] = 1
					}
					if suffix == "in" && h >= 59 && h <= 76 {
						fieldsSeen[kv[0]] = 1
					}
				}
			}
			if kv[0] == "hcl" {
				h := kv[1]
				if len(h) == 7 && h[0] == '#' {
					valid := true
					for _, i := range(h[1:]) {
						if !((i >= '0' && i <= '9') || (i >= 'a' && i <= 'f')) {
							valid = false
							break
						}
					}
					if valid {
						fieldsSeen[kv[0]] = 1
					}
				}
			}
			if kv[0] == "ecl" {
				// amb blu brn gry grn hzl oth
				if kv[1] == "amb" || kv[1] == "blu" || kv[1] == "brn" || kv[1] == "gry" || kv[1] == "grn" || kv[1] == "hzl" || kv[1] == "oth" {
					fieldsSeen[kv[0]] = 1
				}
			}
			if kv[0] == "pid" {
				if len(kv[1]) == 9 {
					if _, err := strconv.Atoi(kv[1]); err == nil {
						fieldsSeen[kv[0]] = 1
					}
				}
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
	input := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := ``
	return strings.Split(input, "\n")
}
