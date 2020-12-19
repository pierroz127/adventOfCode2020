package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func solveDay19Example() {
	lines := []string{
		"42: 9 14 | 10 1",
		"9: 14 27 | 1 26",
		"10: 23 14 | 28 1",
		"1: \"a\"",
		"11: 42 31",
		"5: 1 14 | 15 1",
		"19: 14 1 | 14 14",
		"12: 24 14 | 19 1",
		"16: 15 1 | 14 14",
		"31: 14 17 | 1 13",
		"6: 14 14 | 1 14",
		"2: 1 24 | 14 4",
		"0: 8 11",
		"13: 14 3 | 1 12",
		"15: 1 | 14",
		"17: 14 2 | 1 7",
		"23: 25 1 | 22 14",
		"28: 16 1",
		"4: 1 1",
		"20: 14 14 | 1 15",
		"3: 5 14 | 16 1",
		"27: 1 6 | 14 18",
		"14: \"b\"",
		"21: 14 1 | 1 14",
		"25: 1 1 | 1 14",
		"22: 14 14",
		"8: 42",
		"26: 14 22 | 1 20",
		"18: 15 15",
		"7: 14 5 | 1 21",
		"24: 14 1",
		"",
		"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
		"bbabbbbaabaabba",
		"babbbbaabbbbbabbbbbbaabaaabaaa",
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
		"bbbbbbbaaaabbbbaaabbabaaa",
		"bbbababbbbaaaaaaaabbababaaababaabab",
		"ababaaaaaabaaab",
		"ababaaaaabbbaba",
		"baabbaaaabbaaaababbaababb",
		"abbbbabbbbaaaababbbbbbaaaababb",
		"aaaaabbaabaaaaababaa",
		"aaaabbaaaabbaaa",
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
		"babaaabbbaaabaababbaabababaaab",
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
	}

	rules, messages := parseLines(lines, true /* part 2 */)
	fmt.Println("Patterns:")
	for i, pattern := range rules {
		if len(pattern) > 0 {
			fmt.Printf("- rule[%d]: %s\n", i, pattern)
		}
	}
	fmt.Println("=====")
	fmt.Printf("Messages: %v\n", messages)

	countValidMessages(messages, rules[0])
}

func solveDay19Part1() {
	rules, messages := getDay19Data(false /* part 1 */)
	countValidMessages(messages, rules[0])
}

func countValidMessages(messages []string, pattern string) {
	count := 0
	re := regexp.MustCompile(`^` + pattern + `$`)
	for _, message := range messages {
		found := re.FindString(message)
		if len(found) > 0 {
			fmt.Printf("it's a match for %s !\n", message)
			count++
		}
	}
	fmt.Printf("There are %d valid messages\n", count)
}

func solveDay19Part2() {
	rules, messages := getDay19Data(true /* part 2 */)
	fmt.Printf("Nbr of rules: %d\n", len(rules))
	fmt.Println("=====")
	countValidMessages(messages, rules[0])
}

func getDay19Data(part2 bool) ([]string, []string) {
	return parseLines(getDataFromFile("day19"), part2)
}

func parseLines(lines []string, part2 bool) ([]string, []string) {
	idx := 0
	for strings.Contains(lines[idx], ":") {
		idx++
	}

	rules := parseDecryptRules(sortDecryptRules(lines[:idx]), part2)
	return rules, lines[idx+1:]
}

func sortDecryptRules(lines []string) []string {
	rules := []string{}
	for _, line := range lines {
		arr := strings.Split(line, ":")
		key, _ := strconv.Atoi(arr[0])
		for key >= len(rules) {
			rules = append(rules, "")
		}
		rules[key] = strings.TrimSpace(arr[1])
	}
	return rules
}

func parseDecryptRules(lines []string, part2 bool) []string {
	patterns := []string{}
	for i := 0; i < len(lines); i++ {
		patterns = append(patterns, "")
	}

	patterns[0] = parseDecryptRule(lines[0], lines, patterns, part2)
	return patterns
}

func parseDecryptRule(input string, lines []string, patterns []string, part2 bool) string {
	if strings.Contains(input, "\"") {
		return input[1 : len(input)-1]
	}

	result := ""
	if strings.Contains(input, "|") {
		for _, part := range strings.Split(input, "|") {
			partialPattern := parseDecryptRule(strings.TrimSpace(part), lines, patterns, part2)
			if len(result) == 0 {
				result = partialPattern
			} else {
				result = fmt.Sprintf("%s|%s", result, partialPattern)
			}
		}
		return fmt.Sprintf("(%s)", result)
	}

	for _, sID := range strings.Split(input, " ") {
		ruleID, _ := strconv.Atoi(sID)
		if len(patterns[ruleID]) == 0 {
			var pattern string
			if part2 && ruleID == 8 {
				pattern = fmt.Sprintf("(%s)+", parseDecryptRule("42", lines, patterns, part2))
			} else if part2 && ruleID == 11 {
				pattern42 := parseDecryptRule("42", lines, patterns, part2)
				pattern31 := parseDecryptRule("31", lines, patterns, part2)
				pattern = "(" + pattern42 + pattern31
				for i := 2; i < 40; i++ { // a bit ugly but I need the number of repeated matches for rule 42 to be the same as for rule 31...
					pattern = fmt.Sprintf("%s|%s{%d}%s{%d}",
						pattern,
						pattern42,
						i,
						pattern31,
						i)
				}
				pattern += ")"
			} else {
				pattern = parseDecryptRule(strings.TrimSpace(lines[ruleID]), lines, patterns, part2)
			}
			patterns[ruleID] = pattern
		}
		result += patterns[ruleID]
	}
	return result
}
