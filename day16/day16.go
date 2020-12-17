package day16

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Rules         []Rule
	MyTicket      []int
	NearbyTickets [][]int
}

type Rule struct {
	Name   string
	Ranges [][2]int
}

func Day16() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 16 part 1 answer is %d\n", Part1(input))
	fmt.Printf("Day 16 part 2 answer is %d\n", Part2(input))
}

func Part1(input *Input) int {
	rate := 0

	for _, ticket := range input.NearbyTickets {
		for _, n := range ticket {
			anyValid := false
			for _, rule := range input.Rules {
				if isValid(n, rule) {
					anyValid = true
				}
			}

			if !anyValid {
				rate += n
			}
		}
	}

	return rate
}

func Part2(input *Input) int {
	// validTickets := [][]int{input.MyTicket}

	// for _, ticket := range input.NearbyTickets {
	// 	validTicket := true

	// 	for _, n := range ticket {
	// 		anyValid := false
	// 		for _, rule := range input.Rules {
	// 			if isValid(n, rule) {
	// 				anyValid = true
	// 			}
	// 		}

	// 		if !anyValid {
	// 			validTicket = false
	// 			break
	// 		}
	// 	}
	// 	if validTicket {
	// 		validTickets = append(validTickets, ticket)
	// 	}
	// }

	// iToFields := make([]map[string]bool, len(input.MyTicket))
	// usedFields := map[string]bool{}

	// for _, ticket := range validTickets {
	// 	for i, n := range ticket {
	// 		// already assigned
	// 		if len(iToFields[i]) == 1 {
	// 			var field string
	// 			for f := range iToFields[i] {
	// 				field = f
	// 			}
	// 			usedFields[field] = true
	// 			removeField(&iToFields, i, field, &usedFields)
	// 			continue
	// 		}

	// 		validRules := map[string]bool{}
	// 		for _, rule := range input.Rules {
	// 			if !usedFields[rule.Name] && isValid(n, rule) {
	// 				validRules[rule.Name] = true
	// 			}
	// 		}

	// 		if len(iToFields[i]) == 0 {
	// 			iToFields[i] = validRules
	// 		} else {
	// 			iToFields[i] = intersection(iToFields[i], validRules)
	// 		}

	// 		if len(iToFields[i]) == 1 {
	// 			var field string
	// 			for f := range iToFields[i] {
	// 				field = f
	// 			}
	// 			usedFields[field] = true
	// 			removeField(&iToFields, i, field, &usedFields)
	// 		}

	// 	}
	// }

	// fmt.Println(iToFields)

	return 0
}

func intersection(rules1, rules2 map[string]bool) map[string]bool {
	out := map[string]bool{}

	for r := range rules1 {
		if rules2[r] {
			out[r] = true
		}
	}

	return out
}

func removeField(iToFields *[]map[string]bool, except int, fieldToRemove string, usedFields *map[string]bool) {
	for i, fields := range *iToFields {
		if i == except {
			continue
		}

		delete(fields, fieldToRemove)
	}
}

func getSingleField(fields map[string]bool) string {
	for f := range fields {
		return f
	}
	return ""
}

func isValid(n int, rule Rule) bool {
	valid := false
	for _, r := range rule.Ranges {
		if n >= r[0] && n <= r[1] {
			valid = true
		}
	}

	return valid
}

func parseInput() (*Input, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readInput(f)
}

const (
	ScanningRules int = iota
	ScanningMyTicket
	ScanningNearbyTickets
)

var scanningSequence []int = []int{
	ScanningRules,
	ScanningMyTicket,
	ScanningNearbyTickets,
}

func readInput(r io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanning := ScanningRules
	input := &Input{}
	skipNext := false

	for scanner.Scan() {
		if skipNext {
			skipNext = false
			continue
		}

		line := scanner.Text()
		if line == "" {
			scanning = scanningSequence[scanning+1]
			skipNext = true
			continue
		}

		switch scanning {
		case ScanningRules:
			rule := parseRule(line)
			input.Rules = append(input.Rules, rule)
		case ScanningMyTicket:
			ticket := parseTicket(line)
			input.MyTicket = ticket
		case ScanningNearbyTickets:
			ticket := parseTicket(line)
			input.NearbyTickets = append(input.NearbyTickets, ticket)
		}
	}

	return input, scanner.Err()
}

func parseRule(line string) Rule {
	rule := Rule{}
	x := strings.Split(line, ": ")
	rule.Name = x[0]

	for _, r := range strings.Split(x[1], " or ") {
		s := strings.Split(r, "-")
		from, _ := strconv.Atoi(s[0])
		to, _ := strconv.Atoi(s[1])
		rule.Ranges = append(rule.Ranges, [2]int{from, to})
	}

	return rule
}

func parseTicket(line string) []int {
	x := strings.Split(line, ",")
	ticket := make([]int, len(x))
	for i, s := range x {
		n, _ := strconv.Atoi(s)
		ticket[i] = n
	}
	return ticket
}
