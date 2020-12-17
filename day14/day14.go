package day14

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operator int

const (
	Mask Operator = iota
	MemSet
)

type Command struct {
	Operator Operator
	Operand1 uint64
	Operand2 uint64
}

func Day14() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 14 part 1 answer is %d\n", Part1(input))
	input, err = parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 14 part 2 answer is %d\n", Part2(input))
}

func Part1(input chan *Command) uint64 {
	mask := [2]uint64{0, 0}
	memory := map[uint64]uint64{}

	for command := range input {
		switch command.Operator {
		case Mask:
			mask[0], mask[1] = command.Operand1, command.Operand2
		case MemSet:
			memory[command.Operand1] = applyMask(mask, command.Operand2)
		}
	}

	return sum(memory)
}

func Part2(input chan *Command) uint64 {
	mask := [2]uint64{0, 0}
	memory := map[uint64]uint64{}

	for command := range input {
		switch command.Operator {
		case Mask:
			mask[0], mask[1] = command.Operand1, command.Operand2
		case MemSet:
			addresses := generateAddresses(mask, command.Operand1)
			for _, a := range addresses {
				memory[a] = command.Operand2
			}
		}
	}

	return sum(memory)
}

func generateAddresses(mask [2]uint64, address uint64) []uint64 {
	bitsIdx := []int{}
	floating := mask[0] ^ mask[1]
	address = (^floating) & (address | mask[1])
	for i := 0; floating != 0; i++ {
		lsb := floating & 1
		if lsb == 1 {
			bitsIdx = append(bitsIdx, i)
		}

		floating = floating >> 1
	}

	addresses := []uint64{}
	var helper func(int, uint64)
	helper = func(i int, addr uint64) {
		if i == len(bitsIdx) {
			addresses = append(addresses, addr)
			return
		}

		m := uint64(1 << bitsIdx[i])
		helper(i+1, addr)
		helper(i+1, m|addr)
	}
	helper(0, 0)

	for i, a := range addresses {
		addresses[i] = a | address
	}

	return addresses
}

func applyMask(mask [2]uint64, value uint64) uint64 {
	return (mask[0] | value) & mask[1]
}

func sum(memory map[uint64]uint64) uint64 {
	s := uint64(0)
	for _, v := range memory {
		s += v
	}
	return s
}

func parseInput() (chan *Command, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	return readInput(f)
}

func readInput(r io.ReadCloser) (chan *Command, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	commands := make(chan *Command)

	go func() {
		defer r.Close()
		defer close(commands)

		for scanner.Scan() {
			command := Command{}
			line := scanner.Text()
			if line[:4] == "mask" {
				command.Operator = Mask
				command.Operand1, command.Operand2 = parseMask(line)
			} else {
				command.Operator = MemSet
				command.Operand1, command.Operand2 = parseMemSet(line)
			}

			commands <- &command
		}
	}()

	return commands, scanner.Err()
}

func parseMask(line string) (uint64, uint64) {
	x := strings.Split(line, " = ")
	o1, _ := strconv.ParseUint(strings.ReplaceAll(x[1], "X", "0"), 2, 64)
	o2, _ := strconv.ParseUint(strings.ReplaceAll(x[1], "X", "1"), 2, 64)

	return o1, o2
}

var memSetRe *regexp.Regexp = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

func parseMemSet(line string) (uint64, uint64) {
	x := memSetRe.FindAllStringSubmatch(line, 1)
	o1, _ := strconv.ParseUint(x[0][1], 10, 64)
	o2, _ := strconv.ParseUint(x[0][2], 10, 64)
	return o1, o2
}
