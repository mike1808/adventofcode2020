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
	Rules         []*Rule
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

func findFieldsMapping(input *Input) []string {
	tickets := filterOnlyValidTickets(input.NearbyTickets, input.Rules)
	n := len(input.MyTicket)

	rulesToI := map[*Rule]map[int]bool{}
	mapping := make([]string, n)

	for _, rule := range input.Rules {
		if rulesToI[rule] == nil {
			rulesToI[rule] = map[int]bool{}
		}

		for i := 0; i < n; i++ {
			for _, ticket := range tickets {
				if !rulesToI[rule][i] && isValid(ticket[i], rule) {
					rulesToI[rule][i] = false
				} else {
					rulesToI[rule][i] = true
				}
			}
		}
	}

	changedRules := map[*Rule]bool{}

	for {
		changed := false

		for rule, idxs := range rulesToI {
			if changedRules[rule] {
				continue
			}

			num := 0
			idx := 0
			for i, invalid := range idxs {
				if !invalid {
					num++
					idx = i
				}
			}

			if num == 1 {
				changedRules[rule] = true
				mapping[idx] = rule.Name
				changed = true
				deleteIndex(rule, idx, rulesToI)
			}
		}

		if !changed {
			break
		}
	}

	return mapping
}

func deleteIndex(except *Rule, idx int, rulesToI map[*Rule]map[int]bool) {
	for rule, idxs := range rulesToI {
		if rule == except {
			continue
		}

		if invalid, ok := idxs[idx]; ok && !invalid {
			idxs[idx] = true
		}
	}
}

func Part2(input *Input) int {
	mapping := findFieldsMapping(input)

	res := 1

	for i, field := range mapping {
		if strings.Contains(field, "departure") {
			res *= input.MyTicket[i]
		}
	}

	return res
}

func filterOnlyValidTickets(tickets [][]int, rules []*Rule) [][]int {
	validTickets := [][]int{}

	for _, ticket := range tickets {
		validTicket := true

		for _, n := range ticket {
			anyValid := false
			for _, rule := range rules {
				if isValid(n, rule) {
					anyValid = true
				}
			}

			if !anyValid {
				validTicket = false
				break
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func isValid(n int, rule *Rule) bool {
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

func parseRule(line string) *Rule {
	rule := &Rule{}
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
