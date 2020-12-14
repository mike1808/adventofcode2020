package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

func Day4() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	passportScanner := NewPassportScanner(f)
	fmt.Printf("Day 4 part 1 answer is %d\n", ValidPassports(passportScanner))
}

func ValidPassports(passportScanner *PassportScanner) int {
	valid := 0
	for passportScanner.Scan() {
		p := passportScanner.Passport()
		if isValid(p) {
			valid++
		}
	}

	return valid
}

type validator func(string) bool

var validators = map[string]validator{
	"byr": numberValidator(1920, 2002),
	"iyr": numberValidator(2010, 2020),
	"eyr": numberValidator(2020, 2030),
	"hgt": func(v string) bool {
		if len(v) < 3 {
			return false
		}
		unit := v[len(v)-2:]
		n, err := strconv.Atoi(v[:len(v)-2])
		if err != nil {
			return false
		}
		switch unit {
		case "cm":
			return 150 <= n && n <= 193
		case "in":
			return 59 <= n && n <= 76
		}

		return false
	},
	"hcl": func(v string) bool {
		match, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, v)
		return match
	},
	"ecl": func(v string) bool {
		switch v {
		case
			"amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		}
		return false
	},
	"pid": func(v string) bool {
		match, _ := regexp.MatchString(`^\d{9}$`, v)
		return match
	},
}

func numberValidator(min, max int) validator {
	return func(v string) bool {
		d, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		return d >= min && d <= max
	}
}

func isValid(p *Passport) bool {
	for field, vfunc := range validators {
		if !vfunc((*p)[field]) {
			return false
		}
	}
	return true
}

type PassportScanner struct {
	scanner  *bufio.Scanner
	done     bool
	passport *Passport
}

func NewPassportScanner(r io.Reader) *PassportScanner {
	scanner := bufio.NewScanner(r)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		for i := 0; i < len(data)-1; i++ {
			if data[i] == '\n' && data[i+1] == '\n' {
				return i + 2, data[0:i], nil
			}
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	}

	scanner.Split(split)

	return &PassportScanner{
		scanner: scanner,
	}
}

func (p *PassportScanner) Scan() bool {
	p.done = !p.scanner.Scan()
	if !p.done {
		s := bufio.NewScanner(strings.NewReader(p.scanner.Text()))
		s.Split(bufio.ScanWords)

		passport := Passport{}

		for s.Scan() {
			x := s.Text()
			v := strings.SplitN(x, ":", 2)
			passport[v[0]] = v[1]
		}

		p.passport = &passport
	}

	return !p.done
}

func (p *PassportScanner) Passport() *Passport {
	return p.passport
}
