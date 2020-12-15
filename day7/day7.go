package day7

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const targetBagColor = "shiny gold"

func Day7() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 7 part 1 answer is %d\n", CountBagColorsToContain(input, targetBagColor))
	fmt.Printf("Day 7 part 1 answer is %d\n", CountBagContained(input, targetBagColor))
}

type Bag struct {
	color       string
	containedIn []Edge
	contains    []Edge
}

type Edge struct {
	bag    *Bag
	number int
}

func CountBagColorsToContain(rules []string, target string) int {
	bags := buildBagsGraph(rules)
	return getHowManyCanReachTarget(bags, bags[target])
}

func CountBagContained(rules []string, target string) int {
	bags := buildBagsGraph(rules)
	return getHowManyBagsAreInTarget(bags, bags[target])
}

func getHowManyBagsAreInTarget(bags map[string]*Bag, target *Bag) int {
	var visit func(*Bag) int
	visit = func(bag *Bag) int {
		count := 1

		for _, edge := range bag.contains {
			count += edge.number * visit(edge.bag)
		}

		return count
	}

	return visit(target) - 1
}

func getHowManyCanReachTarget(bags map[string]*Bag, target *Bag) int {
	count := 0

	// true - permanent mark, false - temporary mark
	marked := map[*Bag]bool{}

	var visit func(*Bag)
	visit = func(bag *Bag) {
		if v, ok := marked[bag]; ok && !v {
			panic("contains cycle")
		}

		marked[bag] = false

		for _, edge := range bag.containedIn {
			if !marked[edge.bag] {
				count++
				visit(edge.bag)
			}
		}

		marked[bag] = true
	}

	visit(target)

	return count
}

func buildBagsGraph(rules []string) map[string]*Bag {
	bags := map[string]*Bag{}

	for _, rule := range rules {
		color, content := parseRule(rule)
		if _, ok := bags[color]; !ok {
			bags[color] = &Bag{color: color}
		}
		bag := bags[color]

		for _, c := range content {
			if _, ok := bags[c.color]; !ok {
				bags[c.color] = &Bag{color: c.color}
			}
			containedBag := bags[c.color]
			containedBag.containedIn = append(containedBag.containedIn, Edge{bag: bag, number: c.number})
			bag.contains = append(bag.contains, Edge{bag: containedBag, number: c.number})
		}
	}

	return bags
}

type ruleEdge struct {
	number int
	color  string
}

func parseRule(rule string) (string, []*ruleEdge) {
	sep := " bags contain "
	i := strings.Index(rule, sep)
	color := rule[:i]

	containment := []*ruleEdge{}

	re := regexp.MustCompile(`(\d+) ([\w\s]+) bags?`)
	matches := re.FindAllStringSubmatch(rule[i+len(sep):], -1)
	for _, match := range matches {
		n, _ := strconv.Atoi(match[1])
		c := match[2]
		containment = append(containment, &ruleEdge{n, c})
	}

	return color, containment
}

func parseInput() ([]string, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
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
