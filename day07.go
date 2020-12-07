package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Example data ----------

func solveDay07Example() {
	bagRulePatterns := []string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
		// "light red bags contain 1 bright white bag, 2 muted yellow bags.",
		// "dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		// "bright white bags contain 1 shiny gold bag.",
		// "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		// "shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		// "dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		// "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		// "faded blue bags contain no other bags.",
		// "dotted black bags contain no other bags.",
	}

	bagRules := parseBagRulesByParent(bagRulePatterns)
	count := countAllBags("shiny gold", bagRules)
	fmt.Printf("There are %d bags contained in %s\n bag", count, "shiny gold")
}

// Part 1 ----------

func solveDay07Part1() {
	bagRulePatterns := getDay07Data()
	bagRules := parseBagRules(bagRulePatterns)
	countBag("shiny gold", bagRules)
}

func parseBagRule(bagrule string, rules map[string][]string) {
	idx := strings.Index(bagrule, " bags contain ")
	outter := bagrule[:idx]

	innerBagsRule := bagrule[idx+len(" bags contain "):]
	if innerBagsRule == "no other bags." {
		return
	}
	inners := strings.Split(innerBagsRule, ", ")
	innerRegexp := `^\d+ (?P<bag>\w+(\s\w+)*) bags?\.?$`
	var compRegEx = regexp.MustCompile(innerRegexp)

	for _, inner := range inners {
		match := compRegEx.FindStringSubmatch(inner)
		name := match[1]
		parents, ok := rules[name]
		if !ok {
			parents = []string{outter}
		} else {
			parents = append(parents, outter)
		}
		rules[name] = parents
	}
}

func parseBagRules(bagRules []string) map[string][]string {
	rules := make(map[string][]string)
	for _, rule := range bagRules {
		parseBagRule(rule, rules)
	}
	return rules
}

func countBag(bag string, rules map[string][]string) {
	queue := []string{bag}

	parsed := []string{bag}
	count := 0
	for len(queue) > 0 {
		newQueue := []string{}
		for _, b := range queue {
			parents, ok := rules[b]
			if ok {
				for _, parent := range parents {
					if !stringContains(parsed, parent) {
						newQueue = append(newQueue, parent)
						parsed = append(parsed, parent)
						count++
					}
				}
			}
		}
		queue = newQueue
	}

	fmt.Printf("There are %d outer bags that can contain a %s bag\n", count, bag)
}

// Part 2 ----------

func solveDay07Part2() {
	bagRulePatterns := getDay07Data()
	bagRules := parseBagRulesByParent(bagRulePatterns)
	count := countAllBags("shiny gold", bagRules)
	fmt.Printf("There are %d bags contained in %s\n", count, "shiny gold")
}

func parseBagRulesByParent(bagRules []string) map[string][]string {
	rules := make(map[string][]string)
	for _, rule := range bagRules {
		parseBagRuleByParent(rule, rules)
	}
	return rules
}

func parseBagRuleByParent(bagrule string, rules map[string][]string) {
	idx := strings.Index(bagrule, " bags contain ")
	outter := bagrule[:idx]

	innerBagsRule := bagrule[idx+len(" bags contain "):]
	if innerBagsRule == "no other bags." {
		return
	}

	inners := strings.Split(innerBagsRule, ", ")
	innerRegexp := `^(?P<number>\d+) (?P<bag>\w+(\s\w+)*) bags?\.?$`
	var compRegEx = regexp.MustCompile(innerRegexp)

	children := []string{}
	for _, inner := range inners {
		match := compRegEx.FindStringSubmatch(inner)
		children = append(children, fmt.Sprintf("%s#%s", match[1], match[2]))
	}
	rules[outter] = children
}

func countAllBags(bag string, rules map[string][]string) int {
	children, ok := rules[bag]
	if !ok {
		return 0
	}
	count := 0
	for _, child := range children {
		arr := strings.Split(child, "#")
		nbr, _ := strconv.Atoi(arr[0])

		count += nbr * (1 + countAllBags(arr[1], rules))
	}
	return count
}

// Data ----------

func getDay07Data() []string {
	return getDataFromFile("day07")
}
