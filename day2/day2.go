package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	policy Policy
	pass   string
}

type Policy struct {
	char byte
	lo   int
	hi   int
}

func Day2() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 2 part 1 answer is %d\n", CountValidPasswords1(input))
	fmt.Printf("Day 2 part 2 answer is %d\n", CountValidPasswords2(input))
}

func CountValidPasswords1(input []*Entry) int {
	out := 0

	for _, entry := range input {
		if isValid1(entry) {
			out++
		}
	}

	return out
}

func CountValidPasswords2(input []*Entry) int {
	out := 0

	for _, entry := range input {
		if isValid2(entry) {
			out++
		}
	}

	return out
}

func isValid1(entry *Entry) bool {
	count := 0
	for _, r := range entry.pass {
		if byte(r) == entry.policy.char {
			count++
		}

		if count > entry.policy.hi {
			return false
		}
	}

	return count >= entry.policy.lo
}

func isValid2(entry *Entry) bool {
	return len(entry.pass) >= entry.policy.lo && ((entry.pass[entry.policy.lo-1] == entry.policy.char && (len(entry.pass) >= entry.policy.hi && entry.pass[entry.policy.hi-1] != entry.policy.char ||
		len(entry.pass) < entry.policy.hi)) || (entry.pass[entry.policy.lo-1] != entry.policy.char &&
		len(entry.pass) >= entry.policy.hi && entry.pass[entry.policy.hi-1] == entry.policy.char))

}

func parseInput() ([]*Entry, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	return readInput(f)
}

func readInput(r io.Reader) ([]*Entry, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []*Entry
	for scanner.Scan() {
		// 1-10 a: abcd -> ["1-10 a", "abcd"]
		x := strings.SplitN(scanner.Text(), ": ", 2)
		// 1-10 a -> ["1-10", "a"]
		p := strings.SplitN(x[0], " ", 2)
		// 1-10 -> ["1", "10"]
		r := strings.SplitN(p[0], "-", 2)
		lo, err := strconv.Atoi(r[0])
		if err != nil {
			return result, err
		}
		hi, err := strconv.Atoi(r[1])
		if err != nil {
			return result, err
		}
		result = append(result, &Entry{
			pass: x[1],
			policy: Policy{
				char: p[1][0],
				lo:   lo,
				hi:   hi,
			},
		})
	}
	return result, scanner.Err()
}
