package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func checkPassport(passport map[string]string) bool {
	return passport["byr"] != "" &&
		passport["iyr"] != "" &&
		passport["eyr"] != "" &&
		passport["hgt"] != "" &&
		passport["hcl"] != "" &&
		passport["ecl"] != "" &&
		passport["pid"] != ""
}

func checkYear(s string, min int, max int) bool {
	if len(s) != 4 {
		return false
	}
	var year, err = strconv.Atoi(s)
	if err != nil {
		return false
	}
	return year >= min && year <= max
}
func checkIssueYear(s string) bool {
	if len(s) != 4 {
		return false
	}
	var year, err = strconv.Atoi(s)
	if err != nil {
		return false
	}
	return year >= 2010 && year <= 2020
}

func checkExpirationYear(s string) bool {
	if len(s) != 4 {
		return false
	}
	var year, err = strconv.Atoi(s)
	if err != nil {
		return false
	}
	return year >= 2020 && year <= 2030
}

func checkheight(s string) bool {
	if len(s) < 4 {
		return false
	}
	if s[len(s)-2:] == "in" {
		var height, err = strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return false
		}
		return height >= 59 && height <= 76
	} else if s[len(s)-2:] == "cm" {
		var height, err = strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return false
		}
		return height >= 150 && height <= 193
	}
	return false
}

func checkHair(s string) bool {
	if len(s) == 0 || s[0] != '#' || len(s) != 7 {
		return false
	}
	for _, c := range s[1:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	return true
}

func checkEye(s string) bool {
	return s == "amb" ||
		s == "blu" ||
		s == "brn" ||
		s == "gry" ||
		s == "grn" ||
		s == "hzl" ||
		s == "oth"
}

func checkPid(s string) bool {
	if len(s) != 9 {
		return false
	}
	var pid, err = strconv.Atoi(s)
	if err != nil {
		return false
	}
	return pid <= 999999999 && pid >= 0
}

func secureCheckPassport(passport map[string]string) bool {
	return checkYear(passport["byr"], 1920, 2002) &&
		checkYear(passport["iyr"], 2010, 2020) &&
		checkYear(passport["eyr"], 2020, 2030) &&
		checkheight(passport["hgt"]) &&
		checkHair(passport["hcl"]) &&
		checkEye(passport["ecl"]) &&
		checkPid(passport["pid"])
}

type check func(map[string]string) bool

func ParseFile(fn check) uint {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	//var passports []map[string]string
	var validpassports uint
	var replacer = strings.NewReplacer("\n", " ")
	for _, lines := range strings.Split(string(data), "\n\n") {
		var passportstring = replacer.Replace(lines)
		var passport = make(map[string]string)
		for _, field := range strings.Split(passportstring, " ") {
			var tuple = strings.Split(field, ":")
			var key = tuple[0]
			var value = tuple[1]
			passport[key] = value
		}
		if fn(passport) {
			validpassports++
		}
	}
	return validpassports
}

func main() {
	fmt.Println(ParseFile(checkPassport))
	fmt.Println(ParseFile(secureCheckPassport))
}
