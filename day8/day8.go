package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Operator string
	Operand  int
}

func Day8() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	acc, _ := GetAccBeforeLoop(input)
	fmt.Printf("Day 8 part 1 answer is %d\n", acc)
	fmt.Printf("Day 8 part 2 answer is %d\n", FixCodeAndGetAcc(input))
}

func GetAccBeforeLoop(commands []*Command) (int, bool) {
	ran := map[int]bool{}

	ip := 0
	acc := 0

	for !ran[ip] && ip < len(commands) {
		ran[ip] = true
		command := commands[ip]
		switch command.Operator {
		case "acc":
			acc += command.Operand
			ip++
		case "jmp":
			ip += command.Operand
		case "nop":
			ip++
		}
	}

	return acc, ip == len(commands)
}

// ugly bruteforce
// every time we see nop or jmp, try to change it and see if the program finishes on the last line
func FixCodeAndGetAcc(commands []*Command) int {
	if acc, finished := GetAccBeforeLoop(commands); finished {
		return acc
	}

	for _, command := range commands {
		switch command.Operator {
		case "nop":
			command.Operator = "jmp"
			if acc, finished := GetAccBeforeLoop(commands); finished {
				fmt.Println(command)
				return acc
			}
			command.Operator = "nop"
		case "jmp":
			command.Operator = "nop"
			if acc, finished := GetAccBeforeLoop(commands); finished {
				fmt.Println(command)
				return acc
			}
			command.Operator = "jmp"
		}
	}

	return 0
}

func parseInput() ([]*Command, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readInput(f)
}

func readInput(r io.Reader) ([]*Command, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []*Command
	for scanner.Scan() {
		result = append(result, parseCommand(scanner.Text()))
	}
	return result, scanner.Err()
}

func parseCommand(s string) *Command {
	parts := strings.SplitN(s, " ", 2)
	n, _ := strconv.Atoi(parts[1])
	return &Command{parts[0], n}
}
