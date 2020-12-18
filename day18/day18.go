package day18

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Day18() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 18 part 1 answer is %d\n", Part1(input))
	fmt.Printf("Day 18 part 2 answer is %d\n", Part2(input))
}

func Part1(input []string) int {
	sum := 0

	for _, expr := range input {
		sum += calculate(expr)
	}

	return sum
}

// Reverse string, push to stack, evaluate when opening paren is met (opening paren because we are going from the back)
func calculate(expr string) int {
	stack := &Stack{}
	for i := len(expr) - 1; i >= 0; i-- {
		ch := expr[i]

		switch ch {
		case ' ':
			continue
		case '(':
			res := evaluate(stack)
			stack.Pop()
			stack.Push(res)
		case '+':
			stack.Push(Add)
		case '*':
			stack.Push(Mul)
		case ')':
			stack.Push(CloseParen)
		default:
			stack.Push(atoi(ch))
		}
	}

	return evaluate(stack)
}

func evaluate(stack *Stack) int {
	res := 0
	if len(*stack) != 0 {
		res = stack.Pop()
	}

	for len(*stack) > 0 && stack.Top() != CloseParen {
		op := stack.Pop()

		switch op {
		case Add:
			res += stack.Pop()
		case Mul:
			res *= stack.Pop()
		}
	}

	return res
}

func atoi(b byte) int {
	return int(b - '0')
}

func Part2(input []string) int {
	sum := 0

	for _, expr := range input {
		sum += calculate2(expr)
	}

	return sum
}

func calculate2(expr string) int {
	stack := &Stack{}

	var num int
	operator := Mul

	for i := 0; i < len(expr); i++ {
		ch := expr[i]

		switch ch {
		case ' ':
			continue
		case '(':
			stack.Push(operator)
			operator = Mul
		case ')':
			eval2(stack, operator, num)
			num = 1
			for len(*stack) > 0 && (stack.Top() != Add && stack.Top() != Mul) {
				num *= stack.Pop()
			}
			operator = stack.Pop()
		case '+':
			eval2(stack, operator, num)
			operator = Add
		case '*':
			eval2(stack, operator, num)
			operator = Mul
		default:
			num = atoi(ch)
		}

		if i == len(expr)-1 {
			eval2(stack, operator, num)
		}
	}

	res := 1

	for len(*stack) > 0 {
		res *= stack.Pop()
	}

	return res
}

func eval2(stack *Stack, operator int, num int) {
	if operator == Add {
		stack.Push(num + stack.Pop())
	} else {
		stack.Push(num)
	}
}

const (
	Add = iota - 1
	Mul
	OpenParen
	CloseParen
)

type Stack []int

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}

	t := s.Top()
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *Stack) Push(b ...int) {
	*s = append(*s, b...)
}

func parseInput() ([]string, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readInput(f)
}

func readInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
